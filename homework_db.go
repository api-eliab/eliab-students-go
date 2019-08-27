package main

import (
	"database/sql"
	apigo "github.com/josuegiron/api-golang"
)


func (s *School)getHomeworksDB(studentID int64) (homeworks []HomeworkResponse, err error) {

	query := `SELECT cg.id, c.id AS courseID, cg.name AS taskName, cg.weightage AS points, cg.content AS descripcion_de_tarea, cgp.comment AS observaciones_del_maestro, cg.deliver_date
	FROM course c
	JOIN course_goal cg ON cg.course_id = c.id AND cg.deleted_at IS NULL
	LEFT JOIN course_goal_person cgp ON cgp.goal_id = cg.id AND cgp.person_id = @studentID AND cgp.deleted_at IS NULL
	where cg.deliver_date > NOW()`
	query, err = getQueryString(
		query,
		sql.Named("studentID", studentID),
	)
	if apigo.Checkp(err) {
		return
	}

	rows, errR := s.db.Query(query)
	if apigo.Checkp(errR) {
		return homeworks, errR
	}

	for rows.Next() {
		var homework HomeworkResponse
		var shortDescription, longDescription sql.NullString
		err = rows.Scan(
			&homework.ID,
			&homework.ClasroomID,
			&homework.Title,
			&homework.Points,
			&shortDescription,
			&longDescription,
			&homework.DeliveryDate,
		)
		if apigo.Checkp(err) {
			return
		}

		homework.LongDescription = longDescription.String
		homework.ShortDescription = shortDescription.String
		homework.DeliveryHour = "07:00AM" // *HORA QUEMADA, HABLAR CON AQUELLOS

		homeworks = append(homeworks, homework)
	}

	return

} // end getHomeworksDB

func getHomeworkDetailDB(homeworkID int64, studentID int64) (homeworkDetail HomeworkDetail, err error) {

	query := `SELECT cg.id, c.id AS courseID, cg.name AS taskName, cg.weightage AS points, cg.content AS descripcion_de_tarea, cgp.comment AS observaciones_del_maestro, cg.deliver_date, c.name AS classroomName
	FROM course c
	JOIN course_goal cg ON cg.course_id = c.id AND cg.deleted_at IS NULL
	LEFT JOIN course_goal_person cgp ON cgp.goal_id = cg.id AND cgp.person_id = @studentID AND cgp.deleted_at IS null
	WHERE cg.id = @homeworkID `
	query, err = getQueryString(
		query,
		sql.Named("homeworkID", homeworkID),
		sql.Named("studentID", studentID),
	)
	if err != nil {
		return
	}

	row := db.QueryRow(query)

	var shortDescription, longDescription sql.NullString

	err = row.Scan(
		&homeworkDetail.ID,
		&homeworkDetail.ClassroomID,
		&homeworkDetail.Title,
		&homeworkDetail.Points,
		&shortDescription,
		&longDescription,
		&homeworkDetail.DeliveryDate,
		&homeworkDetail.ClassroomName,
	)

	homeworkDetail.LongDescription = longDescription.String
	homeworkDetail.ShortDescription = shortDescription.String
	homeworkDetail.DeliveryHour = "07:00AM" // *HORA QUEMADA, HABLAR CON AQUELLOS

	return

} // end getHomeworkDetailDB
