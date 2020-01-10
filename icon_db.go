package main

import (
	"database/sql"
	"errors"
)

func updateIconDB(studentID, iconID int64) error {

	query := `UPDATE mas_person SET avatar = @iconID WHERE id = @studentID`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
		sql.Named("iconID", iconID),
	)
	if err != nil {
		return err
	}

	result, err := db.Exec(query)
	if err != nil {
		return err
	}

	numRowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numRowAffected == 0 {
		return errors.New("No fue posible actualizar la tupla")
	}

	return nil

}
