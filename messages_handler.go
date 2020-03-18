package main

import (
	"net/http"

	apigolang "github.com/josuegiron/api-golang"
)

func getNotificationsHandler(w http.ResponseWriter, r *http.Request) {

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

	response = getNotifications(studentID, ownerID)
	apigolang.SendResponse(response, w)
	return

}

func getAnnouncementsHandler(w http.ResponseWriter, r *http.Request) {

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

	response = getAnnouncements(studentID, ownerID)
	apigolang.SendResponse(response, w)
	return

}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {

	var message MessageRequest

	request := apigolang.Request{
		HTTPReq:    r,
		JSONStruct: &message,
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

	classroomID, response := request.GetURLParamInt64("classroomID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	sectionID, response := request.GetURLParamInt64("sectionID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	response = request.UnmarshalBody()
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	response = sendMessage(studentID, ownerID, classroomID, sectionID, message.Message)
	apigolang.SendResponse(response, w)
	return

}

func getConversationHandler(w http.ResponseWriter, r *http.Request) {

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

	classroomID, response := request.GetURLParamInt64("classroomID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	sectionID, response := request.GetURLParamInt64("sectionID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
	}

	response = getMessages(studentID, ownerID, classroomID, sectionID)
	apigolang.SendResponse(response, w)
	return

}
