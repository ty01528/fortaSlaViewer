package systemInit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(engine *gin.Engine) {
	user := "tianyuan"
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"user":       user,
			"userDetail": "老废物",
		})
	})
}
