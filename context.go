package catgo

import (
	"github.com/gin-gonic/gin"
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

	c.JSON(200, gin.H{
		"msg":  msg,
		"data": data,
		"code": code,
	})
}
