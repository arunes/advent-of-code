package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	Name      string
	CleanName string
	SectorId  int
	Checksum  string
}

func (room Room) String() string {
	return fmt.Sprintf("Name: %s, SectorId: %d, Checksum: %s", room.CleanName, room.SectorId, room.Checksum)
}

func getRoom(line string) Room {
	parts := strings.Split(line, "-")
	name := strings.Join(parts[0:len(parts)-1], "-")
	cleanName := strings.ReplaceAll(name, "-", "")
	lastParts := strings.Split(parts[len(parts)-1], "[")
	sectorId, _ := strconv.Atoi(lastParts[0])
	checksum := strings.Trim(lastParts[1], "]")

	return Room{Name: name, CleanName: cleanName, SectorId: sectorId, Checksum: checksum}
}

func isRealRoom(room Room) bool {
	letters := []rune(room.CleanName)

	sort.Slice(letters, func(i, j int) bool {
		left := strings.Count(room.CleanName, string(letters[i]))
		right := strings.Count(room.CleanName, string(letters[j]))

		if left != right {
			return left > right
		} else {
			return letters[i] < letters[j]
		}
	})

	letters = slices.Compact(letters)[0:len(room.Checksum)]
	checksum := string(letters)
	return checksum == room.Checksum
}

func decodeName(name string, sectorId int) string {
	decoded := []rune{}
	for _, char := range name {
		if char == '-' {
			decoded = append(decoded, ' ')
			continue
		}

		newPos := ((int(char) - 97 + sectorId) % 26) + 97
		decoded = append(decoded, rune(newPos))
	}

	return string(decoded)
}

func main() {
	file, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(file), "\n")

	northPoleSectorId := 0
	sectorIdsSum := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}

		room := getRoom(line)
		if !isRealRoom(room) {
			continue
		}

		if northPoleSectorId == 0 {
			sectorIdsSum += room.SectorId
			name := decodeName(room.Name, room.SectorId)
			if strings.Contains(name, "north") {
				northPoleSectorId = room.SectorId
			}
		}
	}

	fmt.Printf("Sum of real room sector ids: %d, North pole object storage sector id: %d\n", sectorIdsSum, northPoleSectorId)
}
