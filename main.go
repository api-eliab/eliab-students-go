package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	apigolang "github.com/josuegiron/api-golang"
)

func main() {
	LoadConfiguration()
	router := mux.NewRouter()

	middlewares := apigolang.MiddlewaresChain(apigolang.BasicAuth, apigolang.RequestHeaderJson, apigolang.GetRequestBodyMiddleware)

	router.HandleFunc("/v1.0/student/{studentID}/homeworks", middlewares(getHomeworks)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/homework/{homeworkID}", middlewares(getHomeworkDetail)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/events", middlewares(getEvents)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/event/{eventID}/confirm_assisstant", middlewares(confirmEventAssistant)).Methods("POST")
	router.HandleFunc("/v1.0/student/{studentID}/classrooms", middlewares(getClassrooms)).Methods("Get")
	router.HandleFunc("/v1.0/student/{studentID}/classroom/{classroomID}", middlewares(getClassroomDetail)).Methods("Get")

	log.Println("Starting server on port ", config.General.ServerAddress)
	if startServerError := http.ListenAndServe(config.General.ServerAddress, router); startServerError != nil {
		panic(startServerError)
	}
}
