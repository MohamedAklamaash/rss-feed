package handlers

import (
	"net/http"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/MohamedAklamaash/rss-feed/utils"
)

func (apicfg *APIConfig) ReturnPosts(w http.ResponseWriter, r *http.Request, user database.User) {

	data, err := apicfg.Db.ListAllFeedPosts(r.Context(), user.ID)
	if err != nil {
		utils.ResponseWithError(w, "Error getting all posts", http.StatusUnauthorized)
		return
	}
	utils.RespondwithJSON(w, http.StatusOK, data)
}
