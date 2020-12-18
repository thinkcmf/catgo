package catgo

import (
	"testing"
)

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

func TestUtils(t *testing.T) {
	IsEmail("xx")
	IsExistingEmail("dd")
}

func TestStructCopy(t *testing.T) {
	println("--------------TestStructCopy----------------")
	src := &struct {
		Name  string
		Name2 string
		Name3 string
	}{
		Name:  "dddd",
		Name2: "name2",
		Name3: "name3",
	}

	des := &struct {
		Name  string
		Name3 string
	}{}

	StructCopy(src, des)

	Dump(des)
	println("--------------TestStructCopy----------------")

}
