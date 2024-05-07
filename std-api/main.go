package main

import (
	"log"
	"net/http"
)

func main() {
	handler := InitAPI()

	middleware := MiddlewareChain(
		LoggerMiddleware,
	)

	server := &http.Server{
		Addr:    ":8000",
		Handler: middleware(handler),
	}

	log.Printf("Server listening on port %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
