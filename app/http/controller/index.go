package controller

import (
	"gin-admin/app/utility/response"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Context(c).View("index")
	return
}
func Dashboard(c *gin.Context) {
	response.Context(c).View("dashboard")
	return
}
