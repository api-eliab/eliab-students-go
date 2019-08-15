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

	var response ClassroomsResponse
	response.Classrooms = classrooms

	return apigo.Success{
		Content: response,
	}
}

func getClassroomDetailFromUser(studentID, classroomID int64) apigo.Response {
	var err error
	classroomDetail, err := getClassroomDetilDB(studentID, classroomID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los cursos del estudiante!",
			Message: "Error al consultar la información de los cursos del estudiante!",
		}
	}

	classroomDetail.Teachers, err = getTeachersDB(classroomID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los profesores!",
			Message: "Error al consultar la información de los profesores!",
		}
	}

	classroomDetail.CourseDist, err = getCourseDistDB(studentID, classroomID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los bimestres!",
			Message: "Error al consultar la información de los bimestres!",
		}
	}

	for i := 0; i < len(classroomDetail.CourseDist); i++ {

		classroomDetail.CourseDist[i].Tasks, err = getTasksDB(studentID, classroomDetail.CourseDist[i].ID)
		if err != nil {
			log.Error(err)
			return apigo.Error{
				Title:   "Error al consultar la información de las tareas!",
				Message: "Error al consultar la información de las tareas!",
			}
		}

		var currentPoints float64

		for _, task := range classroomDetail.CourseDist[i].Tasks {

			currentPoints += task.Mark

		}

		classroomDetail.CourseDist[i].CurrentPoints = currentPoints

	}

	var response ClassroomDetailResponse
	response.Classroom = classroomDetail

	return apigo.Success{
		Content: response,
	}
}
