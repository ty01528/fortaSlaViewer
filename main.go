package main

import (
	"ginProject/src/systemInit"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	systemInit.SourceInit(r)
	systemInit.SetRoute(r)
	err := r.Run()
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
