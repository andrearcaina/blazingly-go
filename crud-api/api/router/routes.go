package router

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"

	"blazingly-go/crud-api/api/http/courses"
	"blazingly-go/crud-api/api/http/professors"
	"blazingly-go/crud-api/api/http/students"
	"blazingly-go/crud-api/models"
)

func InitAPI(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

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

	base := models.BaseHandler{DB: db}
	studentsHandler := students.Handler{BaseHandler: base}
	coursesHandler := courses.Handler{BaseHandler: base}
	professorsHandler := professors.Handler{BaseHandler: base}

	r.Route("/api", func(r chi.Router) {
		r.Mount("/students", studentsHandler.InitRoutes())
		r.Mount("/courses", coursesHandler.InitRoutes())
		r.Mount("/professors", professorsHandler.InitRoutes())
	})

	return r
}
