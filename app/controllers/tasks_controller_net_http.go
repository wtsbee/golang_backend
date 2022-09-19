package controllers

import (
	"fmt"
	"mypackage/app/models"
	"net/http"
)

// func StartNetHttp() err {

// }

func GetAllTasksNetHttp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	fmt.Println("********GetAllTasksNetHttp*********")
	var tasks []models.Task
	db := sqlConnect()
	db.Find(&tasks)
	fmt.Fprintln(w, tasks)
}
