package crons

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/MohamedAklamaash/rss-feed/handlers"
	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/robfig/cron"
)

func InitCronScheduler() *cron.Cron {
	c := cron.New()
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
		return nil
	}
	conn, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalln(errors.New(err.Error()))
		return nil
	}
	apicfg := &handlers.APIConfig{
		Db: database.New(conn),
	}
	err = c.AddFunc("@every 1h", func() {
		err := apicfg.ProcessUnprocessedFeeds(context.Background())
		if err != nil {
			log.Println("feed cron error:", err)
		}
	})
	if err != nil {
		return nil
	}
	c.Start()
	fmt.Println("Cron scheduler initialized")
	return c
}