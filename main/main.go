package main

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"log"
)

func main() {
	str2 := `
name: test
test:
  - test: test
  - test2: test
`
	data := []byte(str2)
	yamlDecoder:= yaml.NewDecoder(bytes.NewBuffer(data))


	var out map[interface{}]interface{}
	yamlDecoder.Decode(&out)

	log.Println(out)
}
