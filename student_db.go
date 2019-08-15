package main

import (
	"database/sql"
	"errors"
)

func setIconDB(studentID, iconID int64) (err error) {

	query := `UPDATE mas_person SET avatar = @iconID WHERE id = @studentID`
	query, err = getQueryString(
		query,
		sql.Named("studentID", studentID),
		sql.Named("iconID", iconID),
	)
	if err != nil {
		return
	}

	rows, err := db.Exec(query)
	if err != nil {
		return
	}

	rowAff, err := rows.RowsAffected()
	if err != nil {
		return
	}

	if rowAff == 0 {
		return errors.New("No se actualizó la tupla")
	}

	return

} // setIcon
