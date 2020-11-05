package response

import (
	"fmt"
	"gin-admin/app/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Wrapper struct {
	*gin.Context
}

func Context(c *gin.Context) *Wrapper {
	return &Wrapper{c}
}

type Response struct {
	Code    int         `json:"code"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (wrapper *Wrapper) View(name string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	wrapper.HTML(http.StatusOK, fmt.Sprintf("%s.html", name), responseData)
	wrapper.Abort()
	return
}

func (wrapper *Wrapper) Success(data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	wrapper.JSON(http.StatusOK, Response{
		Code:    0,
		Status:  http.StatusOK,
		Message: "Success",
		Data:    responseData,
	})
	wrapper.Abort()
	return
}

func (wrapper *Wrapper) Error(errCode int, message ...string) {
	responseMessage := consts.GetMsg(errCode)
	if len(message) > 0 {
		responseMessage = message[0]
	}
	wrapper.JSON(http.StatusOK, Response{
		Code:    0,
		Status:  errCode,
		Message: responseMessage,
	})
	wrapper.Abort()
	return
}
