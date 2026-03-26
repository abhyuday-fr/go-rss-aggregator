package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abhyuday-fr/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}


func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT") // we don't want to use 'export PORT' all the time so we are gonna get the github.com/joho/godotenv package and do 'go mod vendor'
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL") // we don't want to use 'export PORT' all the time so we are gonna get the github.com/joho/godotenv package and do 'go mod vendor'
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil{
		log.Fatal("Can't Connect to database: ", err)
	}

	apiCfg := apiConfig{
		
	}

	//routee
	router := chi.NewRouter()

	//cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//to check status, mounted router
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) // if we use HandleFunc instead of GET then it will respond to POST too
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	

	fmt.Println("Port:", portString)
}
