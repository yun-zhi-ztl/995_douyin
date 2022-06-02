package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("Key")

type JwtClaims1 struct {
	*jwt.StandardClaims
	//用户编号
	UserId uint
}

func CreateJwtToken(id uint) (string, error) {
	// 定义过期时间,7天后过期
	expireToken := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := JwtClaims1{
		&jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // token信息生效时间
			ExpiresAt: int64(expireToken),              // 过期时间
			Issuer:    "douyin",                        // 发布者
		},
		id,
	}
	//SigningMethodHS256,HS256对称加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//通过自定义令牌加密
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

// ParseToken 解析JWT
func ParserToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims1{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return 0, fmt.Errorf("token不可用")
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return 0, fmt.Errorf("token过期")
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return 0, fmt.Errorf("无效的token")
			} else {
				return 0, fmt.Errorf("token不可用")
			}
		}
	}
	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	if claims, ok := token.Claims.(*JwtClaims1); ok && token.Valid {
		return int(claims.UserId), nil
	}
	return 0, fmt.Errorf("token无效")
}
