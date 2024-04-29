package professors

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/blazingly-go/crud-api/database"
	"github.com/blazingly-go/crud-api/models"
)

func GetAllProfessors(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		professors, err := database.GetAll[models.ModelFields](db, "professors", &models.Professors{})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error getting students: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(professors)
	}
}
