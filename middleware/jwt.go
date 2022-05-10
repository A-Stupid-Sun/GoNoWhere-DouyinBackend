package middleware

import (
	"douyin/config"
	"douyin/errno"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JWT struct {
	JwtKey []byte
}

// NewJWT 创建JWT 实例
func NewJWT() *JWT {
	return &JWT{JwtKey: []byte(config.JwtKey)}
}

type MyClaims struct {
	UserID int64
	jwt.StandardClaims
}

var (
	TokenExpired     error = errors.New("token 已过期，请重新登录")
	TokenNotValidYet error = errors.New("token 无效，请重新登录")
	TokenMalFormed   error = errors.New("token 不正确，请重新登录")
	TokenInvalid     error = errors.New("这不是一个Token，请重新登录")
)

// CreateToken 创建token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.JwtKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalFormed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}

	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenExpired
	}

	return nil, TokenInvalid
}

func JWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := "", false
		if token, ok = c.GetPostForm("token"); !ok || token == "" {
			c.JSON(http.StatusOK, gin.H{
				"Msg": "No Token",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParserToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": errno.ErrTokenExpired.Code,
					"msg":  errno.ErrTokenExpired.Message,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"code": errno.ErrValidateFail.Code,
				"msg":  errno.ErrValidateFail.Message,
			})
			c.Abort()
			return
		}
		c.Set("id", claims.UserID) //把解析出来的userID放进头部  方便后续逻辑处理

	}
}
