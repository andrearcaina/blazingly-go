package courses

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/blazingly-go/crud-api/database"
	"github.com/blazingly-go/crud-api/models"
)

func GetAllCourses(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses, err := database.GetAll[models.ModelFields](db, "courses", &models.Courses{})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error getting courses: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(courses)
	}
}
