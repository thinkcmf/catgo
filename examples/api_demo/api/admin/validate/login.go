package validate

import "fmt"

type LoginForm struct {
	Username string `form:"username" validate:"required|CustomValidator"`
	Password string `form:"password" validate:"required"`
}

// Messages you can custom validator error messages.
func (f LoginForm) Messages() map[string]string {
	return map[string]string{
		"required":      "oh! the {field} is required",
		"Username.required": "{field} is required",
		"Username.CustomValidator": "the {field} CustomValidator is not pass!",
	}
}

// Translates you can custom field translates.
func (f LoginForm) Translates() map[string]string {
	return map[string]string{
		"Username": "User Name",
		"Email":    "User Email",
	}
}

// CustomValidator custom validator in the source struct.
func (f LoginForm) CustomValidator(val string) bool {
	fmt.Println(f)
	return len(val) == 4
}
