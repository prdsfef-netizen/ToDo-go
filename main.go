package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

//Check verficar estado de api
func Check(w http.ResponseWriter, r *http.Request) {
	log.Info("API 200")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/api", Check).Methods("GET")
	http.ListenAndServe(":8000", router)
}
