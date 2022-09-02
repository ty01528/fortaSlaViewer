package systemInit

import "github.com/gin-gonic/gin"

func SourceInit(r *gin.Engine) {
	r.Static("/static", "./static")
	r.Static("/dist", "./dist")
	r.LoadHTMLGlob("src/*.html")
}
