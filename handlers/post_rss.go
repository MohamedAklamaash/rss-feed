package handlers

import (
	"context"
	"database/sql"
	"log"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/MohamedAklamaash/rss-feed/utils"
	"github.com/google/uuid"
)

func (apicfg *APIConfig) ProcessUnprocessedFeeds(ctx context.Context) error {

	feeds, err := apicfg.Db.FeedsWithoutProcess(ctx)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, feed := range feeds {

		xmlFeeds, err := utils.ParseRssXML(feed.Url)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, item := range xmlFeeds {
			_, err := apicfg.Db.CreatePost(ctx, database.CreatePostParams{
				ID:    uuid.New(),
				Title: item.Title,
				Description: sql.NullString{
					String: item.Description,
					Valid:  item.Description != "",
				},
				Link:        item.Link,
				Publishedat: item.PublishedAt.Time,
				Feedid:      feed.ID,
			})
			if err != nil {
				log.Println(err)
				break
			}
		}

		_, err = apicfg.Db.MarkFeedProcessed(ctx, feed.ID)
		if err != nil {
			continue
		}
	}
	return nil
}