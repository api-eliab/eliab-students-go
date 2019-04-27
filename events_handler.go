package main

import (
	"encoding/json"
	"log"
	"net/http"

	apigolang "github.com/josuegiron/api-golang"
)
 
func getEvents(w http.ResponseWriter, r *http.Request) {
	request := apigolang.Request{
		HTTPReq: r,
	}
	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	log.Println(studentID)

	muckData := []byte(`
	{
		"events": [{
			"id": 1,
			"title": "Dia de las Madres",
			"description": "Invitamos este 10 de Mayo a todas las madres...",
			"icon" : 10,
			"need_assisstance" : true,
			"assisstance" : false,
			"need_payment" : true,
			"paid" : false,
			"formatted_date": "10 de Mayo 10:00 AM"
		},
		{
			"id": 2,
			"title": "Reunion de entrega de notas",
			"description": "La entrega de notas se realizará durante toda la mañana,se suspenden clases de los alumnos.",
			"icon" : 11,
			"need_assisstance" : true,
			"assisstance" : true,
			"need_payment" : true,
			"paid" : false,
			"formatted_date": "15 de Mayo 08:00 AM"
		},
		{
			"id": 3,
			"title": "Campamento Escolar",
			"description": "Cada año el colegio tiene el agrado de ...",
			"icon" : 11,
			"need_assisstance" : true,
			"assisstance" : false,
			"need_payment" : true,
			"paid" : false,
			"formatted_date": "20 de Julio 08:00 AM"
		},
		{
			"id": 4,
			"title": "Reunion de entrega de notas finales",
			"description": "La entrega de notas se realizará durante toda la mañana,se suspenden clases de los alumnos.",
			"icon" : 11,
			"need_assisstance" : true,
			"assisstance" : true,
			"need_payment" : true,
			"paid" : false,
			"formatted_date": "15 de Mayo 08:00 AM"
		}
	]
	}
	`)

	muckStruct := EventsResponse{}

	if err := json.Unmarshal(muckData, &muckStruct); err != nil {
		panic(err)
	}

	// Set Sesson
	w.Header().Set("SessionId", "MySession")

	apigolang.SuccesContentResponse("Tareas del estudiante", "¡Esta es información de prueba!", muckStruct, w)
	return
}

func confirmEventAssistant(w http.ResponseWriter, r *http.Request) {
	var eventRequest EventConfirmAssitRequest
	request := apigolang.Request{
		HTTPReq: r,
		JSONStruct: &eventRequest,
	}

	responseUM:= request.UnmarshalBody()
	if responseUM != nil {
		apigolang.SendResponse(responseUM, w)
		return
	}

	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	eventID, responsee := request.GetURLParamInt64("eventID")
	if responsee != nil {
		apigolang.SendResponse(responsee, w)
		return
	}

	log.Println(studentID)
	log.Println(eventID)

	// Set Sesson
	w.Header().Set("SessionId", "MySession")
	if eventRequest.Confirmed {
		apigolang.SuccesResponse("¡Esta es información de prueba!", "Asistiras al evento", w)
		return
	}
	
	apigolang.SuccesResponse("¡Esta es información de prueba!", "Declinaste el evento", w)
	return
}