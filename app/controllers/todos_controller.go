package controllers

import (
	"bytes"
	"fmt"
	"image/jpeg"
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
	Title  string
	Source []byte `gorm:"size:70000"`
}

type Test struct {
	Source []byte
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
	fmt.Println("****************")
	// fmt.Println(len(results))
	// fmt.Println(results[len(results)-1])
	// fmt.Println(results[len(results)-1].Source)

	// c.JSON(http.StatusOK, gin.H{"Source": results[len(results)-1].Source})
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

	file, err := os.Open("/app/images/himawari.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("*********CreateTodo**Start***********")
	fmt.Println(file)
	fmt.Printf("%T\n", file)
	fmt.Println("*********CreateTodo**End***********")

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", img)
	file.Close()

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}
	imageBytes := buffer.Bytes()

	// ｆ := &File{Source: imageBytes}
	// db.Create(ｆ)
	// db.Close()

	todo := &Todo{Title: req.Title, Source: imageBytes}
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
	fmt.Printf("%T\n", c)
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

	// file, err := os.Open("image.JPG")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()

	// buffer := new(bytes.Buffer)
	// if err := jpeg.Encode(buffer, img, nil); err != nil {
	// 	log.Println("unable to encode image.")
	// }
	// imageBytes := buffer.Bytes()

	// ｆ := &File{Source: imageBytes}
	// db.Create(ｆ)
	// db.Close()

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
