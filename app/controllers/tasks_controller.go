package controllers

import (
	"fmt"
	"log"
	"mypackage/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	Db = sqlConnect()
	Db.AutoMigrate(&models.Task{})
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	Db.Find(&tasks)
	fmt.Println("*********GetTasks**Start***********")
	fmt.Println(tasks)
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	fmt.Println("*********GetTask**Start***********")
	// task := []models.Task{}
	var task models.Task
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
	Db.First(&task, id)
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	fmt.Println("*********CreateTask**Start***********")
	var req models.Task
	fmt.Println(req)
	c.BindJSON(&req)
	fmt.Println(req)
	fmt.Println(req.Title)

	task := &models.Task{Title: req.Title}
	Db.Create(task)

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var req models.Task
	c.BindJSON(&req)
	var task models.Task
	Db.First(&task, req.ID)
	task.Title = req.Title
	Db.Save(&task)
	c.JSON(http.StatusOK, task)
}
