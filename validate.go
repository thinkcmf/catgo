package catgo

import (
	"github.com/gookit/validate"
)

func Validate(data interface{}, scene ...string) (bool, string) {
	v := validate.New(data)
	v.StopOnError = true

	result := true
	msg := ""

	if !v.Validate(scene...) {
		result = false
		msg = v.Errors.One()
	}
	return result, msg
}

func ValidateBatch(data interface{}, scene ...string) (bool, map[string]map[string]string) {
	v := validate.New(data)
	v.StopOnError = false

	result := true
	msg := map[string]map[string]string{}

	if !v.Validate(scene...) {
		result = false
		for key, messages := range v.Errors {
			println("key:" + key)
			msg[key] = map[string]string{}
			for key2, value := range messages {
				println("key2:" + key2)
				println("value:" + value)
				msg[key][key2] = value
			}
		}
	}
	return result, msg
}
