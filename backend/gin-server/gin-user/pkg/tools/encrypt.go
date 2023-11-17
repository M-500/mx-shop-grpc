// Package tools
// Date        : 2023/2/15 20:49
// Version     : 1.0.0
// Author      : 代码小学生王木木
// Email       : 18574945291@163.com
// Description :
package tools

import (
	"crypto/md5"
	"fmt"
	"io"
)

/** 加密方式 **/

func Md5ByString(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(any(err))
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
