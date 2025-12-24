package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/MohamedAklamaash/rss-feed/utils"
)

func (apicfg *APIConfig) PostRssFeedHandler(w http.ResponseWriter, r *http.Request, user database.User) {

	type params struct {
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	parameter := params{}
	err := decoder.Decode(&parameter)
	if err != nil {
		utils.ResponseWithError(w,"Post RSS Feed Handler Error", http.StatusUnprocessableEntity)
		return
	}
	utils.ParseRssXML(parameter.Url)
}