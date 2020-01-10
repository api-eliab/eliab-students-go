package main

import (
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

// import (
// 	"fmt"

// 	apigo "github.com/josuegiron/api-golang"
// 	"github.com/josuegiron/log"
// )

// func getStudents() {
// 	user.Sons, err = getUserSons(user.ID)
// 	if err != nil {
// 		log.Error(err)
// 		return apigo.Error{
// 			Title:   "Error al consultar la información de los hijos!",
// 			Message: "Error al consultar la información de los hijos!",
// 		}
// 	}

// 	for _, son := range user.Sons {
// 		var newSon ResponseSon
// 		newSon.ID = son.ID
// 		newSon.FirstName = son.FirstName
// 		newSon.LastName = fmt.Sprintf("%v %v", son.FirstLastName, son.SecondLastName)
// 		newSon.Avatar = son.Avatar
// 		respData.User.Sons = append(respData.User.Sons, newSon)
// 	}

// 	return apigo.Success{
// 		Content: respData,
// 	}
// }

func getHomeworks(studentID int64) apigo.Response {

	homeworks, err := getHomeworksDB(studentID)
	if err != nil {
		log.Error(err)
		return apigo.Error{
			Title:   "Error al consultar la información de los cursos del estudiante!",
			Message: "Error al consultar la información de los cursos del estudiante!",
		}
	}

	response := HomeworksResponse{Homeworks: homeworks}
	return apigo.Success{
		Content: response,
	}

}
