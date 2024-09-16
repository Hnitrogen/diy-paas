package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	// 设置密钥 TODO viper.get()
	var mySigningKey = []byte("xakjwaopdmsllldajiwjinik")

	// 创建声明
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 过期时间
			Issuer:    "base-app",
		},
	}

	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并返回令牌字符串
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}
