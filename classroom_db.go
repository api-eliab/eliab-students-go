package main

import (
	"database/sql"
	"fmt"
)

// Classroom doc...
type Classroom struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// ClassroomDetail ...
type ClassroomDetail struct {
	Name       string       `json:"name"`
	Grade      string       `json:"grade"`
	Teachers   []Teacher    `json:"teachers"`
	CourseDist []CourseDist `json:"course_dist"`
}

// Teacher ...
type Teacher struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// CourseDist ...
type CourseDist struct {
	Perfect       bool    `json:"perfect"`
	Name          string  `json:"name"`
	ID            int64   `json:"id"`
	CurrentPoints float64 `json:"current_points"`
	Tasks         []Task  `json:"tasks"`
}

// Task ...
type Task struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Points  string  `json:"points"`
	Type    bool    `json:"type"`
	Mark    float64 `json:"mark"`
	Comment string  `json:"comment"`
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

func getClassroomDetilDB(studentID, classroomID int64) (classroom ClassroomDetail, err error) {
	query := `SELECT mc.name, g.name FROM assignation a 
	JOIN section s ON s.id = a.section_id
	JOIN mas_period mp ON mp.id = s.period_id AND mp.current = 1
	JOIN mas_section ms ON ms.id = s.mas_section_id
	JOIN mas_grade g ON g.id = ms.grade_id
	JOIN mas_course mc ON mc.grade_id = g.id
	JOIN mas_level l ON l.id = g.mas_level_id
	WHERE a.person_id = @studentID AND mc.id = @classroomID`

	query, err = getQueryString(
		query,
		sql.Named("studentID", studentID),
		sql.Named("classroomID", classroomID),
	)
	if err != nil {
		return
	}

	row := db.QueryRow(query)

	err = row.Scan(
		&classroom.Name,
		&classroom.Grade,
	)
	if err != nil {
		return
	}

	return
}

func getTeachersDB(classroomID int64) (teachers []Teacher, err error) {
	query := `SELECT p.id, p.first_name, p.first_last_name FROM course_owner c
	JOIN mas_person p ON p.id = c.person_id
	JOIN mas_course mc ON mc.id = c.course_id
	WHERE mc.id = @classroomID`
	query, err = getQueryString(
		query,
		sql.Named("classroomID", classroomID),
	)
	if err != nil {
		return
	}

	rows, errR := db.Query(query)
	if errR != nil {
		return teachers, errR
	}

	for rows.Next() {
		var teacher Teacher
		err = rows.Scan(
			&teacher.ID,
			&teacher.Name,
			&teacher.LastName,
		)
		if err != nil {
			return
		}

		teachers = append(teachers, teacher)
	}

	return
}
func getCourseDistDB(studentID, classroomID int64) (courseDists []CourseDist, err error) {
	query := `SELECT c.id, c.name FROM assignation a 
	JOIN mas_period mp ON a.period_id = mp.id AND mp.current = 1
	JOIN mas_person p ON a.person_id = p.id
	LEFT JOIN course c ON c.section_id = a.section_id AND c.deleted_at IS NULL
	LEFT JOIN mas_course mc ON mc.id = c.mas_course_id
	WHERE a.person_id = @studentID AND mc.id = @classroomID 
	ORDER BY c.created_at ASC`
	query, err = getQueryString(
		query,
		sql.Named("classroomID", classroomID),
		sql.Named("studentID", studentID),
	)
	if err != nil {
		return
	}

	rows, errR := db.Query(query)
	if errR != nil {
		return courseDists, errR
	}

	for rows.Next() {
		var courseDist CourseDist
		err = rows.Scan(
			&courseDist.ID,
			&courseDist.Name,
		)
		if err != nil {
			return
		}

		courseDists = append(courseDists, courseDist)
	}

	return
}
func getTasksDB(studentID, courseDistID int64) (tasks []Task, err error) {

	query := `SELECT cg.id, cg.name AS tarea, cg.weightage AS puntos, cgp.score * cg.weightage/100 AS puntos_recibidos, cgp.comment AS observaciones_del_maestro FROM course c
		JOIN course_goal cg ON cg.course_id = c.id AND cg.deleted_at IS NULL
		LEFT JOIN course_goal_person cgp ON cgp.goal_id = cg.id AND cgp.person_id = @studentID AND cgp.deleted_at IS NULL
		WHERE c.id = @courseDistID`

	query, err = getQueryString(
		query,
		sql.Named("courseDistID", courseDistID),
		sql.Named("studentID", studentID),
	)
	if err != nil {
		return
	}

	rows, errR := db.Query(query)
	if errR != nil {
		return tasks, errR
	}

	for rows.Next() {
		var mark sql.NullFloat64
		var comment sql.NullString
		var task Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Points,
			&mark,
			&comment,
		)
		if err != nil {
			return
		}

		task.Mark = mark.Float64
		task.Comment = comment.String

		if mark.Valid {
			task.Points = fmt.Sprintf("%0.0f", mark.Float64) + "/" + task.Points + " pts"
			task.Type = true
		} else {
			task.Points = task.Points + " pts"
			task.Type = false
		}
		tasks = append(tasks, task)
	}

	return
}
