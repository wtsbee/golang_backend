package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"mypackage/app/route"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

var err error

type Memo struct {
	gorm.Model
	Title string `json:"title`
}

func init() {
	Db = sqlConnect()
}

func GetMemos(c *gin.Context) {
	Db.AutoMigrate(&Memo{})
	var results []Memo
	Db.Order("created_at asc").Find(&results)

	memos := []Memo{}
	for _, v := range results {
		memos = append(memos, v)
	}

	// c.JSON(http.StatusOK, gin.H{"memos": memos})
	c.JSON(http.StatusOK, memos)
}

func GetMemo(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("id is not a number")
	}
	var memo Memo
	err1 := Db.First(&memo, id).Error
	if err1 != nil {
		c.JSON(http.StatusNotFound, "Not Found")
	} else {
		c.JSON(http.StatusOK, memo)
	}
}

func CreateMemo(c *gin.Context) {
	var req Memo
	c.BindJSON(&req)
	fmt.Println(req)

	memo := &Memo{Title: req.Title}
	Db.Create(memo)

	c.JSON(200, memo)
}

func DeleteMemo(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("id is not a number")
	}

	var memo Memo
	err1 := Db.First(&memo, id).Error
	if err1 != nil {
		c.JSON(http.StatusNotFound, "Not Found")
	} else {
		Db.Delete(&memo)
		c.JSON(http.StatusOK, memo)
	}
}

func EditMemo(c *gin.Context) {
	// b, err := c.GetRawData()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(b)

	n := c.PostForm("id")
	title := c.PostForm("title")
	// fmt.Println(n, title)
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("id is not a number")
		// fmt.Println("エラー")
	}

	var memo Memo
	err1 := Db.First(&memo, id).Error
	if err1 != nil {
		c.JSON(http.StatusNotFound, "Not Found")
	} else {
		memo.Title = title
		Db.Save(&memo)
		c.JSON(http.StatusOK, memo)
	}
}

func MemosController() {
	fmt.Println("++++++++++++++++++")
	Db.AutoMigrate(&Memo{})
	router := route.RouterCommon()

	router.GET("/memos", func(c *gin.Context) {

		var results []Memo
		Db.Order("created_at asc").Find(&results)

		memos := []Memo{}
		for _, v := range results {
			memos = append(memos, v)
		}

		// c.JSON(http.StatusOK, gin.H{"memos": memos})
		c.JSON(http.StatusOK, memos)
	})

	router.GET("/Memos/:id", func(c *gin.Context) {
		// db := sqlConnect()
		// defer db_v2.Close()

		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("id is not a number")
		}
		var memo Memo
		err1 := Db.First(&memo, id).Error
		if err1 != nil {
			c.JSON(http.StatusNotFound, "Not Found")
		} else {
			c.JSON(http.StatusOK, memo)
		}
	})

	router.POST("/memos", func(c *gin.Context) {
		// db := sqlConnect()
		// defer db_v2.Close()

		// var form map[string]interface{}
		// c.BindJSON(&form)
		// fmt.Println(form)

		fmt.Println("***********")
		var req Memo
		c.BindJSON(&req)
		fmt.Println(req)

		memo := &Memo{Title: req.Title}
		Db.Create(memo)

		c.JSON(200, memo)
	})

	router.DELETE("/memos/:id", func(c *gin.Context) {
		// db := sqlConnect()
		// defer db_v2.Close()

		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("id is not a number")
		}

		var memo Memo
		err1 := Db.First(&memo, id).Error
		if err1 != nil {
			c.JSON(http.StatusNotFound, "Not Found")
		} else {
			Db.Delete(&memo)
			c.JSON(http.StatusOK, memo)
		}
	})

	router.POST("/memos/:id", func(c *gin.Context) {
		// db := sqlConnect()
		// defer db_v2.Close()
		fmt.Println("+++++++++++++++++++++")
		// n := c.Param("id")
		n := c.PostForm("id")
		fmt.Println(n)
		title := c.PostForm("title")
		fmt.Println(title)
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("id is not a number")
		}

		var memo Memo
		err1 := Db.First(&memo, id).Error
		if err1 != nil {
			c.JSON(http.StatusNotFound, "Not Found")
		} else {
			fmt.Println("更新")
			// title := c.PostForm("title")
			// req := c.Request
			// fmt.Println(req)
			// req.ParseForm()
			// fmt.Println(req.Method)
			// fmt.Println(req.Form)
			// fmt.Println(title)
			fmt.Println("更新aa")
			memo.Title = title
			Db.Save(&memo)
			c.JSON(http.StatusOK, memo)
		}
	})

	// db.DropTable("Memos")
	// Db.Migrator().DropTable("memos")
	router.Run()
}

func sqlConnect() (database *gorm.DB) {
	// DBMS := "mysql"
	USER := "usr"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "golang_todo_app_db"

	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	// db, err := gorm.Open(DBMS, CONNECT)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
			// db, err = gorm.Open(DBMS, CONNECT)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		}
	}

	return db
}
