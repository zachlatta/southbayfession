package misc

import (
	"errors"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/coopernurse/gorp"
	"github.com/zachlatta/southbayfession/models"
)

func FetchLatestTweetsManager() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"),
		os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))

	db := models.Dbm

	for {
		FetchAndCommitLatestTweets(api, db)
		time.Sleep(5 * time.Second)
	}
}

func FetchAndCommitLatestTweets(api *anaconda.TwitterApi, db gorp.SqlExecutor) {
	lastTweet, err := GetLastTweet(api, db)
	if err != nil {
		return
	}

	tweets, err := TweetsAfter(api, lastTweet)
	if err != nil {
		return
	}

	for _, tweet := range tweets {
		err := db.Insert(&tweet)
		if err != nil {
			return
		}
	}
}

func GetLastTweet(api *anaconda.TwitterApi, db gorp.SqlExecutor) (*models.Tweet, error) {
	var tweets []models.Tweet
	_, err := db.Select(&tweets, "select * from tweets order by id limit 1")
	if err != nil {
		return nil, err
	}

	var tweet *models.Tweet
	if len(tweets) == 0 {
		tweet = &models.Tweet{TwitterId: "1"}
	} else {
		tweet = &tweets[0]
	}
	return tweet, nil
}

func TweetsAfter(api *anaconda.TwitterApi, tweet *models.Tweet) (
	[]models.Tweet, error) {
	anacondaTweets, err := api.GetUserTimeline(url.Values{
		"screen_name": []string{"Southbayfession"},
		"since_id":    []string{tweet.TwitterId},
	})
	if err != nil {
		return nil, err
	}

	tweets := make([]models.Tweet, len(anacondaTweets))
	for i, t := range anacondaTweets {
		tweet := models.Tweet{
			CreatedAt: t.CreatedAt,
			TwitterId: t.IdStr,
			Text:      t.Text,
		}

		school, err := SchoolOfTweet(&tweet)
		if err != nil {
			tweet.School = "N/A"
		} else {
			tweet.School = school
		}

		tweets[i] = tweet
	}

	return tweets, nil
}

func SchoolOfTweet(tweet *models.Tweet) (string, error) {
	schoolRegex := regexp.MustCompile(`( |\w)*$`)

	matchedSchool := schoolRegex.FindString(tweet.Text)
	matchedSchool = strings.TrimSpace(matchedSchool)
	matchedSchool = strings.ToLower(matchedSchool)

	var school string
	for potentialSchool, potentialMatches := range Schools {
		for _, potentialMatch := range potentialMatches {
			if matchedSchool == potentialMatch {
				school = potentialSchool
			}
		}
	}

	if len(school) == 0 {
		return school, errors.New("No school matched.")
	}

	return school, nil
}
