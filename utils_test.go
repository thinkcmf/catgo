package catgo

import "testing"

func TestCheckMobile(t *testing.T) {
	mobile := "15100001234"
	result := CheckMobile(mobile)
	if !result {
		t.Error(mobile + " CheckMobile error")
	}

	mobile = "86-15100001234"
	result = CheckMobile(mobile)

	if !result {
		t.Error(mobile + " CheckMobile error")
	}

	mobile = "86-05100001234"
	result = CheckMobile(mobile)

	if result {
		t.Error(mobile + " CheckMobile error")
	}

}
