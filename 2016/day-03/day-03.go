package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSides(line string) []int {
	splitFn := func(c rune) bool {
		return c == ' '
	}

	sides := strings.FieldsFunc(line, splitFn)
	side1, _ := strconv.Atoi(sides[0])
	side2, _ := strconv.Atoi(sides[1])
	side3, _ := strconv.Atoi(sides[2])

	return []int{
		side1, side2, side3,
	}
}

func isTriangle(side1, side2, side3 int) bool {
	return side1+side2 > side3 && side1+side3 > side2 && side2+side3 > side1
}

func countHorizontally(lines []string) int {
	triangles := 0
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		sides := getSides(line)

		if isTriangle(sides[0], sides[1], sides[2]) {
			triangles++
		}
	}
	return triangles
}

func countVertically(lines []string) int {
	triangles := 0
	for i := 0; i < len(lines)-2; i += 3 {
		line1 := strings.TrimSpace(lines[i])
		line2 := strings.TrimSpace(lines[i+1])
		line3 := strings.TrimSpace(lines[i+2])

		if line1 == "" || line2 == "" || line3 == "" {
			continue
		}

		sides1 := getSides(line1)
		sides2 := getSides(line2)
		sides3 := getSides(line3)

		if isTriangle(sides1[0], sides2[0], sides3[0]) {
			triangles++
		}

		if isTriangle(sides1[1], sides2[1], sides3[1]) {
			triangles++
		}

		if isTriangle(sides1[2], sides2[2], sides3[2]) {
			triangles++
		}

	}
	return triangles

}

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")

	// Part 1
	trianglesSameRow := countHorizontally(lines)

	// Part 2
	trianglesSameColumn := countVertically(lines)

	fmt.Printf("Possible triangles, part 1: %d, part 2: %d\n", trianglesSameRow, trianglesSameColumn)
}
