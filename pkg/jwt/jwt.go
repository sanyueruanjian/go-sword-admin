package jwt

import (
	"errors"
	"time"

	"project/utils/config"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int, username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.JwtConfig.Timeout) * time.Second).Unix(), // 过期时间
			Issuer:    "my-project",                                                                 // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(config.JwtConfig.Secret))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(config.JwtConfig.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
