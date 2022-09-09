package controller

import (
	"ginProject/src/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EditAddressName(context *gin.Context) {
	address := context.PostForm("editAddress")
	name := context.PostForm("editName")
	err := handler.FromAddressSaveName(address, name)
	if err != nil {
		return
	}
	scores := handler.GetScoreFromLeveldb()
	println(scores)
	context.HTML(http.StatusOK, "index.html", gin.H{
		"user":       "田源",
		"userDetail": "欢迎关注",
		"Addresses":  scores,
	})
	context.Done()
}
