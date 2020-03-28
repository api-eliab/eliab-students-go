package main

import (
	"database/sql"
	"errors"
	"strings"
)

// CourseTaskFile doc ...
type CourseTaskFile struct {
	ID          int64
	CourseID    int64
	StudentID   int64
	UserID      int64
	Description string
	File        interface{}
	Extension   string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

func saveFileInDB(courseTaskFile CourseTaskFile) error {

	query := `INSERT INTO course_task_file(
				course_id, 
				student_id, 
				user_id, 
				file, 
				description,
				extension, 
				created_at, 
				updated_at, 
				deleted_at
			) 
		VALUES 
			(
				(
					SELECT 
						id 
					FROM 
						course 
					WHERE 
						mas_course_id = @courseID 
						AND period_phase_id = (SELECT id FROM mas_period_phase WHERE NOW() between start_date and end_date AND deleted_at IS NULL) 
						AND section_id = (SELECT section_id FROM assignation WHERE person_id = @studentID AND deleted_at IS NULL)
						AND deleted_at IS NULL
						ORDER BY id DESC
						LIMIT 1
				), 
				 @studentID , 
				 @userID , 
				 @file , 
				 @description , 
				 @extension , 
				NOW(), 
				NULL, 
				NULL
			)
		`

	query, err := getQueryString(
		query,
		sql.Named("courseID", courseTaskFile.CourseID),
		sql.Named("studentID", courseTaskFile.StudentID),
		sql.Named("userID", courseTaskFile.UserID),
		sql.Named("file", courseTaskFile.File),
		sql.Named("description", courseTaskFile.Description),
		sql.Named("extension", courseTaskFile.Extension),
	)
	if err != nil {
		return err
	}

	// log.Info(query)

	query = strings.Replace(query, "unknown", "?", -1)

	result, err := db.Exec(query, courseTaskFile.File)
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
