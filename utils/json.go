package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondwithJSON(w  http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(data)
	if err != nil {
		return
	}
}

func ResponseWithError(w http.ResponseWriter, message string, status int) {

	type errResponse struct {
		Error string `json:"error"`
	}

	RespondwithJSON(w, status, errResponse{
		Error: message,
	})
}