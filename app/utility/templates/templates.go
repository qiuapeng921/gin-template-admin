package templates

import (
	"gin-admin/app/utility/system"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
)

func InitTemplate(router *gin.Engine) {
	router.Static("/assets", "./public/assets")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.HTMLRender = loadTemplates("./templates")
}

//多模板（模板继承）
func loadTemplates(templatesDir string) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	var (
		layouts, pages, subPages, newPages []string
		err                                error
	)

	// 加载后台模板
	layouts, err = filepath.Glob(templatesDir + "/layouts/backend_base.html")
	system.SecurePanic(err)

	pages, err = filepath.Glob(templatesDir + "/backend/*.html")
	system.SecurePanic(err)

	subPages, err = filepath.Glob(templatesDir + "/backend/**/*.html")
	system.SecurePanic(err)

	if len(subPages) > 0 {
		for _, page := range subPages {
			newPages = append(pages, page)
		}
	} else {
		newPages = pages
	}

	for _, page := range newPages {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, page)
		newPage := strings.Replace(filepath.ToSlash(page), "templates/backend/", "", -1)
		renderer.AddFromFilesFuncs(newPage, FuncMap(), files...)
	}
	return renderer
}
