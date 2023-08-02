package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Name string
	Pass string
}

func Login(c *gin.Context) {
	name := c.Query("name")
	pass := c.Query("pass")
	user := User{Name: name, Pass: pass}

	// めんどいからここで全部処理
	dsn := "host=db user=postgres pass=postgres dbname=postgres	port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// userがいるかを検索
	result := db.Find(&user)
	if result != nil {
		// いる   -> login success
		return
	} else {
		// いない -> login failed
		return
	}
}
