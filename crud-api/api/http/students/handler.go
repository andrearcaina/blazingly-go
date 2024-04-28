package students

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/blazingly-go/crud-api/models"
)

func GetALlStudents(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := db.Query("SELECT * FROM students")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("Error getting students: %v", err)
		}

		defer query.Close()

		var studentsList []models.Students

		for query.Next() {
			var student models.Students

			err = query.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Age, &student.Major)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatalf("Error scanning students: %v", err)
			}

			studentsList = append(studentsList, student)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(studentsList)
	}
}
