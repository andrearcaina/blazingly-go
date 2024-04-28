package router

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/blazingly-go/crud-api/api/http/courses"
	"github.com/blazingly-go/crud-api/api/http/professors"
	"github.com/blazingly-go/crud-api/api/http/students"
)

func Server(r chi.Router, db *sql.DB) {
	r.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			text := map[string]string{"message": "Hello, World!"}

			_, err := json.Marshal(text) // this is a way to convert a map to a json object

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(text) // this is a way to write a json object to the response with encoding
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/students", students.GetALlStudents(db))
		r.Get("/courses", courses.GetALlCourses(db))
		r.Get("/professors", professors.GetALlProfessors(db))
	})
}
