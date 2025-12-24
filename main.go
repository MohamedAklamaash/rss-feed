package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MohamedAklamaash/rss-feed/handlers"
	"github.com/MohamedAklamaash/rss-feed/internal/database"
	"github.com/MohamedAklamaash/rss-feed/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // included but called using sqlc
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(errors.New(err.Error()))
		return
	}
	port := os.Getenv("PORT")
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
		return
	}
	if port == "" {
		log.Fatalln(errors.New("PORT not set"))
	}

	conn, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalln(errors.New(err.Error()))
		return
	}

	apicfg := &handlers.APIConfig{
		Db: database.New(conn),
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
				AllowedHeaders:   []string{"Accept", "Content-Type", "X-Requested-With", "X-CSRF-Token", "Authorization"},
				ExposedHeaders:   []string{"Link"},
				MaxAge:           300,
				AllowCredentials: false,
			},
		),
	)

	version1Router := chi.NewRouter()
	version1Router.Get("/healthz", utils.HandlerReadiness)
	version1Router.Get("/errorz", utils.HandleError)

	//  user endpoints
	version1Router.Post("/user/create", apicfg.HandlecreateUser)
	version1Router.Get("/user/getuser", apicfg.GetUserByAPIKey)

	// feed endpoints
	version1Router.Post("/feed/create", apicfg.AuthMiddleware(apicfg.CreateFeed))
	version1Router.Get("/feed/{id}",apicfg.AuthMiddleware(apicfg.GetSpecificFeed))
	version1Router.Get("/feed/all",apicfg.AuthMiddleware(apicfg.GetAllFeeds))
	version1Router.Get("/user/feed",apicfg.AuthMiddleware(apicfg.GetUserFeed))

	// feed followers endpoints

	// Rss destructing
	version1Router.Post("/post/create",apicfg.AuthMiddleware(apicfg.PostRssFeedHandler))
	router.Mount("/v1", version1Router)

	log.Println("Running on port:", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
		return
	}
}
