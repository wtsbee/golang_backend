package controllers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterCommon() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))
	return router
}

func GetRouter() *gin.Engine {
	r := RouterCommon()
	// r := gin.Default()
	r.GET("/meigens", GetMeigens)
	r.GET("/memos", GetMemos)
	r.GET("/memo/:id", GetMemo)
	r.POST("/memos", CreateMemo)
	r.DELETE("/memos/:id", DeleteMemo)
	r.POST("/memos/:id", EditMemo)

	r.GET("/todos", GetTodos)
	r.GET("/todo/:id", GetTodo)
	r.POST("/todos", CreateTodo)
	r.DELETE("/todo/:id", DeleteTodo)
	r.POST("/todo/:id", UpdateTodo)
	r.POST("/todo/upload_file", UploadFile)

	return r
}
