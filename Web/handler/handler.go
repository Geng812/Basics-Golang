package handler

import (
	"DataBase/controller"

	"github.com/gin-gonic/gin"
)

func Messages(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "Handler-Test1",
	//})
	c.HTML(200, "index.html", nil)
}

func GetData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": controller.FindeOne,
	})
	//c.HTML(200, "index.html", nil)
}
