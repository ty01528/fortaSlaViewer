package controller

import (
	"ginProject/src/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveAddressGetScores(context *gin.Context) {
	handler.PhraseAndSaveAddress(context)
	context.Redirect(301, "./index")
	context.Done()
}
func GetScores(context *gin.Context) {
	scores := handler.GetScoreFromLeveldb()
	println(scores)
	context.HTML(http.StatusOK, "index.html", gin.H{
		"user":       "田源",
		"userDetail": "欢迎关注",
		"Addresses":  scores,
	})
	context.Done()
}
