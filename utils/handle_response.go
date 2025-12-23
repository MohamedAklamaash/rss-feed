package utils

import (
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request){
	RespondwithJSON(w, 200, struct{}{})
}

func HandleError(w http.ResponseWriter, r *http.Request){
	msg :=  "Internal error in the server"
	ResponseWithError(w, msg, http.StatusInternalServerError)
}