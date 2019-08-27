package main

import (
	"net/http"
	"database/sql"

	"github.com/gorilla/mux"
	apigo "github.com/josuegiron/api-golang"
	"github.com/josuegiron/log"
)

func main() {

	loadConfiguration()
	router := mux.NewRouter()

	connect()


	middlewares := apigo.MiddlewaresChain(apigo.BasicAuth, apigo.RequestHeaderJson, apigo.GetRequestBodyMiddleware)

	router.HandleFunc("/v1.0/schools/{schoolID}/student/{studentID}/homeworks", middlewares(getHomeworksHandler)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/homework/{homeworkID}", middlewares(getHomeworkDetailHandler)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/events", middlewares(getEvents)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/event/{eventID}/confirm_assisstant", middlewares(confirmEventAssistant)).Methods("POST")
	router.HandleFunc("/v1.0/student/{studentID}/classrooms", middlewares(getClassrooms)).Methods("GET")
	router.HandleFunc("/v1.0/student/{studentID}/classroom/{classroomID}", middlewares(getClassroomDetail)).Methods("GET")

	router.HandleFunc("/v1.0/student/{studentID}/icon/{iconID}", middlewares(setIconHandler)).Methods("POST")

	log.Println("Starting server on port ", config.General.ServerAddress)

	apigo.Check(http.ListenAndServe(config.General.ServerAddress, router))

}// end main


func connect(){
	
	catalog = make(map[string]*sql.DB)

	for dbid, _ := range config.Databases {
		if !dbConnect(dbid) {
			log.Info("Error al conectar a la base de datos!")
		}
	}

	log.Info(catalog)
	
}
