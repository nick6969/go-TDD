package jwt

import (
	"errors"
	"time"

	"github.com/pascaldekloe/jwt"
)

type JWT struct {
	Secret string
}

func (j JWT) GenerateUserToken(id uint) (token string, err error) {
	var claims jwt.Claims
	now := time.Now()
	claims.Issued = jwt.NewNumericTime(now.Round(time.Second))
	claims.Expires = jwt.NewNumericTime(now.AddDate(0, 0, 1).Round(time.Second))
	claims.Set = map[string]interface{}{
		"id": id,
	}
	t, err := claims.HMACSign("HS512", []byte(j.Secret))

	if err != nil {
		return
	}
	token = string(t)
	return
}

func (j JWT) VerifyUserToken(token string) (id uint, err error) {
	claim, err := jwt.HMACCheck([]byte(token), []byte(j.Secret))

	if err != nil {
		return
	}

	tokenNotCorrect := "Token is not Correct."

	if claim.Expires.Time().Before(time.Now()) {
		err = errors.New(tokenNotCorrect)
		return
	}

	userID, ok := claim.Set["id"].(float64)
	if !ok {
		err = errors.New(tokenNotCorrect)
		return
	}
	id = uint(userID)

	return
}
