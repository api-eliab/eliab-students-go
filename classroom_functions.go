package main

import (
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
	//"github.com/josuegiron/log"
)

func getClassroomsFromStudent(studentID int64) apigo.Response {
	var classrooms []Classroom
	classrooms, err := getClassroomsDB(studentID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los cursos del estudiante!",
			Message: "Error al consultar la información de los cursos del estudiante!",
		}
	}

	log.Println(classrooms)

	var response ClassroomsResponse
	response.Classrooms = classrooms

	return apigo.Success{
		Content: response,
	}
}

func getClasroomDetail(studentID, clasroomID int64) apigo.Response {
	return nil
}
