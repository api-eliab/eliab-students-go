package main

import (
	"strconv"

	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func getHomeworks(studentID int64) apigo.Response {

	homeworks, err := getHomeworksDB(studentID)
	if err != nil {

		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de las tareas del estudiante!",
			Message: "Error al consultar la información de las tareas del estudiante!",
		}

	}

	var response HomeworksResponse
	response.Homeworks = homeworks

	return apigo.Success{
		Content: response,
	}

} //	end getHomehorks

func getHomeworkDetail(homeworkID, studentID int64) apigo.Response {

	homeworkDetail, err := getHomeworkDetailDB(homeworkID, studentID)
	if err != nil {

		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de las tareas del estudiante!",
			Message: "Error al consultar la información de las tareas del estudiante!",
		}

	}

	teachers, err := getTeachersDB(homeworkDetail.ClassroomID)

	if err != nil {

		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de las tareas del estudiante!",
			Message: "Error al consultar la información de las tareas del estudiante!",
		}

	}

	homeworkDetail.TeachersID = strconv.FormatInt(teachers[0].ID, 10)
	homeworkDetail.TeachersName = teachers[0].Name + " " + teachers[0].LastName

	var response HomeworkDetailResponse
	response.HomeworkDetail = homeworkDetail

	return apigo.Success{
		Content: response,
	}

} //	end getHomeworkDetail
