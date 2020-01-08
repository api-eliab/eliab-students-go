package main

import (
	"encoding/json"
	"log"
	"net/http"

	apigo "github.com/josuegiron/api-golang"
	apigolang "github.com/josuegiron/api-golang"
)

func getHomeworks(w http.ResponseWriter, r *http.Request) {
	request := apigolang.Request{
		HTTPReq: r,
	}
	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
	}

	log.Println(studentID)

	muckData := []byte(`
	{
		"homeworks":[
			{
				"id":1,
				"points": 5,
				"title": "El cuerpo humano",
				"short_description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
				"long_description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				"classroom_id": "123",
				"delivery_date": "12 de abril",
				"delivery_hour": "11:00AM"
			}, {
				"id":2,
				"points": 15,
				"title": "El Universo",
				"short_description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
				"long_description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				"classroom_id": "124",
				"delivery_date": "12 de abril",
				"delivery_hour": "15:00AM"
			}, {
				"id":3,
				"points": 10,
				"title": "Ejercicios pagina 20",
				"short_description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
				"long_description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				"classroom_id": "133",
				"delivery_date": "24 de abril",
				"delivery_hour": "10:00AM"
			}
			
		]
	}`)

	muckStruct := HomeworksResponse{}

	if err := json.Unmarshal(muckData, &muckStruct); err != nil {
		panic(err)
	}

	// Set Sesson
	w.Header().Set("SessionId", "MySession")

	apigolang.SuccesContentResponse("Tareas del estudiante", "¡Esta es información de prueba!", muckStruct, w)
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
