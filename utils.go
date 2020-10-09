package catgo

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func init() {
	path, _ := GetCurrentPath()

	log.Println(path)
}

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
