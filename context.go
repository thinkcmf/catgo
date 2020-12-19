package catgo

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkcmf/catgo/errors"
)

type Context struct {
	*gin.Context
}

type HandlerFunc func(*Context)

// msg
// params[0] data
// params[1] map[string]string header
func (c *Context) Success(msg string, params ...interface{}) {
	c.Result(msg, 1, params...)
}

// msg
// params[0] data
// params[1] map[string]string header
func (c *Context) Error(msg string, params ...interface{}) {
	c.Result(msg, 0, params...)
}

// msg
// code 0|1
// params[0] data
// params[1] map[string]string header
func (c *Context) Result(msg string, code int, params ...interface{}) {
	var data interface{}
	length := len(params)

	switch length {
	case 1:
		data = params[0]
	case 2:
		data = params[0]
		if headers, ok := params[1].(map[string]string); ok {
			for key, value := range headers {
				c.Header(key, value)
			}
		}
	}

	responseError := errors.HttpResponseError{
		HTTPCode: 200,
		Data: gin.H{
			"msg":  msg,
			"data": data,
			"code": code,
		},
		Message: msg,
		Type:    "JSON",
	}
	panic(responseError)
}

func (c *Context) SetUserId(userId string) {
	c.Set("CATGO_USER_ID", userId)
}

func (c *Context) UserId() string {
	value, exists := c.Get("CATGO_USER_ID")

	if exists {
		return value.(string)
	}

	return ""
}

// Validate data then fail error
func (c *Context) ValidateFailError(data interface{}, scene ...string) {

	if result, msg := Validate(data, scene...); !result {
		c.Error(msg)
	}
}

// ShouldBind & Validate data then fail error
func (c *Context) ShouldBindValidateFailError(obj interface{}, msg1scene2 ...string) {
	msg1scene2Length := len(msg1scene2)
	err := c.ShouldBind(obj)
	if err != nil {
		if msg1scene2Length > 0 {
			c.Error(msg1scene2[0])
		} else {
			c.Error(err.Error())
		}
	}

	scene := make([]string, 0)

	if msg1scene2Length > 1 {
		scene = msg1scene2[1:]
	}

	if result, msg := Validate(obj, scene...); !result {
		c.Error(msg)
	}

}

// ShouldBind  data then fail error
func (c *Context) ShouldBindFailError(obj interface{}, msg ...string) {

	err := c.ShouldBind(obj)
	if err != nil {
		if len(msg) > 0 {
			c.Error(msg[0])
		} else {
			c.Error(err.Error())
		}
	}
}
