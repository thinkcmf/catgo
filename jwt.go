package catgo

import (
	"github.com/dgrijalva/jwt-go"
)

func JwtToken(data map[string]interface{}, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))

	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtParseToken(token string, secret string) (map[string]interface{}, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return claim.Claims.(jwt.MapClaims), nil
}
