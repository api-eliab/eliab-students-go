package main

import (
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func setIcon(studentID, iconID int64) apigo.Response {

	err := setIconDB(studentID, iconID)
	if err != nil {

		log.Error(err)
		return apigo.Error{
			Title:   "Error al cambiar el icono del estudiante!",
			Message: "Error al cambiar el icono del estudiante!",
		}

	}

	return apigo.Success{
		Title:   "Icono cambiado satisfactoriamente",
		Message: "Has cambiado tu icono!",
	}

} // setIcon
