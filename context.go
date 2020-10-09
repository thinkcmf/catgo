package catgo

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
}

type HandlerFunc func(*Context)

func (c *Context) Success(msg string, data interface{}) {
	c.JSON(200, gin.H{
		"msg":  msg,
		"data": data,
		"code": 1,
	})
}
