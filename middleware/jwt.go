package middleware

import (
	"douyin/config"
	"douyin/errno"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	JwtKey []byte
}

// NewJWT 创建JWT 实例
func NewJWT() *JWT {
	return &JWT{JwtKey: []byte(config.JwtKey)}
}

// MyClaims 自定义Claim
type MyClaims struct {
	UserID int64 //用户ID
	jwt.StandardClaims
}

var (
	TokenExpired     error = errors.New("token 已过期，请重新登录")
	TokenNotValidYet error = errors.New("token 无效，请重新登录")
	TokenMalFormed   error = errors.New("token 不正确，请重新登录")
	TokenInvalid     error = errors.New("这不是一个Token，请重新登录")
)

// SetUpToken 设置 claims，为生成 token 制作准备 "claim"
func SetUpToken(userID int64) (string, error) {
	j := NewJWT()
	claims := MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 240,
			ExpiresAt: time.Now().Unix() + config.TokenExpiredTime,
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return "", errors.New("token 生成失败")
	}
	return token, nil
}

// CreateToken 通过加密和claim创建token
func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token，返回定义的 Claims
// 如果出现错误，则返回对应的错误信息
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

// JWTToken 解析、验证token，并把解析出来的user_id 通过ctx.Set() 方法增加到 gin.Context 头部中
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
