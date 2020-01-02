package controller

import (
	"DataBase/connect"
	"DataBase/model"
)

var data string

func FindeOne() (err error) {
	//func (t *Test) FindeOne(c *gin.Context) {
	db := connect.DB
	defer db.Close()

	var porgs []model.Tests
	result := db.First(&porgs)

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
