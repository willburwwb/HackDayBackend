package verify

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 加密密钥 随机
var jwtKey = []byte("secretKey")

type Claims struct {
	// user id
	Id uint
	jwt.RegisteredClaims
}

func ReleaseToken(id uint, admin bool) (string, error) {
	// last for 15 days
	expireTime := time.Now().Add(15 * 24 * time.Hour)

	claims := &Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			Issuer:    "backend",
			Subject:   "token",
		},
	}

	//get token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey) //利用密钥生成token字符串

	if err != nil {
		return "获取失败", err
	}

	return tokenString, nil
}

// parse token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
