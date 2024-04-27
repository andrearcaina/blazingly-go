package api

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

			json.NewEncoder(w).Encode(text)
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/students", students.GetALlStudents(db))
		r.Get("/courses", courses.GetALlCourses(db))
		r.Get("/professors", professors.GetALlProfessors(db))
	})
}
