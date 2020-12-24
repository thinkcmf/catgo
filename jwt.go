package catgo

import (
	"github.com/dgrijalva/jwt-go"
)

func JwtToken(data map[string]interface{}, secret ...string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))

	secretStr := ""
	if len(secret) > 0 {
		secretStr = secret[0]
	} else {
		if value, ok := DBConfig["jwt_secret"]; ok {
			secretStr = value
		}
	}

	token, err := at.SignedString([]byte(secretStr))
	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtParseToken(token string, secret ...string) (map[string]interface{}, error) {
	secretStr := ""
	if len(secret) > 0 {
		secretStr = secret[0]
	} else {
		if value, ok := DBConfig["jwt_secret"]; ok {
			secretStr = value
		}
	}

	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretStr), nil
	})
	if err != nil {
		return nil, err
	}

	return claim.Claims.(jwt.MapClaims), nil
}
