package route

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
