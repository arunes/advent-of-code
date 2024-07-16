package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func getErrorCorrectedMessage(lines *[]string, mostCommon bool) string {
	size := len((*lines)[0])
	result := make([]rune, size)
	columns := map[int]string{}

	for i := 0; i < len(*lines); i++ {
		line := []rune((*lines)[i])
		if len(line) == 0 {
			continue
		}

		for j := 0; j < size; j++ {
			columns[j] = fmt.Sprintf("%s%s", columns[j], string(line[j]))
		}
	}

	for i := 0; i < size; i++ {
		current := []rune(columns[i])

		sort.Slice(current, func(i, j int) bool {
			left := strings.Count(string(current), string(current[i]))
			right := strings.Count(string(current), string(current[j]))

			if mostCommon {
				return left > right
			} else {
				return left < right
			}
		})

		result[i] = current[0]
	}

	return string(result)
}

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")

	part1Message := getErrorCorrectedMessage(&lines, true)
	part2Message := getErrorCorrectedMessage(&lines, false)

	fmt.Printf("Part 1: %s, part 2: %s\n", part1Message, part2Message)
}
