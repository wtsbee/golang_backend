package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"mypackage/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// "gorm.io/gorm"
)

type Meigen struct {
	gorm.Model
	Meigen string
}

func main() {
	fmt.Println("***********************************")
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)
	db := sqlConnect()
	db.AutoMigrate(&Meigen{})
	defer db.Close()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	router.GET("/meigens", func(c *gin.Context) {
		db := sqlConnect()
		defer db.Close()

		var results []Meigen
		db.Order("created_at asc").Find(&results)

		meigens := []Meigen{}
		for _, v := range results {
			meigens = append(meigens, v)
		}

		c.JSON(http.StatusOK, gin.H{"meigens": meigens})
	})

	router.GET("/meigens/:id", func(c *gin.Context) {
		db := sqlConnect()
		defer db.Close()

		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("id is not a number")
		}
		var meigen Meigen
		if db.First(&meigen, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, "Not Found")
		} else {
			c.JSON(http.StatusOK, meigen)
		}
	})

	router.POST("/meigens", func(c *gin.Context) {
		db := sqlConnect()
		defer db.Close()

		var req Meigen
		c.BindJSON(&req)

		meigen := &Meigen{Meigen: req.Meigen}
		db.Create(meigen)

		c.JSON(200, meigen)
	})

	router.DELETE("/meigens/:id", func(c *gin.Context) {
		db := sqlConnect()
		defer db.Close()

		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("id is not a number")
		}

		var meigen Meigen
		if db.First(&meigen, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, "Not Found")
		} else {
			db.Delete(&meigen)
			c.JSON(http.StatusOK, meigen)
		}
	})

	// db.DropTable("meigens")
	router.Run()
}

func sqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "usr"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "golang_todo_app_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}

	return db
}
