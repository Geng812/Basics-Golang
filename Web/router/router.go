package router

import (
	"net/http"

	"DataBase/controller"

	"github.com/gin-gonic/gin"
)

func Init() {

  r := gin.Default()
  
  r.GET("/", Home)

  page := r.Group("/report")
  {
    page.GET("\daily", ReportDaily)
  }
  
  r.RUN(":8080")
}
