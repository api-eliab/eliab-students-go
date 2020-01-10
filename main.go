package main

import (
	"net/http"

	"github.com/gorilla/mux"
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func main() {

	loadConfiguration()
	router := mux.NewRouter()

	if !dbConnect() {
		log.Panic("Error al conectar a la base de datos!")
	}

	middlewares := apigo.MiddlewaresChain(apigo.BasicAuth, apigo.RequestHeaderJson, apigo.GetRequestBodyMiddleware)

	router.HandleFunc("/v1.0/student/{studentID}/homeworks", middlewares(getHomeworksHandler)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/homework/{homeworkID}", middlewares(getHomeworkDetail)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/events", middlewares(getEvents)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/event/{eventID}/confirm_assisstant", middlewares(confirmEventAssistant)).Methods("POST")
	router.HandleFunc("/v1.0/student/{studentID}/classrooms", middlewares(getClassrooms)).Methods("Get")
	router.HandleFunc("/v1.0/student/{studentID}/classroom/{classroomID}", middlewares(getClassroomDetailHandler)).Methods("Get")
	router.HandleFunc("/v1.0/students", middlewares(getStudentsHandler)).Methods("Get")
	router.HandleFunc("/v1.0/student/{studentID}/icon/{iconID}", middlewares(setIconHandler)).Methods("POST")

	log.Println("Starting server on port ", config.General.ServerAddress)
	if startServerError := http.ListenAndServe(config.General.ServerAddress, router); startServerError != nil {
		panic(startServerError)
	}
}
