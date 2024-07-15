package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction = rune

const (
	Left  Direction = 'L'
	Up    Direction = 'U'
	Right Direction = 'R'
	Down  Direction = 'D'
)

func standardKeyMap() map[string]rune {
	keyMap := map[string]rune{
		"1L": '1', "1U": '1', "1R": '2', "1D": '4',
		"2L": '1', "2U": '2', "2R": '3', "2D": '5',
		"3L": '2', "3U": '3', "3R": '3', "3D": '6',
		"4L": '4', "4U": '1', "4R": '5', "4D": '7',
		"5L": '4', "5U": '2', "5R": '6', "5D": '8',
		"6L": '5', "6U": '3', "6R": '6', "6D": '9',
		"7L": '7', "7U": '4', "7R": '8', "7D": '7',
		"8L": '7', "8U": '5', "8R": '9', "8D": '8',
		"9L": '8', "9U": '6', "9R": '9', "9D": '9',
	}
	return keyMap
}

func diamondKeyMap() map[string]rune {
	keyMap := map[string]rune{
		"1L": '1', "1U": '1', "1R": '1', "1D": '3',
		"2L": '2', "2U": '2', "2R": '3', "2D": '6',
		"3L": '2', "3U": '1', "3R": '4', "3D": '7',
		"4L": '3', "4U": '4', "4R": '4', "4D": '8',
		"5L": '5', "5U": '5', "5R": '6', "5D": '5',
		"6L": '5', "6U": '2', "6R": '7', "6D": 'A',
		"7L": '6', "7U": '3', "7R": '8', "7D": 'B',
		"8L": '7', "8U": '4', "8R": '9', "8D": 'C',
		"9L": '8', "9U": '9', "9R": '9', "9D": '9',
		"AL": 'A', "AU": '6', "AR": 'B', "AD": 'A',
		"BL": 'A', "BU": '7', "BR": 'C', "BD": 'D',
		"CL": 'B', "CU": '8', "CR": 'C', "CD": 'C',
		"DL": 'D', "DU": 'B', "DR": 'D', "DD": 'D',
	}
	return keyMap
}

func getNextKey(currentKey rune, directions []Direction, keyMap map[string]rune) rune {
	if len(directions) == 0 {
		return currentKey
	}

	currentDirection := Direction(directions[0])
	nextKey := keyMap[fmt.Sprintf("%s%s", string(currentKey), string(currentDirection))]

	return getNextKey(nextKey, directions[1:], keyMap)
}

func getKeyCode(startKey rune, instructions []string, keyMap map[string]rune) string {
	keys := []rune{}
	currentKey := startKey

	for i := 0; i < len(instructions); i++ {
		currentLine := strings.TrimSpace(instructions[i])
		if currentLine == "" {
			continue
		}

		directions := []Direction(currentLine)
		currentKey = getNextKey(currentKey, directions, keyMap)
		keys = append(keys, currentKey)
	}

	return string(keys)
}

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")

	partOne := getKeyCode('5', lines, standardKeyMap())
	partTwo := getKeyCode('5', lines, diamondKeyMap())

	fmt.Printf("Part 1: %s, Part 2: %s\n", partOne, partTwo)
}
