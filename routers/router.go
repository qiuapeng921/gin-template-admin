package routers

import (
	"gin-admin/app/http/middleware"
	"gin-admin/app/utility/templates"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(router *gin.Engine) {
	templates.InitTemplate(router)

	router.Use(
		middleware.Cors(),
		middleware.HandleException(),
		middleware.RequestId(),
	)

	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/admin/index")
	})

	// 加载后台路由组
	InitAdminRouter(router)
}
