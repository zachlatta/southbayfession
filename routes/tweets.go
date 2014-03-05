package routes

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"../models"
	"github.com/ChimeraCoder/anaconda"
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
)

func GetTweets(enc Encoder, db gorp.SqlExecutor) (int, string) {
	var tweets []models.Tweet
	_, err := db.Select(&tweets, "select * from tweets order by id")
	if err != nil {
		checkErr(err, "select failed")
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(tweetsToIface(tweets)...))
}

func GetTweet(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Tweet{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.Tweet)
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func AddTweet(entity models.Tweet, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	tweets, err := api.GetUserTimeline(url.Values{
		"screen_name": []string{"Southbayfession"},
	})

	if err != nil {
		panic(err)
	}
	for _, tweet := range tweets {
		tweetToCommit := models.Tweet{
			CreatedAt: tweet.CreatedAt,
			TwitterId: tweet.Id,
			Text:      tweet.Text,
			School:    "ESHS",
		}
		db.Insert(&tweetToCommit)
	}
	err = db.Insert(&entity)
	if err != nil {
		checkErr(err, "insert failed")
		return http.StatusConflict, ""
	}
	w.Header().Set("Location", fmt.Sprintf("/southbayfession/tweets/%d", entity.Id))
	return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func UpdateTweet(entity models.Tweet, enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Tweet{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	oldEntity := obj.(*models.Tweet)

	entity.Id = oldEntity.Id
	_, err = db.Update(&entity)
	if err != nil {
		checkErr(err, "update failed")
		return http.StatusConflict, ""
	}
	return http.StatusOK, Must(enc.EncodeOne(entity))
}

func DeleteTweet(db gorp.SqlExecutor, parms martini.Params) (int, string) {
	id, err := strconv.Atoi(parms["id"])
	obj, _ := db.Get(models.Tweet{}, id)
	if err != nil || obj == nil {
		checkErr(err, "get failed")
		// Invalid id, or does not exist
		return http.StatusNotFound, ""
	}
	entity := obj.(*models.Tweet)
	_, err = db.Delete(entity)
	if err != nil {
		checkErr(err, "delete failed")
		return http.StatusConflict, ""
	}
	return http.StatusNoContent, ""
}

func tweetsToIface(v []models.Tweet) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}
