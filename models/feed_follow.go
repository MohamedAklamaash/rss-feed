package models

import (
	"time"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/google/uuid"
)

type CreateFeedFollowParams struct {
	ID            uuid.UUID `json:"id"`
	Createdat     time.Time `json:"createdat"`
	Updatedat     time.Time `json:"updatedat"`
	UserID        uuid.UUID `json:"user_id"`
	FeedID        uuid.UUID `json:"feed_id"`
	Lastfetchedat time.Time `json:"lastFetchedat"`
}

func DatabaseFeedFollowToFeedFollow(feed *database.Feedfollow) *CreateFeedFollowParams {
	return &CreateFeedFollowParams{
		ID: feed.ID,
		Createdat: feed.Createdat,
		Updatedat: feed.Updatedat,
		UserID: feed.UserID,
		FeedID: feed.FeedID,
		Lastfetchedat: feed.Lastfetchedat,
	}
}