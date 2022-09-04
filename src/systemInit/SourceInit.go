package systemInit

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net/http"
)

func SourceInit(r *gin.Engine, FS embed.FS) {
	tmpl := template.Must(template.New("").ParseFS(FS, "src/index.html"))
	r.SetHTMLTemplate(tmpl)
	fStatic, _ := fs.Sub(FS, "static")
	fDist, _ := fs.Sub(FS, "dist")
	r.StaticFS("/static", http.FS(fStatic))
	r.StaticFS("/dist", http.FS(fDist))
}
