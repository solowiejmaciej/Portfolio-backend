package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"portfolio-api/pkg/handlers/career"
	"portfolio-api/pkg/handlers/projects"
	"portfolio-api/pkg/middleware"
)

var port = ":5020"

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.TraceLevel)
	log.Infof("Application starting...")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	router.HandleFunc("/projects", projects.GetProjects).Methods("GET")
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}).Methods("GET")
	router.HandleFunc("/career", career.GetCareer).Methods("GET")

	router.Use(middleware.LoggingMiddleware)

	log.Infof("Server started at http://localhost%s", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
