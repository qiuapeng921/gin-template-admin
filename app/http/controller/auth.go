package controller

import (
	"fmt"
	"gin-admin/app/service"
	"gin-admin/app/utility/app"
	"gin-admin/app/utility/captcha"
	"gin-admin/app/utility/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authRequestData struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Code     string `json:"code" form:"code" binding:"required"`
}

// @生成验证码
// @Author 邱阿朋
// @Date 16:43 2020/11/04
func Captcha(ctx *gin.Context) {
	captcha.GenerateCaptcha(ctx)
	return
}

// @用户登录
// @Author 邱阿朋
// @Date 16:43 2020/11/04
func Login(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		response.Context(ctx).View("login")
		return
	}

	var request authRequestData
	if err := ctx.ShouldBind(&request); err != nil {
		response.Context(ctx).Error(10000, err.Error())
		return
	}
	if !captcha.VerifyCaptcha(ctx, request.Code) {
		response.Context(ctx).Error(10001, "验证码错误")
		return
	}
	result, code, err := service.HandelAdminAuth(request.Username, request.Password)
	if err != nil {
		response.Context(ctx).Error(code, err.Error())
	} else {
		response.Context(ctx).Success(result)
	}
	return
}

// @Description
// @Author 邱阿朋
// @Date 16:43 2020/11/04
func Logout(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	app.Redis().Del(fmt.Sprintf("token:%d", id))
	response.Context(ctx).Success()
	return
}
