package dao

import "github.com/dgrijalva/jwt-go"

// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 18:18
type CustomClaims struct {
	ID          uint
	NickName    string
	UserName    string
	AuthorityId uint
	jwt.StandardClaims
}
