// Package middleware
// @author ufec https://github.com/ufec
// @date 2022/5/11
package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuth
//  @Description: 类似于JWT中间件
//  @param where string 由于请求token位置不固定，通过指定位置来获取token
//  @return gin.HandlerFunc
func JWTAuth(where string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		// where可选值: query/form/header/body json
		switch where {
		case "query":
			token = c.Query("token")
		case "form":
			token = c.PostForm("token")
		default:
			token = c.Query("token")
		}
		// 不存在该用户token则直接抛出用户不存在错误信息
		// if _, exists := controller.UsersLoginInfo[token]; !exists {
		// 	c.JSON(http.StatusOK, controller.Response{StatusCode: 1, StatusMsg: "token鉴权失败, 非法操作"})
		// 	c.Abort()
		// 	return
		// }
		// 统一获取user的位置 后续流程直接从上下文取 user 即可
		// 此方法不可取 gin 不支持设置指定数据类型的数据，需要通过json序列化 反序列化来完成类型转换 损失性能
		//c.Set("user", controller.UsersLoginInfo[token])

		// 设置Token也是不必要的 从整体来看 我们仅仅只需要用户ID便能唯一确定用户, gin 也支持获取基础数据类型 恰好符合要求
		c.Set("token", token)
		//c.Set("userID", controller.UsersLoginInfo[token].ID)
		c.Next()
	}
}

type JwtClaims struct {
	*jwt.StandardClaims
	//用户编号
	Username     string
	Userpassword string
}

var mySigningKey = []byte("Key")

//创建token
// CreateJwtToken 生成一个jwttoken
func CreateJwtToken(username, password string) (string, error) {
	// 定义过期时间,7天后过期
	expireToken := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // token信息生效时间
			ExpiresAt: int64(expireToken),              // 过期时间
			Issuer:    "douyin",                        // 发布者
		},
		username,
		password,
	}
	//SigningMethodHS256,HS256对称加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//通过自定义令牌加密
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

type JwtClaims1 struct {
	*jwt.StandardClaims
	//用户编号
	UserId int
}

func CreateJwtToken1(id int) (string, error) {
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
func ParserToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token不可用")
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token过期")
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("无效的token")
			} else {
				return nil, fmt.Errorf("token不可用")
			}
		}
	}
	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token无效")
}
