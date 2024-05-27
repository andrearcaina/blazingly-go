package courses

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"

	"blazingly-go/crud-api/database"
	"blazingly-go/crud-api/models"
)

type Handler struct {
	models.BaseHandler
}

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.getAllCourses)
	r.Get("/{id}", h.getCourseByID)
	r.Post("/", h.createCourse)

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

func (h *Handler) getCourseByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	n, _ := strconv.Atoi(id)

	student, err := database.GetID[models.DatabaseObject](h.DB, "courses", &models.Courses{}, n)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error getting course"))
		log.Printf("Error getting course: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (h *Handler) createCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Courses

	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding course"))
		log.Printf("Error decoding course: %v", err)
		return
	}

	item, err := database.CreateObj[models.DatabaseObject](h.DB, "courses", &course)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating course"))
		log.Printf("Error creating course: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
