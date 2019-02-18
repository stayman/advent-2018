package input

import (
	"io/ioutil"
	"strings"
)

func GetInput(path string) []string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(content), "\n")
}
