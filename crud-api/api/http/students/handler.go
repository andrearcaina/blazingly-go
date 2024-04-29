package students

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/blazingly-go/crud-api/database"
	"github.com/blazingly-go/crud-api/models"
)

func GetAllStudents(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		students, err := database.GetAll[models.ModelFields](db, "students", &models.Students{})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error getting students: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(students)
	}
}

func GetStudentByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id") // get the id from the query string
		n, _ := strconv.Atoi(id)

		student, err := database.GetID[models.ModelFields](db, "students", &models.Students{}, n)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error getting student: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(student)
	}
}
