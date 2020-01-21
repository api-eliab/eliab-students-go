package main

import (
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func getMessages(studentID, ownerID int64) apigo.Response {

	messages, err := getMessagesDB(studentID, ownerID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los cursos del estudiante!",
			Message: "Error al consultar la información de los cursos del estudiante!",
		}
	}

	response := MessagesResponse{Messages: messages}
	return apigo.Success{
		Content: response,
	}
}
