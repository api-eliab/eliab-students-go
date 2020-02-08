package main

import (
	"database/sql"
	"time"

	"github.com/josuegiron/log"
)

func getHomeworksDB(studentID int64) ([]Homework, error) {

	var homeworks []Homework

	query := `
		SELECT 
			cg.id, mc.id AS course, 
			cg.name AS tarea, 
			cg.content descripcion, 
			cgp.comment comentario_de_maestra,
			cgp.score * cg.weightage / 100 AS nota_para_suma_final, 
			cg.weightage AS peso_en_zona, 
			cg.deliver_date,
			cgt.name
		FROM course_goal cg
		JOIN course_goal_type cgt ON cgt.id = cg.course_goal_type_id
		JOIN course c ON c.id = cg.course_id
		JOIN section s ON s.id = c.section_id
		JOIN mas_section ms ON ms.id = s.mas_section_id
		JOIN mas_period_phase mpp ON mpp.id = c.period_phase_id
		JOIN mas_course mc ON mc.id = c.mas_course_id
		JOIN mas_grade mg ON mg.id = mc.grade_id
		LEFT JOIN course_goal_person cgp ON cgp.course_person_id = cg.id AND person_id = @studentID 
		WHERE mc.grade_id IN (SELECT  ms.grade_id
			FROM assignation a
			JOIN mas_period mp ON a.period_id = mp.id AND mp.current = 1 AND mp.deleted_at IS NULL
			JOIN section s ON s.id = a.section_id AND s.deleted_at IS NULL
			JOIN mas_section ms ON ms.id = s.mas_section_id AND ms.deleted_at IS NULL
			WHERE a.person_id = @studentID 
		)
		AND cg.deleted_at IS NULL AND
		cg.deliver_date >= NOW()
		ORDER BY 
			cg.deliver_date

	`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
	)
	if err != nil {
		return homeworks, err
	}

	log.Info(query)

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
			&homework.Type,
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

func getHomeworkDetailDB(studentID, homeworkID int64) (HomeworkDetail, error) {

	var homework HomeworkDetail

	query := `
		SELECT 
			cg.id AS tarea_id, 
			cg.weightage AS puntos_de_la_tarea, 
			cg.name titulo, 
			cg.content AS description, 
			cgp.comment comentario_de_maestra,
			mc.id AS mas_course_id, 
			mc.name AS course_name, 
			cg.deliver_date,
			(SELECT CONCAT_WS(' ', mp.first_name, mp.second_name, mp.first_last_name, mp.second_last_name)
				FROM 
					course_owner co
				JOIN mas_person mp ON mp.id = co.person_id AND mp.deleted_at IS NULL
				JOIN mas_course mc2 ON mc2.id = co.course_id AND mc2.deleted_at IS NULL
				WHERE 
					co.deleted_at IS NULL
					AND mc2.id = mc.id limit 1 
			) as teacherName,
			cgt.name AS homework_type
		FROM 
			course_goal cg
		JOIN course_goal_type cgt ON cgt.id = cg.course_goal_type_id
		JOIN course c ON c.id = cg.course_id AND c.deleted_at IS NULL
		JOIN mas_period_phase mpp ON mpp.id = c.period_phase_id AND mpp.deleted_at IS NULL
		JOIN mas_course mc ON mc.id = c.mas_course_id AND mc.deleted_at IS NULL
		JOIN mas_grade mg ON mg.id = mc.grade_id AND mg.deleted_at IS NULL
		JOIN section s ON s.id = c.section_id AND s.deleted_at IS NULL
		LEFT JOIN 
			course_goal_person cgp ON cgp.goal_id = cg.id 
			AND cgp.deleted_at IS NULL 
			AND cgp.person_id = @studentID 
		WHERE 
			cg.id = @homeworkID 
			AND cg.deleted_at IS NULL
	`

	query, err := getQueryString(
		query,
		sql.Named("homeworkID", homeworkID),
		sql.Named("studentID", studentID),
	)
	if err != nil {
		return homework, err
	}

	row := db.QueryRow(query)

	var deliverDateStr string
	var description, comments, teacherName sql.NullString

	err = row.Scan(
		&homework.ID,
		&homework.Points,
		&homework.Title,
		&description,
		&comments,
		&homework.ClassroomID,
		&homework.ClassroomName,
		&deliverDateStr,
		&teacherName,
		&homework.Type,
	)
	if err != nil {
		return homework, err
	}

	deliverDate, err := time.Parse("2006-01-02", deliverDateStr)
	if err != nil {
		return homework, err
	}

	homework.LongDescription = comments.String
	homework.ShortDescription = description.String
	homework.TeachersName = teacherName.String
	homework.DeliveryDate = deliverDate.Format("02 Jan 2006")
	//homework.DeliveryHour = deliverDate.Format("3:04PM")

	if err != nil {
		return homework, err
	}

	return homework, nil

}
