package systemInit

import (
	"ginProject/src/controller"
	"github.com/gin-gonic/gin"
)

func SetRoute(engine *gin.Engine) {
	engine.GET("/", controller.GetScores)
	engine.POST("/insertAddress", controller.SaveAddressGetScores)
	engine.POST("/editName", controller.EditAddressName)
	engine.GET("/index", controller.GetScores)
}
