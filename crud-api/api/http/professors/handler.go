package professors

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

	r.Get("/", h.getAllProfessors)
	r.Get("/{id}", h.getProfessorByID)
	r.Post("/", h.createProfessor)

	return r
}

func (h *Handler) getAllProfessors(w http.ResponseWriter, r *http.Request) {
	professors, err := database.GetAll[models.DatabaseObject](h.DB, "professors", &models.Professors{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting professors table"))
		log.Printf("Error getting students: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(professors)
}

func (h *Handler) getProfessorByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	n, _ := strconv.Atoi(id)

	student, err := database.GetID[models.DatabaseObject](h.DB, "professors", &models.Professors{}, n)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error getting professor"))
		log.Printf("Error getting professor: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func (h *Handler) createProfessor(w http.ResponseWriter, r *http.Request) {
	var professor models.Professors

	err := json.NewDecoder(r.Body).Decode(&professor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding professor"))
		log.Printf("Error decoding professor: %v", err)
		return
	}

	item, err := database.CreateObj[models.DatabaseObject](h.DB, "professors", &professor)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating professor"))
		log.Printf("Error creating professor: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
