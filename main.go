package main

import (
	"embed"
	"flag"
	"fmt"
	"ginProject/src/systemInit"
	"github.com/gin-gonic/gin"
)

//go:embed src/*.html dist/* static/*
var FS embed.FS

func main() {
	var ip string
	var port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip地址，默认为127.0.0.1")
	flag.StringVar(&port, "port", "8080", "端口，默认为8080")
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	systemInit.SourceInit(r, FS)
	systemInit.SetRoute(r)
	addr := ip + ":" + port
	fmt.Printf("启动服务" + addr)
	err := r.Run(addr)
	if err != nil {
		println("服务启动失败，检查端口是否被占用")
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
