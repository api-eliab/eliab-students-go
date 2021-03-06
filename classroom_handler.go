package main

import (
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

	response = getClassroomsFromStudent(studentID)
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}
	return
}

func getClassroomDetailHandler(w http.ResponseWriter, r *http.Request) {

	request := apigolang.Request{
		HTTPReq: r,
	}
	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}
	classroomID, response := request.GetURLParamInt64("classroomID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	response = getClassroomDetail(studentID, classroomID)

	apigo.SendResponse(response, w)
	return

}
