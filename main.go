package main

import (
	"embed"
	"fmt"
	"ginProject/src/systemInit"
	"github.com/gin-gonic/gin"
)

//go:embed src/*.html dist/* static/*
var FS embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	systemInit.SourceInit(r, FS)
	systemInit.SetRoute(r)
	addr := ":8080"
	fmt.Printf("启动服务" + addr)
	err := r.Run(addr)
	if err != nil {
		println("服务启动失败，检查端口是否被占用")
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
