//@Author: wulinlin
//@Description:
//@File:  jwt
//@Version: 1.0.0
//@Date: 2023/03/09 14:06

package middlewares

import (
	"errors"
	"fmt"
	"gin-user/app/config"
	"gin-user/pkg/tools"
	"gin-user/pkg/tools/str"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"gin-user/app/models/dao"
	e "gin-user/pkg/constant"
	ldconst "gin-user/pkg/constant"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("Token格式错误！")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWT struct {
	SigningKey []byte
}

func NewJwt() *JWT {
	jwtCfg := config.ConfigInstance.Jwt
	return &JWT{
		[]byte(jwtCfg.SecretKey), // 可以设置过期时间
	}
}

func (j *JWT) ParseToken(tokenString string) (*dao.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dao.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*dao.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}

func JwtAuth() gin.HandlerFunc {
	jwtCfg := config.ConfigInstance.Jwt
	return func(c *gin.Context) {
		// 从请求头中获取token
		token := c.Request.Header.Get(jwtCfg.JwtHeaderKey)
		//token := c.Request.Header.Get("")
		// 从url参数中获取token
		fmt.Println(token)
		// 从body中获取token
		if str.IsBlank(token) {
			// 直接报错用户未登录
			tools.JsonErrorCodeResp(c, e.NOT_LOGIN)
			c.Abort()
			return
		}
		j := NewJwt()
		// 解析token信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if err == TokenExpired {
					c.JSON(http.StatusUnauthorized, map[string]string{
						"msg": "授权已过期,请重新登录",
					})
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusUnauthorized, "未登陆")
			c.Abort()
			return
		}
		// 将用户信息插入上下文中
		c.Set("claims", claims)
		c.Set(ldconst.JWT_INFO_KEY, claims.ID)
		c.Next()
	}
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &dao.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*dao.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func (j *JWT) CreateToken(claims dao.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
