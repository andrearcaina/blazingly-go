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

type ModelFields interface {
	GetFields() []interface{}
	NewInstance() ModelFields
}

func (s *Students) GetFields() []interface{} {
	return []interface{}{&s.ID, &s.FirstName, &s.LastName, &s.Age, &s.Major}
}

func (c *Courses) GetFields() []interface{} {
	return []interface{}{&c.ID, &c.StudentsID, &c.ProfessorsID, &c.CourseName, &c.CourseCode}
}

func (p *Professors) GetFields() []interface{} {
	return []interface{}{&p.ID, &p.FirstName, &p.LastName, &p.Age, &p.Department}
}

func (s *Students) NewInstance() ModelFields {
	return &Students{}
}

func (c *Courses) NewInstance() ModelFields {
	return &Courses{}
}

func (p *Professors) NewInstance() ModelFields {
	return &Professors{}
}
