package main

import (
	"encoding/json"
	"net/http"
)

func InitAPI() *http.ServeMux {
	mux := http.NewServeMux()

	PostcardHandler := &PostcardHandler{}

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]int{"version": 1})
	})

	mux.HandleFunc("GET /postcards/", PostcardHandler.ListPost)
	mux.HandleFunc("GET /postcards/{id}", PostcardHandler.ListPostByID)
	mux.HandleFunc("POST /postcards/", PostcardHandler.CreatePost)
	mux.HandleFunc("PUT /postcards/{id}", PostcardHandler.UpdatePost)
	mux.HandleFunc("DELETE /postcards/{id}", PostcardHandler.DeletePost)

	v1 := http.NewServeMux()
	// this puts the /api/v1/ prefix on all routes
	// v1 is now the main router that has all the routes previously defined for mux
	// this why we return v1 at the end of this function, not mux
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", mux))

	return v1
}
