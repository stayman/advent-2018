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
	var res uint64
	for _, val := range vals {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("Bad Input: %v", err))
		}
		res += intVal
	}

	fmt.Printf("The result of the sequence is: %d\n", res)

	values := make(map[int]int)
	var res2 int
	var repeat int
	var done bool

	for {
		for _, val := range vals {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				panic(fmt.Sprintf("Bad Input: %v", err))
			}
			res2 += intVal
			values[res2] = values[res2] + 1

			if values[res2] > 1 {
				repeat = res2
				done = true
				break
			}
		}

		if done {
			break
		}
	}

	fmt.Printf("The first duplicated frequency is: %d\n", repeat)
}
