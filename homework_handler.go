package main

import (
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
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}
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

	response = getHomeworkDetail(homeworkID, studentID)
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}
	return
}
