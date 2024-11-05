package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/wonderivan/logger"
	"operations-platform/config"
)

var JWTToken jwtToken

type jwtToken struct {
}

// token解析后对应的结构体, 包含自定义信息和jwt签名信息
// 推荐使用userId代替username
type CustomClaims struct {
	Username string `json: "username"`
	Password string `json: "password"`
	jwt.StandardClaims
}

//验证token是否有效

func (*jwtToken) ParseToken(tokenStr string) (claims *CustomClaims, err error) {
	//第一个参数 token的string类型
	//第二个是 想解析成什么类型
	//第三个是个函数，主要是为了把加密因子传进去
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		logger.Error("parse token failed:", err)
		//处理token解析报错的各种情况
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("TokenMalformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("TokenExpired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("TokenNotValidyet")
			} else {
				return nil, errors.New("TokenInvalid")
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("parse Token failed")
}
