package main

import (
	"log"
	"net/http"

	"blazingly-go/crud-api/api/router"
	"blazingly-go/crud-api/database"
)

func main() {
	db := database.ConnectDB()
	handler := router.InitAPI(db)

	server := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}

	log.Printf("Server listening on port %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
