package main

import (
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func getNotifications(studentID, ownerID int64) apigo.Response {

	notifications, err := getNotificationsDB(studentID, ownerID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los cursos del estudiante!",
			Message: "Error al consultar la información de los cursos del estudiante!",
		}
	}

	response := NotificationsResponse{Notifications: notifications}
	return apigo.Success{
		Content: response,
	}
}

func getAnnouncements(studentID, ownerID int64) apigo.Response {

	notifications, err := getAnnouncementsDB(studentID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los cursos del estudiante!",
			Message: "Error al consultar la información de los cursos del estudiante!",
		}
	}

	response := NotificationsResponse{Notifications: notifications}
	return apigo.Success{
		Content: response,
	}

}

func sendMessage(studentID, ownerID, classroomID, sectionID int64, message string) apigo.Response {

	err := sendMessageDB(studentID, ownerID, classroomID, sectionID, message)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al enviar el mensaje!",
			Message: "Error al enviar el mensaje!",
		}
	}

	return apigo.Success{}
}

func getMessages(studentID, ownerID, classroomID, sectionID int64) apigo.Response {

	messages, err := geMessagesDB(studentID, ownerID, classroomID, sectionID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al obtener los mensajes del estudiante!",
			Message: "Error al obtener los mensajes del estudiante!",
		}
	}

	response := MessagesResponse{Messages: messages}
	return apigo.Success{
		Content: response,
	}
}
