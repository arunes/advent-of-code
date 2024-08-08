package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Instruction struct {
	Id  int
	Bot int

	LowType string
	LowId   int

	HighType string
	HighId   int
}

var valueRegex = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
var instructionRegex = regexp.MustCompile(`bot (\d+) gives low to (output|bot) (\d+) and high to (output|bot) (\d+)`)

func insertOrUpdate(list map[int][]int, id, value int) {
	if _, exists := list[id]; exists {
		list[id] = append(list[id], value)
	} else {
		list[id] = []int{value}
	}
}

func getBots(lines []string) map[int][]int {
	bots := make(map[int][]int)

	for _, line := range lines {
		if !strings.HasPrefix(line, "value") {
			continue
		}

		match := valueRegex.FindStringSubmatch(line)
		value, _ := strconv.Atoi(match[1])
		botId, _ := strconv.Atoi(match[2])

		insertOrUpdate(bots, botId, value)
	}

	return bots
}

func getInstructions(lines []string) []Instruction {
	response := make([]Instruction, 0)

	for index, line := range lines {
		if strings.HasPrefix(line, "value") || line == "" {
			continue
		}

		match := instructionRegex.FindStringSubmatch(line)

		fromBot, _ := strconv.Atoi(match[1])
		toLowType := match[2]
		toLowId, _ := strconv.Atoi(match[3])
		toHighType := match[4]
		toHighId, _ := strconv.Atoi(match[5])

		response = append(response, Instruction{Id: index, Bot: fromBot, LowType: toLowType, LowId: toLowId, HighType: toHighType, HighId: toHighId})
	}

	return response
}

func distribute(
	bots map[int][]int,
	instructions []Instruction,
	processed []int,
	responsible int,
	outputs map[int][]int,
	chip1, chip2 int) (botId int, allOutputs map[int][]int) {

	if len(processed) == len(instructions) {
		return responsible, outputs
	}

	for _, instruction := range instructions {
		if slices.Contains(processed, instruction.Id) {
			continue
		}

		slices.Sort(bots[instruction.Bot])

		bot := bots[instruction.Bot]

		if len(bot) < 2 {
			continue
		}

		low := bot[0]
		high := bot[len(bot)-1]

		if (low == chip1 && high == chip2) || (low == chip2 && high == chip1) {
			responsible = instruction.Bot
		}

		if instruction.LowType == "bot" {
			insertOrUpdate(bots, instruction.LowId, low)
		} else {
			insertOrUpdate(outputs, instruction.LowId, low)
		}

		// remove low
		bots[instruction.Bot] = bot[1:]

		if instruction.HighType == "bot" {
			insertOrUpdate(bots, instruction.HighId, high)
		} else {
			insertOrUpdate(outputs, instruction.HighId, high)
		}

		// remove high
		bots[instruction.Bot] = bot[:len(bot)-1]

		processed = append(processed, instruction.Id)
	}

	return distribute(bots, instructions, processed, responsible, outputs, chip1, chip2)
}

func totalOutputs(outputs map[int][]int, selected []int) int {
	sum := 1
	for _, id := range selected {
		for _, chip := range outputs[id] {
			sum *= chip
		}
	}

	return sum
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	bots := getBots(lines)
	instructions := getInstructions(lines)

	botId, outputs := distribute(
		bots,
		instructions,
		make([]int, 0),
		-1,
		make(map[int][]int),
		17,
		61)

	totalOutputs := totalOutputs(outputs, []int{0, 1, 2})
	fmt.Printf("Part 1: %d, Part 2: %d\n", botId, totalOutputs)
}
