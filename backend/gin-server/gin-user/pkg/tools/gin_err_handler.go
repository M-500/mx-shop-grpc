// Package tools
// Date        : 2023/2/16 10:39
// Version     : 1.0.0
// Author      : 代码小学生王木木
// Email       : 18574945291@163.com
// Description :
package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"ludaxf_django/pkg/validators"
	"strings"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		fmt.Println(err)
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		JsonErrorResp(ctx, err)
		return
	}
	JsonErrorInterfaceResp(ctx, removeTopStruct(errs.Translate(validators.InstanceTrans)))
}
