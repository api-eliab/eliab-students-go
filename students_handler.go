package main

import (
	"encoding/json"
	"log"
	"net/http"

	apigo "github.com/josuegiron/api-golang"
	apigolang "github.com/josuegiron/api-golang"
)

func getHomeworksHandler(w http.ResponseWriter, r *http.Request) {
	request := apigolang.Request{
		HTTPReq: r,
	}
	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
	}

	response = getHomeworks(studentID)
	apigo.SendResponse(response, w)
	return

}

func getHomeworkDetail(w http.ResponseWriter, r *http.Request) {
	request := apigolang.Request{
		HTTPReq: r,
	}
	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}
	homeworkID, responseh := request.GetURLParamInt64("homeworkID")
	if responseh != nil {
		apigolang.SendResponse(responseh, w)
		return
	}

	log.Println(studentID)
	log.Println(homeworkID)

	muckData := []byte(`
	{
			"homework": {
					"id": 1821,
					"points": 5,
					"title": "Cuerpo Humano",
					"short_description": "La tarea consta de realizar una maqueta del cuerpo humano.",
					"long_description" : "La tarea consta de realizar una maqueta del cuerpo humano. Se debe utilizar duroport y plastilina",
					"classroom_id": 3,
					"classroom_name" : "Ciencias Naturales",
					"teachers_id" : "1",
					"teachers_name" : "Odalia Ruiz",
					"delivery_date": "12 de Abril",
					"delivery_hour": "11:00 AM"
				}
	}
		`)

	muckStruct := HomeworkDetailResponse{}

	if err := json.Unmarshal(muckData, &muckStruct); err != nil {
		panic(err)
	}

	// Set Sesson
	w.Header().Set("SessionId", "MySession")

	apigolang.SuccesContentResponse("Detalle de la tarea", "¡Esta es información de prueba!", muckStruct, w)
	return
}

func getStudentsHandler(w http.ResponseWriter, r *http.Request) {

	request := apigolang.Request{
		HTTPReq: r,
	}

	studentID, response := request.GetURLParamInt64("userID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	log.Println(studentID)

	response = getClassroomsFromStudent(studentID)
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}

	return

}
