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

func getClassroomDetail(studentID, classroomID int64) apigo.Response {

	classroom, err := getClassroomDetailDB(studentID, classroomID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información del curso del estudiante!",
			Message: "Error al consultar la información del curso del estudiante!",
		}
	}

	classroom.Teacher, err = getTeacher(classroomID)
	if err != nil {
		log.Error(err)
		// return apigo.Error{
		// 	Title:   "Error al consultar la información del encargado del curso!",
		// 	Message: "Error al consultar la información del encargado del curso!",
		// }
	}

	classroom.CourseDist, err = getCourseDist(classroomID)
	if err != nil {
		log.Error(err)
		// return apigo.Error{
		// 	Title:   "Error al consultar la información de los cursos del estudiante!",
		// 	Message: "Error al consultar la información de los cursos del estudiante!",
		// }
	}

	for i, dist := range classroom.CourseDist {
		classroom.CourseDist[i].Tasks, classroom.CourseDist[i].CurrentPoints, err = getTasksDB(dist.ID, classroomID, studentID)
		if err != nil {
			log.Error(err)
			// return apigo.Error{
			// 	Title:   "Error al consultar la información de los cursos del estudiante!",
			// 	Message: "Error al consultar la información de los cursos del estudiante!",
			// }
		}
	}

	response := ClassroomDetailResponse{ClassroomDetail: classroom}
	return apigo.Success{
		Content: response,
	}

}
