package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func getNotificationsDB(studentID, ownerID int64) ([]Notification, error) {

	var messages []Notification

	query := `
	
		SELECT DISTINCT
			ch.id,
			ch.message,
			IF( ch.owner_id IS NULL, CONCAT(p.first_name, " ", p.first_last_name), CONCAT(p1.first_name, " ", p1.first_last_name) ) AS creator,
			mc.name AS course_name,
			mc.id AS mas_course_id,
			ms.name AS section_name,
			s.id AS section_id,
			g.name AS grade_name,
			g.id AS grade_id,
			n.created_at
		FROM 
			message_chat ch
		JOIN 
			mas_person st ON st.id = ch.student_id
		JOIN 
			message_chat_notification n ON n.message_chat_id = ch.id
		JOIN 
			mas_course mc ON mc.id = ch.course_id
		JOIN 
			section s ON s.id = ch.section_id
		JOIN 
			mas_section ms ON ms.id = s.mas_section_id
		JOIN 
			mas_grade g ON g.id = ms.grade_id
		LEFT 
			JOIN mas_person p ON p.id = ch.teacher_id
		LEFT 
			JOIN mas_person p1 ON p1.id = ch.owner_id
		WHERE
			ch.approved = 1
			AND ch.student_id = @studentID 
			ORDER BY n.created_at DESC
		LIMIT 0,20

	`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
		//sql.Named("ownerID", ownerID),
	)
	if err != nil {
		return messages, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return messages, err
	}

	for rows.Next() {

		var message Notification
		var createdAtStr string

		err = rows.Scan(
			&message.ID,
			&message.Content,
			&message.TeacherName,
			&message.CourseName,
			&message.CourseID,
			&message.SectionName,
			&message.SectionID,
			&message.GradeName,
			&message.GradeID,
			&createdAtStr,
		)

		if err != nil {
			return messages, err
		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return messages, err
		}

		message.Title = message.TeacherName
		message.CreatedAt = createdAt.Format("02 Jan 06 15:04")

		messages = append(messages, message)

	}

	return messages, nil

}

func getAnnouncementsDB(studentID int64) ([]Notification, error) {

	var announcements []Notification

	query := `
	
		SELECT DISTINCT 
			msg.id,
			IF(ISNULL(msg.file), 0, 1) hasfile,
			msg.message,
			CONCAT(p.first_name, " ", p.first_last_name) AS creator,
			msg.mas_course_id,
			IF( msg.mas_course_id IS NULL, '', mc.name ) AS course_name,
			ms.name AS section_name,
			s.id AS section_id,
			g.name AS grade_name,
			g.id AS grade_id,
			msg.created_at
		FROM 	
			message_section_ok ok
		JOIN 	
			message_section msg ON msg.id = ok.message_section_id
		JOIN 	
			mas_person p ON p.id = msg.user_id
		JOIN 	
			section s ON s.id = msg.section_id
		JOIN 	
			mas_section ms ON ms.id = s.mas_section_id
		JOIN 	
			mas_grade g ON g.id = ms.grade_id
		LEFT JOIN 
			mas_course mc ON mc.id = msg.mas_course_id
		WHERE 
			ok.student_id = @studentID 
			AND msg.approve = 1
		ORDER BY 
		created_at DESC

	`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
	)

	if err != nil {
		return announcements, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return announcements, err
	}

	for rows.Next() {

		var announcement Notification
		var createdAtStr string
		var courseID sql.NullInt64

		var hasfile bool

		err = rows.Scan(
			&announcement.ID,
			&hasfile,
			&announcement.Content,
			&announcement.TeacherName,
			&courseID,
			&announcement.CourseName,
			&announcement.SectionName,
			&announcement.SectionID,
			&announcement.GradeName,
			&announcement.GradeID,
			&createdAtStr,
		)

		if err != nil {
			return announcements, err
		}

		if hasfile {

			announcement.File = fmt.Sprintf("https://alfarero.eliabc.com/open/message/file/%v", announcement.ID)

		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return announcements, err
		}

		announcement.Title = announcement.TeacherName
		announcement.CreatedAt = createdAt.Format("02 Jan 06 15:04")
		announcement.CourseID = courseID.Int64

		announcements = append(announcements, announcement)

	}

	return announcements, nil

}

func sendMessageDB(studentID, ownerID, classroomID, sectionID int64, message string) error {

	query := `INSERT INTO message_chat(
			owner_id, 
			teacher_id, 
			section_id, 
			course_id, 
			student_id,
			message, 
			approved, 
				` + "`" + "in" + "`" + `, 
			approved_by, 
			approved_at, 
			created_at, 
			updated_at, 
			deleted_at
		) 
	VALUES 
		(
			 @ownerID , 
			NULL, 
			 @sectionID , 
			 @classroomID , 
			 @studentID , 
			 @message , 
			0, 
			1, 
			NULL, 
			NULL, 
			NOW(), 
			NULL, 
			NULL
		)
	`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
		sql.Named("ownerID", ownerID),
		sql.Named("classroomID", classroomID),
		sql.Named("sectionID", sectionID),
		sql.Named("message", message),
	)
	if err != nil {
		return err
	}

	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	numRowAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numRowAffect == 0 {
		return errors.New("No se pudo ingresar el registro")
	}

	return nil

}

func geMessagesDB(studentID, ownerID, classroomID, sectionID int64) ([]Message, error) {

	var messages []Message

	query := `
	SELECT mc.message, t.first_name, t.first_last_name, ow.first_name, ow.first_last_name, .mc.created_at, mc.` + "`" + "in" + "`" + `, mc.approved
	FROM message_chat mc
	JOIN mas_person s ON s.id = mc.student_id
	JOIN section sect ON sect.id = mc.section_id
	JOIN mas_course masc ON masc.id = mc.course_id
	LEFT JOIN mas_person t ON t.id = mc.teacher_id
	LEFT JOIN mas_person ow ON ow.id = mc.owner_id
	WHERE section_id = @sectionID 
	AND course_id = @classroomID 
	AND student_id = @studentID 
	AND ( (mc.` + "`" + "in" + "`" + ` = 2 AND mc.approved = 1) OR ( mc.` + "`" + "in" + "`" + ` = 1) )
	AND mc.deleted_at IS NULL
	AND mc.created_at >= DATE_SUB(NOW(), INTERVAL 3 MONTH)
	`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
		sql.Named("classroomID", classroomID),
		sql.Named("sectionID", sectionID),
	)
	if err != nil {
		return messages, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return messages, err
	}

	for rows.Next() {

		var message Message

		var typeID int64
		var teacherFirstname, teacherLastname, ownerFirstName, ownerLastName sql.NullString
		var createdAtStr string

		err = rows.Scan(
			&message.Text,
			&teacherFirstname,
			&teacherLastname,
			&ownerFirstName,
			&ownerLastName,
			&createdAtStr,
			&typeID,
			&message.Aproved,
		)

		if err != nil {
			return messages, err
		}

		if typeID == 1 {
			message.Type = "padre"
		} else {
			message.Type = "maestro"
		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return messages, err
		}

		message.Time = createdAt.Format("02 Jan 06 15:04")
		if teacherFirstname.String == "" && teacherLastname.String == "" {
			message.TeacherName = "Pendiente..."
		} else {
			message.TeacherName = fmt.Sprintf("%v %v", teacherFirstname.String, teacherLastname.String)
		}
		message.OwnerName = fmt.Sprintf("%v %v", ownerFirstName.String, ownerLastName.String)
		messages = append(messages, message)

	}

	return messages, nil

}
