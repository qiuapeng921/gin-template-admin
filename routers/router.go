package routers

import (
	"gin-admin/app/http/middleware"
	"gin-admin/app/utility/response"
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

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/index")
	})

	router.NoRoute(func(ctx *gin.Context) {
		response.Context(ctx).View("error", gin.H{"message": "路由异常"})
		return
	})

	// 加载后台路由组
	InitAdminRouter(router)
}
