package utils

import (
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request){
	RespondwithJSON(w, 200, struct{}{})
}

func HandleError(w http.ResponseWriter, r *http.Request){
	ResponseWithError(w, "Internal error in the server", http.StatusInternalServerError)
}