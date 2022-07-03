package controllers

import (
	"fmt"
	"io"
	"log"
	"mypackage/app/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string
}

func init() {
	Db = sqlConnect()
	Db.AutoMigrate(&Todo{})
}

func GetTodos(c *gin.Context) {
	// Db.AutoMigrate(&Todo{})
	// var results []Todo
	// Db.Find(&results)

	results := models.GetAllTTodos()

	c.JSON(http.StatusOK, results)
}

func GetTodo(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("id is not a number")
	}
	var todo Todo
	err1 := Db.First(&todo, id).Error
	if err1 != nil {
		c.JSON(http.StatusNotFound, "Not Found")
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c *gin.Context) {
	// Db.AutoMigrate(&Todo{})
	var req Todo
	c.BindJSON(&req)
	fmt.Println(req)

	todo := &Todo{Title: req.Title}
	Db.Create(todo)

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	fmt.Println("****************")
	n := c.Param("id")
	fmt.Printf("%T, %v\n", n, n)
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("id is not a number")
	}

	var todo Todo
	err1 := Db.First(&todo, id).Error
	if err1 != nil {
		c.JSON(http.StatusNotFound, "Not Found")
	} else {
		Db.Delete(&todo)
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	// n := c.PostForm("id")
	// title := c.PostForm("title")
	var req Todo
	err = c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	var todo Todo
	err1 := Db.First(&todo, req.ID).Error
	if err1 != nil {
		c.JSON(http.StatusNotFound, "Not Found")
	} else {
		todo.Title = req.Title
		Db.Save(&todo)
		c.JSON(http.StatusOK, todo)
	}
}

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	fileName := header.Filename
	dir, _ := os.Getwd()
	out, err := os.Create(dir + "/images/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
