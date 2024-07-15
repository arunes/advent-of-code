package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction rune
type Turn rune
type Square struct {
	X int
	Y int
}

const (
	North Direction = 'N'
	East  Direction = 'E'
	South Direction = 'S'
	West  Direction = 'W'
	Left  Turn      = 'L'
	Right Turn      = 'R'
)

func visitedSquares(oldX, oldY, newX, newY int) []Square {
	response := []Square{}
	for x := min(oldX, newX); x <= max(oldX, newX); x++ {
		for y := min(oldY, newY); y <= max(oldY, newY); y++ {
			if oldX == x && oldY == y {
				continue
			}

			response = append(response, Square{X: x, Y: y})
		}
	}

	return response
}

func getDirection(current Direction, turn Turn) Direction {
	switch {
	case current == North && turn == Left:
		return West
	case current == North && turn == Right:
		return East
	case current == East && turn == Left:
		return North
	case current == East && turn == Right:
		return South
	case current == South && turn == Left:
		return East
	case current == South && turn == Right:
		return West
	case current == West && turn == Left:
		return South
	case current == West && turn == Right:
		return North
	default:
		return North
	}
}

func move(direction Direction, steps, currentX, currentY int) (int, int, []Square) {
	switch direction {
	case North:
		x, y := currentX, currentY-steps
		squares := visitedSquares(currentX, currentY, x, y)
		return x, y, squares

	case East:
		x, y := currentX+steps, currentY
		squares := visitedSquares(currentX, currentY, x, y)
		return x, y, squares

	case South:
		x, y := currentX, currentY+steps
		squares := visitedSquares(currentX, currentY, x, y)
		return x, y, squares

	case West:
		x, y := currentX-steps, currentY
		squares := visitedSquares(currentX, currentY, x, y)
		return x, y, squares

	default:
		return currentX, currentY, []Square{}

	}
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func main() {
	file, _ := os.ReadFile("input.txt")
	directions := strings.Split(string(file), ",")
	visitHistory := map[string]bool{"0x0": true}
	firstReVisitFound := false

	var x, y int
	facing := North
	for i := 0; i < len(directions); i++ {
		current := strings.TrimSpace(directions[i])

		if current == "" {
			continue
		}

		steps, _ := strconv.Atoi(current[1:])
		turn := Turn(rune(current[0]))
		facing = getDirection(facing, turn)

		var visited []Square
		x, y, visited = move(facing, steps, x, y)

		// Part 2
		if !firstReVisitFound {
			for i := 0; i < len(visited); i++ {
				key := fmt.Sprintf("%dx%d", visited[i].X, visited[i].Y)
				if visitHistory[key] {
					firstVisitedDistance := abs(visited[i].X) + abs(visited[i].Y)
					fmt.Printf("First re-visited square, x: %d, y: %d, distance: %d\n", visited[i].X, visited[i].Y, firstVisitedDistance)
					firstReVisitFound = true
				} else {
					visitHistory[key] = true
				}
			}
		}
	}

	distance := abs(x) + abs(y)
	fmt.Printf("Final position, x: %d, y: %d. Distance: %d\n", x, y, distance)
}
