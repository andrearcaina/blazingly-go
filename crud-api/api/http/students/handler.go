package students

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"blazingly-go/crud-api/database"
	"blazingly-go/crud-api/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	models.BaseHandler
}

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.getAllStudents)
	r.Get("/{id}", h.getStudentByID)
	r.Post("/", h.createStudent)

	return r
}

func (h *Handler) getAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := database.GetAll[models.DatabaseObject](h.DB, "students", &models.Students{})

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

	student, err := database.GetID[models.DatabaseObject](h.DB, "students", &models.Students{}, n)

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

func (h *Handler) createStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Students

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding student"))
		log.Printf("Error decoding student: %v", err)
		return
	}

	item, err := database.CreateObj[models.DatabaseObject](h.DB, "students", &student)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating student"))
		log.Printf("Error creating student: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
