//
// @Author: 18574
// @Description:
// @File:  jsonResult
// @Version: 1.0.0
// @Date: 2023/2/15 23:02
//

package tools

import (
	e "gin-user/pkg/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResult struct {
	ErrorCode int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func Json(code int, message string, data interface{}, success bool) *JsonResult {
	return &JsonResult{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   success,
	}
}

// JsonSuccessData
// @Description: 成功时返回
// @param ctx:
// @param msg:
// @param data:
func JsonSuccessData(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, JsonResult{
		ErrorCode: e.SUCCESS,
		Data:      data,
		Message:   msg,
		Success:   true,
	})
}

func JsonErrorResp(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, JsonResult{
		ErrorCode: e.ERROR,
		Data:      nil,
		Message:   err.Error(),
		Success:   false,
	})
}

func JsonErrorInterfaceResp(ctx *gin.Context, err interface{}) {
	ctx.JSON(http.StatusOK, JsonResult{
		ErrorCode: e.ERROR,
		Data:      err,
		Message:   "失败",
		Success:   false,
	})
}

func JsonErrorCodeResp(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, JsonResult{
		ErrorCode: code,
		Data:      nil,
		Message:   e.GetMessage(code),
		Success:   false,
	})
}
func JsonErrorStrResp(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusOK, JsonResult{
		ErrorCode: e.ERROR,
		Data:      nil,
		Message:   err,
		Success:   false,
	})
}
