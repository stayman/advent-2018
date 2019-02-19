package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/stayman/advent/pkg/input"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	vals := input.GetInput(filepath.Join(wd, "cmd", "day-2", "input.txt"))
	var threes int
	var twos int
	for _, val := range vals {
		runeCount := make(map[rune]int)
		for _, char := range val {
			runeCount[char] = runeCount[char] + 1
		}

		if contains(runeCount, 2) {
			twos = twos + 1
		}

		if contains(runeCount, 3) {
			threes = threes + 1
		}

		// Interesting conjecture... Which of these options is actually more performant?
		// Above or below?

		// var threesPassed bool
		// var twosPassed bool
		// for _, count := range runeCount {
		// 	if count == 3 && !threesPassed {
		// 		threes = threes + 1
		// 		threesPassed = true
		// 	}

		// 	if count == 2 && !twosPassed {
		// 		twos = twos + 1
		// 		twosPassed = true
		// 	}
		// }

	}

	fmt.Printf("The first half result (checksum) is %d\n", threes*twos)
}

func contains(dict map[rune]int, expectedVal int) bool {
	for _, count := range dict {
		if count == expectedVal {
			return true
		}
	}

	return false
}
