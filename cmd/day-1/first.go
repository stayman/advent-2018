package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/stayman/advent/pkg/input"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	vals := input.GetInput(filepath.Join(wd, "cmd", "day-1", "input.txt"))
	var res int
	for _, val := range vals {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("Bad Input: %v", err))
		}
		res += intVal
	}

	fmt.Println(res)

}
