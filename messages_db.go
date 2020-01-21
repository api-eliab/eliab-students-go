package main

import (
	"database/sql"

	"github.com/josuegiron/log"
)

func getMessagesDB(studentID, ownerID int64) ([]Message, error) {

	var messages []Message

	query := `SELECT mk.id, ms.message, CONCAT(IFNULL(mo.first_name,''), " ",IFNULL(mo.second_name,""), " ",IFNULL(mo.first_last_name,''), " ",IFNULL(mo.second_last_name,'')) AS encargado, ms.created_at 
	FROM message_student_ok mk
	JOIN message_student ms ON ms.id = mk.message_student_id AND ms.deleted_at IS NULL
	JOIN mas_person mp ON mp.id = ms.student_id AND mp.deleted_at IS NULL
	JOIN mas_person mo ON mo.id = mk.owner_id AND mo.deleted_at IS NULL
	WHERE ms.approve = 1
	AND student_id = @studentID 
	AND owner_id = @ownerID 
	AND mk.deleted_at IS NULL`

	query, err := getQueryString(
		query,
		sql.Named("studentID", studentID),
		sql.Named("ownerID", ownerID),
	)
	if err != nil {
		return messages, err
	}

	log.Info(query)
	rows, err := db.Query(query)
	if err != nil {
		return messages, err
	}

	for rows.Next() {

		var message Message

		err = rows.Scan(
			&message.ID,
			&message.Content,
			&message.Title,
			&message.CreateAt,
		)

		if err != nil {
			return messages, err
		}

		messages = append(messages, message)

	}

	return messages, nil

}
