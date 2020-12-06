package response

// A Success is an Success that is used when Success.
// swagger:response Success
type SuccessResponse struct {
	// The success message
	// in: body
	Body SuccessBody
}

type SuccessBody struct {
	// status code 1 - OK
	//
	// Required: true
	// Example: 1
	Code int `json:"code"`

	// The  message
	//
	// Required: true
	// Example: success
	Msg string `json:"msg"`

	// data mixed
	//
	// Required: true
	Data interface{} `json:"data"`
}
