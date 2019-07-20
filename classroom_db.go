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
	Name       string       `json:"name"`
	Grade      string       `json:"grade"`
	Teachers   []Teacher    `json:"teachers"`
	CourseDist []CourseDist `json:"course_dist"`
}

// Teacher ...
type Teacher struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// CourseDist ...
type CourseDist struct {
	Perfect       bool   `json:"perfect"`
	Name          string `json:"name"`
	ID            int64  `json:"id"`
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
	WHERE a.person_id = @studentID AND mc.id = @classroomID`
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
	query := `SELECT g.id, g.name, g.weightage FROM course c 
	JOIN course_goal g ON g.course_id = c.id
	LEFT JOIN course_goal_person gp ON gp.goal_id = g.id AND gp.person_id = @studentID AND gp.deleted_at IS NULL
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
		var task Task
		err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Points,
		)
		if err != nil {
			return
		}

		tasks = append(tasks, task)
	}

	return
}
