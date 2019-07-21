package controller

import (
	"DataBase/connect"
	"DataBase/model"

	"github.com/gin-gonic/gin"
)

type Test struct{}

func Config() *Test {
  return &Test{}
}

func (t *Test) FindOne(c *gin.Context) {
  db := connect.InitDB()
  defer db.Close()
  
  var n model.Test
  
  result := db.First(&n)
  
}
