package main

import (
	"fmt"
	"os"
	"strings"
)

// taken from: https://github.com/alexchao26/advent-of-code-go/blob/main/2016/day09/main.go
func decompressLength(in string, part int) int {
	var decompressedLen int
	for i := 0; i < len(in); {
		switch in[i] {
		case '(':
			// find index of closing paren, then find total length of substring
			relativeCloseIndex := strings.Index(in[i:], ")")
			closeIndex := relativeCloseIndex + i

			var copyLen, repeat int
			fmt.Sscanf(in[i:closeIndex+1], "(%dx%d)", &copyLen, &repeat)

			substring := in[closeIndex+1 : closeIndex+1+copyLen]
			patternLength := len(substring)
			if part == 2 {
				patternLength = decompressLength(substring, 2)
			}
			decompressedLen += patternLength * repeat
			// jump the closed paren (+1) the length of the substring from THIS
			// function call
			i = closeIndex + 1 + len(substring)
		default:
			decompressedLen++
			i++
		}
	}
	return decompressedLen
}

func main() {
	input, _ := os.ReadFile("input.txt")
	part1 := decompressLength(string(input), 1)
	part2 := decompressLength(string(input), 2)

	fmt.Printf("Part 1: %d, part 2: %d\n", part1, part2)
}
