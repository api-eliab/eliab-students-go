package main

import (
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
	//"github.com/josuegiron/log"
)

func getClassroomsFromStudent(studentID int64) apigo.Response {
	classrooms, err := getClassroomsDB(studentID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los cursos del estudiante!",
			Message: "Error al consultar la información de los cursos del estudiante!",
		}
	}

	response := ClassroomsResponse{Classrooms: classrooms}
	return apigo.Success{
		Content: response,
	}
}
