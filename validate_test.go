package catgo

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// UserForm struct
type UserForm struct {
	Name     string    `validate:"required|minLen:7"`
	Email    string    `validate:"email" message:"email is invalid"`
	Age      int       `validate:"required|int|min:1|max:99" message:"int:age must int| min: age min value is 1"`
	CreateAt int       `validate:"min:1"`
	Safe     int       `validate:"-"`
	UpdateAt time.Time `validate:"required"`
	Code     string    `validate:"customValidator"`
}

// CustomValidator custom validator in the source struct.
func (f UserForm) CustomValidator(val string) bool {
	return len(val) == 4
}

// Messages you can custom validator error messages.
func (f UserForm) Messages() map[string]string {
	return map[string]string{
		"required":      "oh! the {field} is required",
		"Name.required": "message for special field",
	}
}

// Translates you can custom field translates.
func (f UserForm) Translates() map[string]string {
	return map[string]string{
		"Name":  "User Name",
		"Email": "User Email",
	}
}

func TestValidate(t *testing.T) {

	u := &UserForm{
		Name: "inhere",
	}

	if result, msg := Validate(u); !result {
		println(msg)
	}

}

func TestValidateBatch(t *testing.T) {

	u := &UserForm{
		Name: "inhere",
	}

	if result, messages := ValidateBatch(u); !result {
		fmt.Println(messages)

		messagesJson, _ := json.MarshalIndent(messages,"","    ")
		fmt.Println(string(messagesJson))

	}

}
