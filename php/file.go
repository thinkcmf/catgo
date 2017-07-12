package php

import (
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

func FileGetContents(path string) (string, error) {

	content, err := FileGetBytes(path)

	return string(content), err
}

func FileGetBytes(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	return fd, err
}

func FileMd2Html(path string) string {

	contentBytes, _ := FileGetBytes(path)
	contentBytes = blackfriday.MarkdownCommon(contentBytes)

	return string(contentBytes)
}
