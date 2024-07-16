package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getPasswordPart1(doorId *string, numerator int, password string) string {
	fmt.Printf("Decoding part 1 password: %s \n", password)
	if len(password) == 8 {
		return password
	}

	for i := numerator; ; i++ {
		hash := getMD5Hash(fmt.Sprintf("%s%d", *doorId, i))

		if hash[:5] == "00000" {
			return getPasswordPart1(doorId, i+1, fmt.Sprintf("%s%s", password, string(hash[5])))
		}
	}
}

func getPasswordPart2(doorId string) string {
	password := map[int]rune{
		0: '_',
		1: '_',
		2: '_',
		3: '_',
		4: '_',
		5: '_',
		6: '_',
		7: '_',
	}

	getPassword := func() string {
		passwordStr := make([]rune, 8)
		for i := 0; i < 8; i++ {
			passwordStr[i] = password[i]
		}
		return string(passwordStr)
	}

	isPasswordSet := func() bool {
		for i := 0; i < 8; i++ {
			if password[i] == '_' {
				return false
			}
		}
		return true
	}

	for i := 0; ; i++ {
		hash := getMD5Hash(fmt.Sprintf("%s%d", doorId, i))

		if hash[:5] == "00000" {
			position := hash[5]
			if position >= '0' && position <= '7' {
				position, _ := strconv.Atoi(string(position))
				if password[position] == '_' {
					password[position] = rune(hash[6])
					fmt.Printf("Decoding part 2 password: %s \n", getPassword())

					if isPasswordSet() {
						break
					}

				}
			}
		}
	}

	return getPassword()
}

func main() {
	doorId := "abbhdwsy"
	passwordPart1 := getPasswordPart1(&doorId, 0, "")
	passwordPart2 := getPasswordPart2(doorId)

	fmt.Printf("Password part 1: %s, part 2: %s\n", passwordPart1, passwordPart2)
}
