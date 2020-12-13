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
		mAuthCode = DBConfig["authcode"]
	}

	return "###" + php.Md5(php.Md5(mAuthCode+password))
}

func ComparePassword(password, passwordInDb string) bool {
	result := false
	if strings.Index(password, "###") >= 0 {
		result = Password(password) == passwordInDb
	}

	return result
}
