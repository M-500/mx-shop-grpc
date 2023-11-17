//
// @Author: 18574
// @Description:
// @File:  validat
// @Version: 1.0.0
// @Date: 2023/2/15 23:07
//

package validators

import (
	"errors"
	"gin-user/pkg/tools/str"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidatePhone(f1 validator.FieldLevel) bool {
	phone := f1.Field().String()
	ok, _ := regexp.MatchString("^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$", phone)
	if ok {
		return true
	}
	return false
}

func IsPassword(password, rePassword string) error {
	if str.IsBlank(password) {
		return errors.New("请输入密码")
	}
	if str.RuneLen(password) < 6 {
		return errors.New("密码过于简单")
	}
	if password != rePassword {
		return errors.New("两次输入密码不匹配")
	}
	return nil
}
