package test

import (
	"fmt"
	"go-chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//迁移schema
	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{}
	user.Name = "张洁"
	db.Create(user)

	fmt.Println(db.First(user, 1))

	db.Model(user).Update("Password", "1234")
}
