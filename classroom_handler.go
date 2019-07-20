package main

import (
	"encoding/json"
	"log"
	"net/http"

	apigo "github.com/josuegiron/api-golang"
	apigolang "github.com/josuegiron/api-golang"
)

func getClassrooms(w http.ResponseWriter, r *http.Request) {
	request := apigolang.Request{
		HTTPReq: r,
	}
	studentID, response := request.GetURLParamInt64("studentID")
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

func getClassroomDetail(w http.ResponseWriter, r *http.Request) {
	request := apigolang.Request{
		HTTPReq: r,
	}
	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}
	classroomID, responsec := request.GetURLParamInt64("classroomID")
	if responsec != nil {
		apigolang.SendResponse(responsec, w)
		return
	}

	log.Println(studentID)
	log.Println(classroomID)

	muckData := []byte(`
	{
		"classroom": {
					"name": "Matematicas",
					"teacher": {
						"id" : 1,
						"name" : "Odalia",
						"last_name" : "Ruiz"
					},
					"grade": "Tercero Primaria",
					"course_dist":[{  
						"perfect" : true,
						"name" : "Primer Bimestre",
						"id" : 4,
						"current_points" : 10,
						"tasks" : [{
							"id" : 1,
							"name" : "Tarea 1",
							"points" : "5/5 pts",
							"type" : 1
						   },
							{
							"id" : 2,
							"name" : "Tarea 2",
							"points" : "5/5 pts",
							"type" : 1
							}
						]
				}]
	}
	}
	`)

	muckStruct := ClassroomDetailResponse{}

	if err := json.Unmarshal(muckData, &muckStruct); err != nil {
		panic(err)
	}

	// Set Sesson
	w.Header().Set("SessionId", "MySession")

	apigolang.SuccesContentResponse("¡Esta es información de prueba!", "Listado de asignaturas", muckStruct, w)
	return
}
