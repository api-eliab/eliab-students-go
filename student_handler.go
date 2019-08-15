package main

import (
	"net/http"

	apigo "github.com/josuegiron/api-golang"
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
	iconID, responseh := request.GetURLParamInt64("iconID")
	if responseh != nil {
		apigolang.SendResponse(responseh, w)
		return
	}

	response = setIcon(studentID, iconID)
	if response != nil {
		apigo.SendResponse(response, w)
		return
	}
	return

} // setIconHandler
