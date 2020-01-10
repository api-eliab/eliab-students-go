package main

import (
	"net/http"

	apigolang "github.com/josuegiron/api-golang"
)

func setIconHandler(w http.ResponseWriter, r *http.Request) {

	request := apigolang.Request{
		HTTPReq: r,
	}

	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	iconID, response := request.GetURLParamInt64("iconID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	response = setIcon(studentID, iconID)
	apigolang.SendResponse(response, w)
	return

}
