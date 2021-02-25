package news

import (
	"time"

	"github.com/dghubble/go-twitter/twitter"
)

type New struct {
	ID     int64
	ImgUrl string
	Name   string
	Text   string
	Time   time.Time
}

type Content struct {
	News []New
}

func Search(username string) Content {
	tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: username,
	})

	News := []New{}
	for _, tweet := range tweets {

		n := New{}
		n.ID = tweet.ID
		n.ImgUrl = tweet.User.ProfileImageURLHttps
		n.Name = tweet.User.Name
		n.Text = tweet.Text
		n.Time, _ = tweet.CreatedAtTime()
		News = append(News, n)
	}
	c := Content{News}
	return c
}
