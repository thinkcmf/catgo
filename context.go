package catgo

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkcmf/catgo/errors"
	"strconv"
	"strings"
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

// set user id
func (c *Context) SetUserId(userId string) {
	c.Set("CATGO_USER_ID", userId)
}

// get user id
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

// param int64
func (c *Context) ParamInt64(key string) (result int64) {

	result, _ = strconv.ParseInt(c.Param(key), 10, 64)

	return
}

// param uint64
func (c *Context) ParamUint64(key string) (result uint64) {

	result, _ = strconv.ParseUint(c.Param(key), 10, 64)

	return
}

// param int
func (c *Context) ParamInt(key string) (result int) {

	result, _ = strconv.Atoi(c.Param(key))

	return
}

// param uint
func (c *Context) ParamUint(key string) (result uint) {

	r, err := strconv.ParseUint(c.Param(key), 10, 32)

	if err != nil {
		result = 0
	} else {
		result = uint(r)
	}
	return
}

// param int64 with default value defaultResult
func (c *Context) DefaultParamInt64(key string, defaultResult int64) (result int64) {

	result, _ = strconv.ParseInt(c.Param(key), 10, 64)

	if result == 0 {
		result = defaultResult
	}

	return
}

// param uint64 with default value defaultResult
func (c *Context) DefaultParamUint64(key string, defaultResult uint64) (result uint64) {

	result, _ = strconv.ParseUint(c.Param(key), 10, 64)

	if result == 0 {
		result = defaultResult
	}

	return
}

// param int with default value defaultResult
func (c *Context) DefaultParamInt(key string, defaultResult int) (result int) {

	result, _ = strconv.Atoi(c.Param(key))

	if result == 0 {
		result = defaultResult
	}
	return
}

// param uint with default value defaultResult
func (c *Context) DefaultParamUint(key string, defaultResult uint) (result uint) {

	r, err := strconv.ParseUint(c.Param(key), 10, 32)

	if err != nil {
		result = 0
	} else {
		result = uint(r)
	}

	if result == 0 {
		result = defaultResult
	}

	return
}

// query int64
func (c *Context) QueryInt64(key string) (result int64) {

	result, _ = strconv.ParseInt(c.Query(key), 10, 64)

	return
}

// query uint64
func (c *Context) QueryUint64(key string) (result uint64) {

	result, _ = strconv.ParseUint(c.Query(key), 10, 64)

	return
}

// query int
func (c *Context) QueryInt(key string) (result int) {

	result, _ = strconv.Atoi(c.Query(key))

	return
}

// query int
func (c *Context) QueryUint(key string) (result uint) {

	r, err := strconv.ParseUint(c.Query(key), 10, 32)

	if err != nil {
		result = 0
	} else {
		result = uint(r)
	}

	return
}

// query int64 with default value defaultResult
func (c *Context) DefaultQueryInt64(key string, defaultResult int64) (result int64) {

	result, _ = strconv.ParseInt(c.Query(key), 10, 64)

	if result == 0 {
		result = defaultResult
	}

	return
}

// query uint64 with default value defaultResult
func (c *Context) DefaultQueryUint64(key string, defaultResult uint64) (result uint64) {

	result, _ = strconv.ParseUint(c.Query(key), 10, 64)

	if result == 0 {
		result = defaultResult
	}

	return
}

// query int with default value defaultResult
func (c *Context) DefaultQueryInt(key string, defaultResult int) (result int) {

	result, _ = strconv.Atoi(c.Query(key))

	if result == 0 {
		result = defaultResult
	}

	return
}

// query uint with default value defaultResult
func (c *Context) DefaultQueryUint(key string, defaultResult uint) (result uint) {

	r, err := strconv.ParseUint(c.Query(key), 10, 32)

	if err != nil {
		result = 0
	} else {
		result = uint(r)
	}

	if result == 0 {
		result = defaultResult
	}

	return
}

// parse query page
// defaultPageSize 10
func (c *Context) QueryPage(defaultPageSize ...int) (page, pageSize int) {

	pageStr := c.Query("page")

	if pageStr != "" {
		pageArr := strings.Split(pageStr, ",")
		pageArrLength := len(pageArr)
		if len(pageArr) == 1 {
			page, _ = strconv.Atoi(pageArr[0])
		} else if pageArrLength == 2 {
			page, _ = strconv.Atoi(pageArr[0])
			pageSize, _ = strconv.Atoi(pageArr[1])
		}
	}

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 1
	}

	if pageSize > 500 {
		pageSize = 500
	}

	return
}
