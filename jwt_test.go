package catgo

import "testing"

func TestJwt(t *testing.T) {
	secret := "111111"
	token, err := JwtToken(map[string]interface{}{"data": "uid","test":"test2"}, secret)

	println("token:" + token)
	println(err)

	data, err := JwtParseToken(token, secret)

	Dump(data)
	println(err)

}
