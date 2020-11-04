package templates

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func InitTemplate(router *gin.Engine) {
	router.Static("/assets", "./public/assets")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.HTMLRender = loadTemplates("./templates")
}

//多模板（模板继承）
func loadTemplates(templatesDir string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	// 加载后台模板
	backendLayouts, err := filepath.Glob(templatesDir + "/layouts/backend_base.html")
	if err != nil {
		panic(err.Error())
	}

	backendPages, err := filepath.Glob(templatesDir + "/backend/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, page := range backendPages {
		layoutCopy := make([]string, len(backendLayouts))
		copy(layoutCopy, backendLayouts)
		files := append(layoutCopy, page)
		renderer.AddFromFilesFuncs(filepath.Base(page), FuncMap(), files...)
	}

	return renderer
}