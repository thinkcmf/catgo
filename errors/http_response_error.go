package errors

type HttpResponseError struct {
	HTTPCode int // 200,404
	Data     map[string]interface{}
	Message  string
	Type     string // JSON,JSONP
}

func (c *HttpResponseError) Error() string {
	return c.Message
}
