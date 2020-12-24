package catgo

import (
	"github.com/icatgo/php"
	"strings"
)

//
func Password(password string, authCode ...string) string {
	mAuthCode := ""
	if len(authCode) > 0 {
		mAuthCode = authCode[0]
	} else {
		if value, ok := DBConfig["authcode"]; ok {
			mAuthCode = value
		}
	}

	return "###" + php.Md5(php.Md5(mAuthCode+password))
}

func ComparePassword(password, passwordInDb string) bool {
	result := false
	if strings.Index(passwordInDb, "###") >= 0 {
		result = Password(password) == passwordInDb
	}

	return result
}
