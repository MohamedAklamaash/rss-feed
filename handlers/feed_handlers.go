package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/MohamedAklamaash/rss-feed/models"
	"github.com/MohamedAklamaash/rss-feed/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apicfg *APIConfig) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type params struct {
		FeedName      string `json:"feed_name"`
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	parameter := params{}
	err := decoder.Decode(&parameter)
	if err != nil {
		utils.ResponseWithError(w,"Failed to parse request body", http.StatusInternalServerError)
		return
	}

	xmlFeeds, err := utils.ParseRssXML(parameter.Url)
	if err!= nil {
		utils.ResponseWithError(w,"Post RSS Feed Handler Error", http.StatusBadRequest)
		return
	}
	cnt := len(xmlFeeds)
	feed, err := apicfg.Db.CreateFeed(r.Context(),database.CreateFeedParams{
		Name:         parameter.FeedName,
		Url:          parameter.Url,
		Createdat:    time.Now().UTC(),
		Updatedat:    time.Now().UTC(),
		ID:           uuid.New(),
		UserID:       user.ID,
		Feedquantity: int32(cnt),
		Processed:    false,
	})

	if err != nil {
		utils.ResponseWithError(w,"Failed to create feed", http.StatusInternalServerError)
		return
	}

	utils.RespondwithJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}

func (apicfg *APIConfig) GetUserFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds,err := apicfg.Db.GetUserFeeds(r.Context(),user.ID)

	if err != nil {
		utils.ResponseWithError(w,"Failed to get user feeds", http.StatusInternalServerError)
		return
	}
	var CustomFeeds []*models.Feed

	for _, feed := range feeds {
		CustomFeeds = append(CustomFeeds, models.DatabaseFeedToFeed(feed))
	}
	utils.RespondwithJSON(w, http.StatusOK, CustomFeeds)
}

func (apicfg *APIConfig) GetAllFeeds(w http.ResponseWriter, r *http.Request, _ database.User) {
	feeds, err := apicfg.Db.GetAllFeeds(r.Context())
	if err != nil {
		utils.ResponseWithError(w,"Failed to get all feeds", http.StatusInternalServerError)
		return
	}
	var CustomFeeds []*models.Feed
	for _, feed := range feeds {
		CustomFeeds = append(CustomFeeds, models.DatabaseFeedToFeed(feed))
	}
	utils.RespondwithJSON(w, http.StatusOK, CustomFeeds)
}

func (apicfg *APIConfig) GetSpecificFeed(w http.ResponseWriter, r *http.Request, _ database.User) {
	id := chi.URLParam(r, "id")
	feedId, err := uuid.Parse(id)
	if err != nil {
		utils.ResponseWithError(w,"Failed to parse request body or Feed not Found", http.StatusBadRequest)
		return
	}
	feed, err := apicfg.Db.GetSpecificFeed(r.Context(), feedId)
	if err != nil {
		utils.ResponseWithError(w,"Failed to get feed", http.StatusInternalServerError)
		return
	}
	CustomFeed := models.DatabaseFeedToFeed(feed)
	utils.RespondwithJSON(w, http.StatusOK, CustomFeed)
}