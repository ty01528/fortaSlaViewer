package systemInit

import (
	"ginProject/src/controller"
	"github.com/gin-gonic/gin"
)

func SetRoute(engine *gin.Engine) {
	engine.GET("/", controller.GetScores)
	engine.POST("/index", controller.SaveAddressGetScores)
	engine.GET("/index", controller.GetScores)
}
