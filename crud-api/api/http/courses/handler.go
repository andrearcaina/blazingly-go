package courses

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/blazingly-go/crud-api/models"
)

func GetALlCourses(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := db.Query("SELECT * FROM courses")

		if err != nil {
			log.Fatalf("Error getting courses: %v", err)
		}

		defer query.Close()

		var coursesList []models.Courses

		for query.Next() {
			var course models.Courses

			err = query.Scan(&course.ID, &course.StudentsID, &course.ProfessorsID, &course.CourseName, &course.CourseCode)

			if err != nil {
				log.Fatalf("Error scanning courses: %v", err)
			}

			coursesList = append(coursesList, course)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(coursesList)
	}
}
