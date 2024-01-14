package main

import (
	"context"
	"log/slog"
	"osr/db"

	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
)

func Feeds(q *db.Queries) {
	for _, feedUrl := range FeedUrls() {
		slog.Info("fetching feed", "url", feedUrl)
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(feedUrl)
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		f, err := q.CreateFeed(context.Background(), db.CreateFeedParams{
			Title: feed.Title,
			Link:  feed.Link,
		})
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		for _, item := range feed.Items {
			//log.Printf("%#v", item)
			_, err := q.CreateItem(context.Background(), db.CreateItemParams{
				FeedID:      &f.FeedID,
				Title:       item.Title,
				Link:        item.Link,
				Categories:  item.Categories,
				PublishedAt: *item.PublishedParsed,
				Content:     bluemonday.NewPolicy().Sanitize(item.Content),
			})
			if err != nil {
				slog.Error(err.Error())
				continue
			}
		}
	}
}
