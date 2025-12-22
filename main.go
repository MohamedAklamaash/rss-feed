package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/MohamedAklamaash/rss-feed/utils"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(errors.New(err.Error()))
		return 
	}
	port := os.Getenv("PORT")
	if port == ""{
		log.Fatalln(errors.New("PORT not set"))
	}
	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	router.Use(
			cors.Handler(
				cors.Options{
					AllowedOrigins: []string{"https://*", "http://*"},
					AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
					AllowedHeaders: []string{"Accept", "Content-Type", "X-Requested-With","X-CSRF-Token","Authorization"},
					ExposedHeaders: []string{"Link"},
					MaxAge: 300,
					AllowCredentials: false,
					},
				),
		)

	healthRouter := chi.NewRouter()
	healthRouter.Get("/healthz", utils.HandlerReadiness)
	healthRouter.Get("/errorz", utils.HandleError)

	router.Mount("/v1", healthRouter)

	log.Println("Running on port:", port)
	err = server.ListenAndServe()

	if err != nil{
		 log.Fatalln(err)
		 return
	}
}