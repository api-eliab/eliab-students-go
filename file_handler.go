package main

import (
	"net/http"

	apigo "github.com/josuegiron/api-golang"
	apigolang "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {

	var fileInfo FileInfo

	request := apigolang.Request{
		HTTPReq: r,
	}

	ownerID, response := request.GetURLParamInt64("ownerID")
	if response != nil {
		apigolang.SendResponse(response, w)
		return
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

	fileInfo.Extension = request.HTTPReq.FormValue("extension")
	fileInfo.Description = request.HTTPReq.FormValue("description")

	var err error

	fileInfo.File, fileInfo.FileHeader, err = request.HTTPReq.FormFile("file")

	if err != nil {
		log.Error(err)
		apigolang.SendResponse(
			apigo.Error{
				Title:   "¡Error al obtener el archivo!",
				Message: "Por favor intenta de nuevo",
			},
			w,
		)
		return
	}

	response = uploadFile(ownerID, studentID, classroomID, fileInfo)
	apigolang.SendResponse(response, w)
	return

}
