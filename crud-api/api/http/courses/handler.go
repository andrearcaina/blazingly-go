package courses

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"blazingly-go/crud-api/database"
	"blazingly-go/crud-api/models"
)

type Handler struct {
	models.BaseHandler
}

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.getAllCourses)

	return r
}

func (h *Handler) getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := database.GetAll[models.DatabaseObject](h.DB, "courses", &models.Courses{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting courses table"))
		log.Printf("Error getting courses: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)
}
