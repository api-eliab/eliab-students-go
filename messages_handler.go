package main

import (
	"net/http"

	apigolang "github.com/josuegiron/api-golang"
)

func getMessagesHandler(w http.ResponseWriter, r *http.Request) {

	request := apigolang.Request{
		HTTPReq: r,
	}

	studentID, response := request.GetURLParamInt64("studentID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	ownerID, response := request.GetURLParamInt64("ownerID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	response = getMessages(studentID, ownerID)
	apigolang.SendResponse(response, w)
	return

}
