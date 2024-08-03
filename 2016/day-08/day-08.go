package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getBoard(width, height int) [][]bool {
	array := make([][]bool, height)

	for i := range array {
		array[i] = make([]bool, width)
	}

	return array
}

func rect(board *[][]bool, x, y int) {
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			(*board)[row][col] = true
		}
	}
}

func rotateColumn(board *[][]bool, column, by int) {
	isOn := (*board)[len(*board)-1][column]
	for row := 0; row < len(*board); row++ {
		current := (*board)[row][column]
		(*board)[row][column] = isOn
		isOn = current
	}

	if by > 1 {
		rotateColumn(board, column, by-1)
	}
}

func rotateRow(board *[][]bool, row, by int) {
	isOn := (*board)[row][len((*board)[row])-1]
	for col := 0; col < len((*board)[row]); col++ {
		current := (*board)[row][col]
		(*board)[row][col] = isOn
		isOn = current
	}

	if by > 1 {
		rotateRow(board, row, by-1)
	}
}

func getCount(board *[][]bool, isLit bool) int {
	count := 0

	for row := 0; row < len(*board); row++ {
		for col := 0; col < len((*board)[row]); col++ {
			if (*board)[row][col] == isLit {
				count++
			}
		}
	}

	return count
}

func display(board *[][]bool) {
	for row := 0; row < len(*board); row++ {
		for col := 0; col < len((*board)[row]); col++ {
			if (*board)[row][col] {
				print("#")
			} else {
				print("·")
			}
		}
		print("\n")
	}

	print("\n")
}

var pattern = regexp.MustCompile(`(rect|rotate row|rotate column).+?(\d+).+?(\d+)`)

func parseInstructions(line string) (action string, number1 int, number2 int) {
	match := pattern.FindStringSubmatch(line)
	action = match[1]
	number1, _ = strconv.Atoi(match[2])
	number2, _ = strconv.Atoi(match[3])
	return
}

func followInstructions(board *[][]bool, line string) {
	action, n1, n2 := parseInstructions(line)

	switch action {
	case "rect":
		rect(board, n1, n2)
		break

	case "rotate row":
		rotateRow(board, n1, n2)
		break

	case "rotate column":
		rotateColumn(board, n1, n2)
		break
	}
}

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")
	board := getBoard(50, 6)

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		followInstructions(&board, lines[i])
	}

	display(&board)

	lit := getCount(&board, true)
	fmt.Printf("Total lit: %d\n", lit)
}
