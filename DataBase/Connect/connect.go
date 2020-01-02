package connect

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err := gorm.Open("mysql", "account:password@/WebT?charset=utf8")
	//defer DB.Close()

	if err != nil {
		panic("failed to connect database !")
	}

	if DB.Error != nil {
		panic("database error ")
	}
	//return DB
}
