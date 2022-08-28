package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string
	Source []byte `gorm:"size:70000"`
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

func GetAllTTodos() []Todo {
	var todos []Todo
	db := sqlConnect()
	db.Find(&todos)
	return todos

}
