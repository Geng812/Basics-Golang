package tmp

import (
	"Web/handler"

	"github.com/gin-gonic/gin"
)

func Table() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("/Users/nick/go/src/templates/html/*")

	homeRouting := r.Group("/")
	{
		homeRouting.GET("", handler.Messages)
	}

	dbRouting := r.Group("/db")
	{
		dbRouting.GET("", handler.GetData)
	}

	return r
}
