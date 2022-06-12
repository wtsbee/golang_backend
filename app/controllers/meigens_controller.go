package controllers

import (
	"net/http"

	"mypackage/app/route"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var Db *gorm.DB

// var err error

type Meigen struct {
	gorm.Model
	Meigen string
}

func init() {
	Db = sqlConnect()
}

func MeigensController() {
	Db.AutoMigrate(&Meigen{})
	router := route.RouterCommon()

	router.GET("/meigens", func(c *gin.Context) {

		var results []Meigen
		Db.Order("created_at asc").Find(&results)

		meigens := []Meigen{}
		for _, v := range results {
			meigens = append(meigens, v)
		}

		c.JSON(http.StatusOK, gin.H{"meigens": meigens})
	})

	// router.GET("/meigens/:id", func(c *gin.Context) {
	// 	// db := sqlConnect()
	// 	// defer db_v2.Close()

	// 	n := c.Param("id")
	// 	id, err := strconv.Atoi(n)
	// 	if err != nil {
	// 		panic("id is not a number")
	// 	}
	// 	var meigen Meigen
	// 	err1 := Db.First(&meigen, id).Error
	// 	if err1 != nil {
	// 		c.JSON(http.StatusNotFound, "Not Found")
	// 	} else {
	// 		c.JSON(http.StatusOK, meigen)
	// 	}
	// })

	// router.POST("/meigens", func(c *gin.Context) {
	// 	// db := sqlConnect()
	// 	// defer db_v2.Close()

	// 	// var form map[string]interface{}
	// 	// c.BindJSON(&form)
	// 	// fmt.Println(form)

	// 	var req Meigen
	// 	c.BindJSON(&req)
	// 	fmt.Println(req)

	// 	meigen := &Meigen{Meigen: req.Meigen}
	// 	Db.Create(meigen)

	// 	c.JSON(200, meigen)
	// })

	// router.DELETE("/meigens/:id", func(c *gin.Context) {
	// 	// db := sqlConnect()
	// 	// defer db_v2.Close()

	// 	n := c.Param("id")
	// 	id, err := strconv.Atoi(n)
	// 	if err != nil {
	// 		panic("id is not a number")
	// 	}

	// 	var meigen Meigen
	// 	err1 := Db.First(&meigen, id).Error
	// 	if err1 != nil {
	// 		c.JSON(http.StatusNotFound, "Not Found")
	// 	} else {
	// 		Db.Delete(&meigen)
	// 		c.JSON(http.StatusOK, meigen)
	// 	}
	// })

	// // db.DropTable("meigens")
	router.Run()
}

// func sqlConnect() (database *gorm.DB) {
// 	// DBMS := "mysql"
// 	USER := "usr"
// 	PASS := "password"
// 	PROTOCOL := "tcp(db:3306)"
// 	DBNAME := "golang_todo_app_db"

// 	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
// 	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

// 	count := 0
// 	// db, err := gorm.Open(DBMS, CONNECT)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		for {
// 			if err == nil {
// 				fmt.Println("")
// 				break
// 			}
// 			fmt.Print(".")
// 			time.Sleep(time.Second)
// 			count++
// 			if count > 180 {
// 				fmt.Println("")
// 				panic(err)
// 			}
// 			// db, err = gorm.Open(DBMS, CONNECT)
// 			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 		}
// 	}

// 	return db
// }
