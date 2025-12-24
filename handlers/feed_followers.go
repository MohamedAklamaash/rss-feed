package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/MohamedAklamaash/rss-feed/models"
	"github.com/MohamedAklamaash/rss-feed/utils"
	"github.com/google/uuid"
)

func (apicfg *APIConfig) CreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		FeedId string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	parameter := params{}
	err := decoder.Decode(&parameter)
	if err != nil {
		utils.ResponseWithError(w, "Error in parsing request body", http.StatusBadRequest)
		return
	}
	feedId, err := uuid.Parse(parameter.FeedId)
	if err != nil {
		utils.ResponseWithError(w, "Error in parsing request body of followerUserId", http.StatusBadRequest)
		return
	}
	feedFollow, feedErr := apicfg.Db.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:            uuid.New(),
		Createdat:     time.Now().UTC(),
		Updatedat:     time.Now().UTC(),
		UserID:        user.ID,
		FeedID:        feedId,
		Lastfetchedat: time.Now(),
	})
	if feedErr != nil {
		utils.ResponseWithError(w, "Error in creating feed follower", http.StatusBadRequest)
		return
	}
	feed := models.DatabaseFeedFollowToFeedFollow(&feedFollow)
	utils.RespondwithJSON(w, http.StatusCreated, feed)
}
