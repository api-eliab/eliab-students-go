package main

import (
	"bytes"
	"io"
	"mime/multipart"

	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func uploadFile(ownerID, studentID, classroomID int64, fileInfo FileInfo) apigo.Response {

	if fileInfo.FileHeader.Size > 5000000 {
		log.Info(fileInfo.FileHeader.Size)
		return apigo.Error{
			Title:   "¡Tu documento es mayor a 5MB!",
			Message: "Elige un docuento menor a 5MB",
		}
	}

	file, err := convertFileToBytes(fileInfo.File, fileInfo.FileHeader)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "¡Error al guardar el archivo!",
			Message: "Por favor intenta de nuevo",
		}
	}

	courseTaskFile := CourseTaskFile{
		UserID:      ownerID,
		StudentID:   studentID,
		CourseID:    classroomID,
		File:        file,
		Description: fileInfo.Description,
		Extension:   fileInfo.Extension,
	}

	err = saveFileInDB(courseTaskFile)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "¡Error al cargar el archivo!",
			Message: "Por favor intenta de nuevo",
		}
	}

	return apigo.Success{
		Title:   "¡Archivo cargado exitosamente!",
		Message: "¡Has cargado exitosamente el archivo!",
	}

}

func convertFileToBytes(file multipart.File, fileHeader *multipart.FileHeader) (interface{}, error) {

	defer file.Close()

	buf := bytes.NewBuffer(nil)

	_, err := io.Copy(buf, file)
	if err != nil {
		return "", err
	}

	return buf.Bytes(), nil

}
