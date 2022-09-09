package controller

import (
	"ginProject/src/handler"
	"github.com/gin-gonic/gin"
)

func EditAddress(context *gin.Context) {
	address := context.PostForm("editAddress")
	name := context.PostForm("editName")
	err := handler.FromAddressSaveName(address, name)
	if err != nil {
		return
	}
	context.Redirect(301, "./index")
	context.Done()
}
func DeleteAddress(context *gin.Context) {
	address := context.PostForm("editAddress")
	handler.DeleteAddress(address)
	context.Redirect(301, "./index")
	context.Done()
}
