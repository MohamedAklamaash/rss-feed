package handlers

import (
	"log"
	"net/http"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *APIConfig) AuthMiddleware(handler authHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := apicfg.GetUserByAPIKeyWithReturn(w, r)

		if err != nil {
			log.Println(err)
			return
		}
		handler(w, r, user)
	}
}
