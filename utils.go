package catgo

import (
	"github.com/asaskevich/govalidator"
	"github.com/rocket049/gostructcopy"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func GetCurrentPath() (string, error) {
	dir, _ := os.Executable()
	path := filepath.Dir(dir)
	//log.Println(exPath)
	//
	//file, err := exec.LookPath(os.Args[0])
	//if err != nil {
	//	return "", err
	//}
	//path, err := filepath.Abs(file)
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println("path111:", path)
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	//fmt.Println("path222:", path)
	//i := strings.LastIndex(path, "/")
	//if i < 0 {
	//	return "", errors.New(`Can't find "/" or "\".`)
	//}
	//fmt.Println("path333:", path)
	return path, nil
}

func CheckMobile(mobile string) bool {

	result, _ := regexp.MatchString(`(^(13\d|14\d|15\d|16\d|17\d|18\d|19\d)\d{8})$`, mobile)
	if !result {
		result, _ = regexp.MatchString(`^\d{1,4}-[1-9]\d{4,10}$`, mobile)
	}

	return result

}

//StructCopy copy the exported value of a struct to a likely struct , with reflect.
func StructCopy(src, des interface{}) error {
	return gostructcopy.StructCopy(src, des)
}

var (
	IsEmail         = govalidator.IsEmail
	IsExistingEmail = govalidator.IsExistingEmail
	IsAlpha         = govalidator.IsAlpha
	IsAlphanumeric  = govalidator.IsAlphanumeric
	IsBase64        = govalidator.IsBase64
	IsFloat         = govalidator.IsFloat
	IsIP            = govalidator.IsIP
)
