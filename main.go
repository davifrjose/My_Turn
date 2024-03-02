package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/davifrjose/My_Turn/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable must be set")
	}

	dbURL := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	
	apiCfg := apiConfig{
		DB: dbQueries,
	}
	

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Post("/users", apiCfg.handlerUsersCreate)

	r.Mount("/v1",v1Router)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: r,
	}

	log.Printf("Serving on port: %s\n", port)
	
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}