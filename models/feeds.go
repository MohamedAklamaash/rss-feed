package models

import (
	"time"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"createdat"`
	Updatedat time.Time `json:"updatedat"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserID    uuid.NullUUID `json:"user_id"`
}

func DatabaseFeedToFeed(feed database.Feed) *Feed {
	return &Feed{
		ID:        feed.ID,
		Createdat: feed.Createdat,
		Updatedat: feed.Updatedat,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    uuid.NullUUID{UUID:  feed.UserID, Valid: true},
	}
}