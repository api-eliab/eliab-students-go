package main

import (
	"database/sql"
)

// Classroom doc...
type Classroom struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ClassroomDetail ...
type ClassroomDetail struct {
	Name    string `json:"name"`
	Teacher struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"last_name"`
	} `json:"teacher"`
	Grade      string       `json:"grade"`
	CourseDist []CourseDist `json:"course_dist"`
}

// CourseDist ...
type CourseDist struct {
	Perfect       bool   `json:"perfect"`
	Name          string `json:"name"`
	ID            int    `json:"id"`
	CurrentPoints int    `json:"current_points"`
	Tasks         []Task `json:"tasks"`
}

// Task ...
type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Points string `json:"points"`
	Type   int    `json:"type"`
}

func getClassroomsDB(studentID int64) (classrooms []Classroom, err error) {
	query := `SELECT DISTINCT mc.id, mc.name FROM assignation a 
	JOIN course c ON c.section_id = a.section_id AND c.deleted_at IS NULL
	JOIN mas_course mc ON mc.id = c.mas_course_id
	WHERE a.person_id = @studentID`
	query, err = getQueryString(
		query,
		sql.Named("studentID", studentID),
	)
	if err != nil {
		return
	}

	rows, errR := db.Query(query)
	if errR != nil {
		return classrooms, errR
	}

	for rows.Next() {
		var classroom Classroom
		err = rows.Scan(
			&classroom.ID,
			&classroom.Name,
		)
		if err != nil {
			return
		}

		classrooms = append(classrooms, classroom)
	}

	return
}

func getClassroomsDetilDB(studentID, classroomID int64) (classrooms []Classroom, err error) {
	query := `SELECT c.id, c.name, g.name FROM assignation a 
	JOIN mas_period mp ON a.period_id = mp.id AND mp.current = 1
	JOIN mas_person p ON a.person_id = p.id
	JOIN course c ON c.section_id = a.section_id AND c.deleted_at IS NULL
	JOIN mas_course mc ON mc.id = c.mas_course_id
	JOIN mas_grade g ON g.id = mc.grade_id
	WHERE a.person_id = @studentID`
	query, err = getQueryString(
		query,
		sql.Named("studentID", studentID),
	)
	if err != nil {
		return
	}

	rows, errR := db.Query(query)
	if errR != nil {
		return classrooms, errR
	}

	for rows.Next() {
		var classroom Classroom
		err = rows.Scan(
			&classroom.ID,
			&classroom.Name,
		)
		if err != nil {
			return
		}

		classrooms = append(classrooms, classroom)
	}

	return
}
