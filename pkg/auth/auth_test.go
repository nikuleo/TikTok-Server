package auth

import (
	"TikTokServer/pkg/tlog"
	"fmt"
	"testing"
)

func init() {
	tlog.InitLog()
	InitJWT()
}

func TestInitJWT(t *testing.T) {
	fmt.Println(secret, expireHour)
}

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(2233, "niku")
	fmt.Println(token, err)
}

func TestValidateToken(t *testing.T) {
	token, err := CreateToken(2233, "niku")
	fmt.Println(token, err)
	isValidated, err := ValidateToken(token)
	fmt.Println(isValidated, err)
}

func TestGetUserIDByToken(t *testing.T) {
	token, err := CreateToken(1, "niku")
	fmt.Println("token:", token, err)
	userID, err := GetUserIDByToken(token)
	fmt.Println("userID:", userID, err)
}
