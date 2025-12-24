package models

import (
	"time"

	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Createdat: dbUser.Createdat,
		Updatedat: dbUser.Updatedat,
		ApiKey:    dbUser.ApiKey,
	}
}
