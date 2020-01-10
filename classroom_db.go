package main

import (
	"database/sql"
)

// Classroom doc...
type Classroom struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getClassroomsDB(studentID int64) (classrooms []Classroom, err error) {

	query := `SELECT DISTINCT mc.id, mc.name
	FROM assignation a 
	JOIN mas_period mp ON a.period_id = mp.id AND mp.current = 1 AND mp.deleted_at IS NULL
	JOIN mas_person p ON a.person_id = p.id 
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
