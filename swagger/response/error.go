package response

// A Error is an error that is used when error.
// swagger:response Error
type Error struct {
	// The error message
	// in: body
	Body struct {
		// status code 0 - ERR
		//
		// Required: true
		Code int `json:"code"`

		// The  message
		//
		// Required: true
		// Example: error
		Msg string `json:"msg"`

		// data mixed

		// Required: true
		Data interface{} `json:"data"`
	}
}
