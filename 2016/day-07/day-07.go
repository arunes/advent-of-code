package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func isABBA(chars []rune) bool {

	openBrackets := 0
	response := false
	for i := 0; i < len(chars)-3; i++ {
		if chars[i] == '[' {
			openBrackets++
			continue
		} else if chars[i] == ']' {
			openBrackets--
			continue
		}

		current := chars[i : i+2]
		next := chars[i+2 : i+4]

		// skip check if parts has brackets
		if strings.ContainsAny(string(current), "[]") || strings.ContainsAny(string(next), "[]") {
			continue
		}

		isPair := current[0] != current[1] && current[0] == next[1] && current[1] == next[0]
		if isPair && openBrackets > 0 {
			return false
		} else if isPair {
			response = isPair
		}
	}

	return response
}

func isABA(chars []rune) bool {
	openBrackets := 0
	for i := 0; i < len(chars)-2; i++ {
		if chars[i] == '[' {
			openBrackets++
			continue
		} else if chars[i] == ']' {
			openBrackets--
			continue
		}

		first := chars[i]
		middle := chars[i+1]
		last := chars[i+2]

		// skip check, if has brackets and it not a pattern
		if openBrackets > 0 || first != last || strings.ContainsAny(string(chars[i:i+3]), "[]") {
			continue
		}

		pattern := regexp.MustCompile(fmt.Sprintf("\\[[^\\]]*%c%c%c[^\\[]*\\]", middle, first, middle))
		found := pattern.MatchString(string(chars))

		if found {
			return true
		}
	}

	return false
}

func getTotalPairs(lines []string, pairFn func([]rune) bool) int {
	totalPairs := 0
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		if pairFn([]rune(line)) {
			totalPairs++
		}
	}

	return totalPairs
}

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")

	totalABBA := getTotalPairs(lines, isABBA)
	totalABA := getTotalPairs(lines, isABA)
	fmt.Printf("Part 1: %d, part 2: %d\n", totalABBA, totalABA)

}
