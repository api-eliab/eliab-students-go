package main

import (
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func setIcon(studentID, iconID int64) apigo.Response {

	err := updateIconDB(studentID, iconID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al actualizar el icono del estudiante!",
			Message: "Error al actualizar el icono del estudiante!",
		}
	}

	return apigo.Success{}

}
