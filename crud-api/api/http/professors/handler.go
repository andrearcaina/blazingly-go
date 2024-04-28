package professors

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/blazingly-go/crud-api/models"
)

func GetALlProfessors(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := db.Query("SELECT * FROM professors")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("Error getting professors: %v", err)
		}

		defer query.Close()

		var professorsList []models.Professors

		for query.Next() {
			var professor models.Professors

			err = query.Scan(&professor.ID, &professor.FirstName, &professor.LastName, &professor.Age, &professor.Department)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Fatalf("Error scanning professors: %v", err)
			}

			professorsList = append(professorsList, professor)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(professorsList)
	}
}
