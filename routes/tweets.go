package routes

import (
	"log"
	"net/http"

	"github.com/coopernurse/gorp"
	"github.com/zachlatta/southbayfession/models"
)

func GetTweets(enc Encoder, db gorp.SqlExecutor) (int, string) {
	var tweets []models.Tweet
	_, err := db.Select(&tweets, "select * from Tweet order by Id desc limit 20")
	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, ""
	}
	return http.StatusOK, Must(enc.Encode(TweetsToIface(tweets)...))
}

func TweetsToIface(v []models.Tweet) []interface{} {
	if len(v) == 0 {
		return nil
	}
	ifs := make([]interface{}, len(v))
	for i, v := range v {
		ifs[i] = v
	}
	return ifs
}
