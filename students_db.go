package main

import (
	"database/sql"
	"time"
)

func getHomeworksDB(studentID int64) ([]Homework, error) {

	var homeworks []Homework

	query := `SELECT cg.id, mc.id AS course, cg.name AS tarea, cg.content descripcion, cgp.comment comentario_de_maestra,
	cgp.score * cg.weightage / 100 AS nota_para_suma_final, cg.weightage AS peso_en_zona, cg.deliver_date
	FROM course_goal cg
	JOIN course c ON c.id = cg.course_id AND c.deleted_at IS NULL
	JOIN mas_period_phase mpp ON mpp.id = c.period_phase_id AND mpp.deleted_at IS NULL
	JOIN mas_period p ON p.id = mpp.period_id
	JOIN mas_course mc ON mc.id = c.mas_course_id AND mc.deleted_at IS NULL
	LEFT JOIN course_goal_person cgp ON cgp.goal_id = cg.id AND cgp.deleted_at IS NULL AND cgp.person_id = @studentID
	WHERE cg.deleted_at IS NULL
	AND p.current = 1
	AND cg.deliver_date >= now()`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
	)
	if err != nil {
		return homeworks, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return homeworks, err
	}

	for rows.Next() {

		var homework Homework
		var nota, total sql.NullInt64
		var deliverDateStr string
		var description, comments sql.NullString
		err = rows.Scan(
			&homework.ID,
			&homework.ClassroomID,
			&homework.Title,
			&description,
			&comments,
			&nota,
			&total,
			&deliverDateStr,
		)

		if err != nil {
			return homeworks, err
		}

		deliverDate, err := time.Parse("2006-01-02", deliverDateStr)
		if err != nil {
			return homeworks, err
		}

		homework.ShortDescription = comments.String
		homework.LongDescription = description.String
		homework.Points = total.Int64
		homework.DeliveryDate = deliverDate.Format("02 Jan 2006")
		//homework.DeliveryHour = deliverDate.Format("3:04PM")
		homework.DeliveryDateFormatted = deliverDate.Format("2006-01-02")

		homeworks = append(homeworks, homework)
	}

	return homeworks, nil

}