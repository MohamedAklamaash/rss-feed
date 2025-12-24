package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/MohamedAklamaash/rss-feed/models"
	"github.com/MohamedAklamaash/rss-feed/utils"
	"github.com/google/uuid"
)

func (apicfg *APIConfig) HandlecreateUser(w http.ResponseWriter, r *http.Request){
	type params struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	parameter := params{}
	err := decoder.Decode(&parameter)
	if err != nil {
		utils.ResponseWithError(w,"Error in creating user", http.StatusInternalServerError)
		return 
	}
	user, err := apicfg.Db.CreateUser(
		r.Context(),
		database.CreateUserParams{
			ID: uuid.New(),
			Name: parameter.Name,
			Createdat: time.Now().UTC(),
			Updatedat: time.Now().UTC(),
			},
		)

	if err != nil {
		utils.ResponseWithError(w,"Error in creating user with db connection error", http.StatusInternalServerError)
		return 
	}

	utils.RespondwithJSON(w, http.StatusCreated, models.DatabaseUserToUser(user))
}

func (apicfg *APIConfig) GetUserByAPIKey(w http.ResponseWriter, r *http.Request){
	apiKey, err := utils.GetApiKey(r)
	if err!=nil{
		utils.ResponseWithError(w,"Error in getting api key", http.StatusInternalServerError)
		return
	}
	user, err := apicfg.Db.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		utils.ResponseWithError(w,"User not found", http.StatusNotFound)
		return
	}
	utils.RespondwithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

func (apicfg *APIConfig) GetUserByAPIKeyWithReturn(w http.ResponseWriter, r *http.Request) (database.User, error){
	apiKey, err := utils.GetApiKey(r)
	if err!=nil{
		utils.ResponseWithError(w,"Error in getting api key", http.StatusInternalServerError)
		return database.User{}, errors.New("error in getting api key")
	}
	user, err := apicfg.Db.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		utils.ResponseWithError(w,"User not found", http.StatusNotFound)
		return database.User{}, errors.New("user not found")
	}
	return user, nil
}