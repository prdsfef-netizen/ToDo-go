package main

import (
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var db, _ = gorm.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

type TodoItemModel struct {
	Id int `gorm:"primary_key"`
	Description string
	Completed bool
}

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
	defer db.Close()
	db.Debug().DropTableIfExists(&TodoItemModel{})
	db.Debug().AutoMigrate(&TodoItemModel{})
	
	log.Info("Starting Todolist API server")
	router := mux.NewRouter()
	router.HandleFunc("/api", Check).Methods("GET")
	http.ListenAndServe(":8000", router)
}
