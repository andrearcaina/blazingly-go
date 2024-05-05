package router

import (
	http2 "blazingly-go/crud-api/api/http"
	"blazingly-go/crud-api/api/http/professors"
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"blazingly-go/crud-api/api/http/courses"
	"blazingly-go/crud-api/api/http/students"
)

func Server(r chi.Router, db *sql.DB) {
	base := http2.BaseHandler{DB: db}
	studentsHandler := students.Handler{BaseHandler: base}
	coursesHandler := courses.Handler{BaseHandler: base}
	professorsHandler := professors.Handler{BaseHandler: base}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		text := map[string]string{"message": "Hello, World!"}

		output, err := json.Marshal(text) // convert the defined map to a json object

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println(string(output)) // print the json object to the console

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(text) // write a map to the response with json encoding
	})

	r.Route("/api", func(r chi.Router) {
		r.Mount("/students", studentsHandler.InitRoutes())
		r.Mount("/courses", coursesHandler.InitRoutes())
		r.Mount("/professors", professorsHandler.InitRoutes())
	})
}
