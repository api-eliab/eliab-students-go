package main

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
