package students

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	http2 "blazingly-go/crud-api/api/http"
	"blazingly-go/crud-api/database"
	"blazingly-go/crud-api/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	http2.BaseHandler
}

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.getAllStudents)
	r.Get("/{id}", h.getStudentByID)

	return r
}

func (h *Handler) getAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := database.GetAll[models.ModelFields](h.DB, "students", &models.Students{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting student"))
		log.Printf("Error getting students: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}

func (h *Handler) getStudentByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	n, _ := strconv.Atoi(id)

	student, err := database.GetID[models.ModelFields](h.DB, "students", &models.Students{}, n)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error getting student"))
		log.Printf("Error getting student: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}
