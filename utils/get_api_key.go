package utils

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	key := strings.Split(authHeader, " ")
	if len(key) <= 1 {
		return "", errors.New("no auth header sent")
	}
	if key[1] == "" {
		return "", errors.New("no auth header")
	}
	return key[1], nil
}
