package main

import (
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

	expire, _ := request.GetQueryParamInt64("expire")

	response = getHomeworks(studentID, expire)
	apigo.SendResponse(response, w)
	return

}

func getHomeworkDetailHandler(w http.ResponseWriter, r *http.Request) {

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

	response = getHomeworkDetail(studentID, homeworkID)
	apigo.SendResponse(response, w)
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

	apigo.SendResponse(response, w)
	return

}
