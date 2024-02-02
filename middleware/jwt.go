package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"new_demo/utils"
	"new_demo/utils/errmsg"
	"strings"
	"time"
)

var mySigningKey = []byte(utils.MySigningKey)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetToken 创建token
// GetToken 函数为提供的用户名生成一个 JWT token。
// token 包含用户名和过期时间。
// 它使用 MyCustomClaims 结构体，该结构体扩展了 jwt.RegisteredClaims 结构体。
// 使用的签名密钥是 mySigningKey。
func GetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)

	Claims := MyCustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return tokenString, errmsg.SUCCESS
}

// ParseToken 解析 JWT Token
// 它解析给定的 token 字符串，使用 mySigningKey 作为签名密钥，
// 并返回解析后的 MyCustomClaims 对象和可能的错误。
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("无效的token")
}

// JWTAuthMiddleware 中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			code1 := errmsg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code": code1,
				"msg":  errmsg.GetCode(code1),
			})
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			code2 := errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code2,
				"msg":  errmsg.GetCode(code2),
			})
			c.Abort()
			return
		}
		mc, err := ParseToken(parts[1])
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				code3 := errmsg.ERROR_TOKEN_RUNTIME
				c.JSON(http.StatusOK, gin.H{
					"code": code3,
					"msg":  errmsg.GetCode(code3),
				})
				c.Abort()
				return
			}

		}
		c.Set("username", mc.Username)
		c.Next()
	}
}
