package main

import (
	"database/sql"
	"fmt"

	"github.com/josuegiron/log"
)

// Classroom doc...
type Classroom struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getClassroomsDB(studentID int64) (classrooms []Classroom, err error) {

	query := `SELECT mc.id, mc.name FROM mas_course mc
	WHERE mc.grade_id IN (SELECT ms.grade_id
	FROM assignation a
	JOIN mas_period mp ON a.period_id = mp.id AND mp.current = 1 AND mp.deleted_at IS NULL
	JOIN section s ON s.id = a.section_id AND s.deleted_at IS NULL
	JOIN mas_section ms ON ms.id = s.mas_section_id AND ms.deleted_at IS NULL
	WHERE a.person_id = @studentID )
	AND mc.deleted_at IS NULL`

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

func getClassroomDetailDB(studentID, courseID int64) (ClassroomDetail, error) {

	var cassroomDetail ClassroomDetail

	query := `
		SELECT 
			mc.id, 
			a.section_id, 
			mc.name, 
			mg.name
		FROM 
			mas_course mc
		JOIN 
			mas_grade mg 
				ON mg.id = mc.grade_id
		JOIN (
			SELECT 
				a.section_id, 
				ms.grade_id 
			FROM 
				assignation a
			JOIN 
				mas_period mp 
					ON a.period_id = mp.id 
					AND mp.current = 1 
					AND mp.deleted_at IS NULL
			JOIN 
				section s 
					ON s.id = a.section_id
			JOIN 
				mas_section ms 
					ON ms.id = s.mas_section_id
			WHERE 
				person_id = @studentID 
		) a ON a.grade_id = mg.id

		WHERE 
			mc.grade_id IN (
				SELECT  
					ms.grade_id
				FROM 
					assignation a
				JOIN 
					mas_period mp 
						ON a.period_id = mp.id 
						AND mp.current = 1 
						AND mp.deleted_at IS NULL
				JOIN 
					section s 
						ON s.id = a.section_id 
						AND s.deleted_at IS NULL
				JOIN 
					mas_section ms 
						ON ms.id = s.mas_section_id 
						AND ms.deleted_at IS NULL
				WHERE 
					a.person_id = @studentID 
			)

		AND mc.deleted_at IS NULL
		AND mc.id = @courseID 
	`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
		sql.Named("courseID", courseID),
	)
	if err != nil {
		return cassroomDetail, err
	}

	row := db.QueryRow(query)
	err = row.Scan(
		&cassroomDetail.ID,
		&cassroomDetail.SectionID,
		&cassroomDetail.Name,
		&cassroomDetail.Grade,
	)

	if err != nil {
		return cassroomDetail, err
	}

	return cassroomDetail, nil
}

func getTeacher(courseID int64) (Teacher, error) {

	var teacher Teacher

	query := `SELECT mp.id, mp.first_name, mp.second_name, mp.first_last_name, mp.second_last_name 
	FROM course_owner co
	JOIN mas_person mp ON mp.id = co.person_id AND mp.deleted_at IS NULL
	JOIN mas_course mc ON mc.id = co.course_id AND mc.deleted_at IS NULL
	WHERE co.deleted_at IS NULL
	AND mc.id = @courseID`

	query, err := getQueryString(
		query,
		sql.Named("courseID", courseID),
	)
	if err != nil {
		return teacher, err
	}

	var firstname, secondname, firstlastname, secondlastname sql.NullString

	row := db.QueryRow(query)
	err = row.Scan(
		&teacher.ID,
		&firstname,
		&secondname,
		&firstlastname,
		&secondlastname,
	)

	if err != nil {
		return teacher, err
	}

	teacher.Name = fmt.Sprintf("%v %v", firstname.String, secondname.String)
	teacher.LastName = fmt.Sprintf("%v %v", firstlastname.String, secondlastname.String)

	return teacher, nil

}

func getCourseDist(courseID int64) ([]CourseDist, error) {

	var courseDist []CourseDist

	query := `SELECT mpp.id AS phase_id, mpp.name AS period_phase_name
	FROM mas_period_phase mpp
	JOIN mas_period p ON p.id = mpp.period_id
	WHERE p.current = 1 `

	query, err := getQueryString(
		query,
		sql.Named("courseID", courseID),
	)
	if err != nil {
		return courseDist, err
	}

	log.Info(query)

	rows, err := db.Query(query)
	if err != nil {
		return courseDist, err
	}

	for rows.Next() {

		var dist CourseDist

		err = rows.Scan(
			&dist.ID,
			&dist.Name,
		)
		if err != nil {
			return courseDist, err
		}

		courseDist = append(courseDist, dist)

	}

	return courseDist, nil

}

func getTasksDB(distID, courseID, userID int64) ([]Task, int64, error) {

	var tasks []Task

	query := `
		SELECT 
			cg.id, 
			cg.name AS tarea, 
			cgp.score * cg.weightage / 100 AS nota_para_suma_final, 
			cg.weightage AS peso_en_zona, 
			cgp.comment
		FROM
			course_goal cg
		JOIN 
			course c 
				ON c.id = cg.course_id 
				AND c.deleted_at IS NULL
		JOIN 
			mas_period_phase mpp 
				ON mpp.id = c.period_phase_id 
				AND mpp.deleted_at IS NULL 
				AND mpp.id = @distID 
		JOIN 
			mas_period p 
				ON p.id = mpp.period_id
		JOIN 
			mas_course mc 
				ON mc.id = c.mas_course_id 
				AND mc.deleted_at IS NULL 
				AND mc.id = @courseID 
		LEFT JOIN 
			course_goal_person cgp 
				ON cgp.goal_id = cg.id 
				AND cgp.deleted_at IS NULL 
				AND cgp.person_id = @userID 
		WHERE 
			cg.deleted_at IS NULL
			AND p.current = 1
	`

	query, err := getQueryString(
		query,
		sql.Named("distID", distID),
		sql.Named("courseID", courseID),
		sql.Named("userID", userID),
	)
	if err != nil {
		return tasks, 0, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return tasks, 0, err
	}

	currentPoints := int64(0)

	for rows.Next() {

		var task Task
		var nota, total sql.NullInt64
		var comment sql.NullString

		err = rows.Scan(
			&task.ID,
			&task.Name,
			&nota,
			&total,
			&comment,
		)

		if err != nil {
			return tasks, 0, err
		}

		currentPoints += nota.Int64
		task.Points = fmt.Sprintf("%v/%v pts", nota.Int64, total.Int64)
		task.Comment = comment.String

		tasks = append(tasks, task)

	}

	return tasks, currentPoints, nil

}
