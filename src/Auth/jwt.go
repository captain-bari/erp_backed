package auth

import (
	"fmt"
	"os"
	"time"

	log "erp/log"

	"github.com/golang-jwt/jwt"
)

var key []byte

func GenerateJWT(email, role string) (string, error) {

	if len(key) == 0 {
		b, err := os.ReadFile("./key_open.pem")
		if err != nil {
			log.Errorln(err)
		}
		key = b
	}

	claims := &jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Second * 5).Unix(),
		"role":       role,
		"email":      role,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func CheckJwt(tokenString string) (valid bool, role string, err error) {
	tokenresp, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	claimsResp := tokenresp.Claims.(jwt.MapClaims)
	role = fmt.Sprintln(claimsResp["role"])
	valid = tokenresp.Valid
	return
}
