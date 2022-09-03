package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAddresses(context *gin.Context) {
	addresses := context.PostForm("addresses")
	println(addresses)
	context.HTML(http.StatusOK, "index.html", gin.H{})
}
