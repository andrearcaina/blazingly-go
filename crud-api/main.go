package main

import (
	"log"
	"net/http"

	"github.com/blazingly-go/crud-api/api/router"
	"github.com/blazingly-go/crud-api/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.ConnectDB()
	router.Server(r, db)

	log.Println("Server running on port :8000")
	http.ListenAndServe(":8000", r)
}
