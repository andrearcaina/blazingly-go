package models

type Students struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Major     string `json:"major"`
}

type Courses struct {
	ID           int    `json:"id"`
	StudentsID   int    `json:"students_id"`
	ProfessorsID int    `json:"professors_id"`
	CourseName   string `json:"course_name"`
	CourseCode   string `json:"course_code"`
}

type Professors struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}
