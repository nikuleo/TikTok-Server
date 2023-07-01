package auth

import (
	"TikTokServer/pkg/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	ID   int64  `json:"user_id"`
	Name string `json:"user_name"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	ID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

var (
	secret     []byte
	expireHour int
	// refreshExpireHour int
)

func InitJWT() {
	cfg := config.GetConfig("jwtConfig")
	viper := cfg.Viper
	secret = []byte(viper.GetString("JWT.secret"))
	expireHour = viper.GetInt("JWT.expireHour")
	// refreshExpireHour = viper.GetInt("JWT.refreshExpireHour")
}

func CreateToken(userID int64, userName string) (token string, err error) {
	//FIXME: middleware err: token is expired by 12.031758694s
	expireTime := time.Now().Add(time.Hour * time.Duration(expireHour))
	claims := &JwtClaims{
		ID:   userID,
		Name: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(requestToken string) (bool, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil // HS256 对称加密（公钥与私钥相同），若非对称加密则这里填写公钥(私钥签署， 公钥验证)
	})
	if err != nil {
		return false, err
	}
	if _, ok := token.Claims.(*JwtClaims); !ok && !token.Valid {
		return false, fmt.Errorf("token is invalid")
	}
	return true, nil
}

func GetUserIDByToken(requestToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(requestToken, &JwtClaims{}, func(token *jwt.Token) (any, error) {
		return secret, nil
	})

	if err != nil {
		return -1, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims.ID, nil
	}

	return -1, err
}
