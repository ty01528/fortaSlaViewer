package systemInit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRoute(engine *gin.Engine) {
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"user":       "田源",
			"userDetail": "欢迎关注",
		})
	})
}
