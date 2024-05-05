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

	student, err := database.GetID[models.DatabaseObject](h.DB, "professors", &models.Students{}, n)

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
