package day04

import (
	"aoc2016/utils"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Room struct {
	name string
	id int
	checksum string
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func parseRoom(str string) Room {
	result := Room{}
	re, _ := regexp.Compile("([a-z]+\\-)+")
	name := re.FindString(str)
	result.name = name[:len(name)-1]
	re, _ = regexp.Compile("[0-9]+")
	result.id, _ = strconv.Atoi(re.FindString(str))
	re, _ = regexp.Compile("\\[.{5}\\]")
	checksum := re.FindString(str)
	result.checksum = checksum[1:len(checksum)-1]
	return result
}

func decrypt(roomName string, roomId int) string {
	result := strings.Builder{}
	for _, character := range roomName {
		if character == '-' {
			result.WriteRune(' ')
		} else {
			i := strings.IndexRune(alphabet, character)
			i = (i + roomId) % len(alphabet)
			result.WriteByte(alphabet[i])
		}
	}
	return result.String()
}

func countOccurrences(roomName string, letter rune) int {
	count := 0
	for _, char := range roomName {
		if char == letter {
			count++
		}
	}
	return count
}

func maxValue(letterOccurrences map[rune]int) rune {
	max := 0
	var result rune
	for letter, count := range letterOccurrences {
		if count > max {
			max = count
			result = letter
		} else if count == max {
			if letter < result {
				result = letter
			}
		}
	}
	return result
}

func fiveMostCommonLetters(roomName string) string {
	occurrences := make(map[rune]int)
	for _, letter := range roomName {
		if unicode.IsLetter(letter) {
			occurrences[letter] = countOccurrences(roomName, letter)
		}
	}
	mostCommon := strings.Builder{}
	for i := 0; i < 5; i++ {
		letter := maxValue(occurrences)
		mostCommon.WriteRune(letter)
		occurrences[letter] = 0
	}
	return mostCommon.String()
}

func readRooms() []Room {
	result := make([]Room, 100)
	for _, line := range utils.ReadLines("assets/input04.txt") {
		room := parseRoom(line)
		result = append(result, room)
	}
	return result
}

func Solution1() int {
	idSum := 0
	for _, room := range readRooms() {
		checksum := fiveMostCommonLetters(room.name)
		if checksum == room.checksum {
			idSum += room.id
		}
	}
	return idSum
}

func Solution2() int {
	for _, room := range readRooms() {
		decryptedName := decrypt(room.name, room.id)
		if strings.Contains(decryptedName, "north pole") || strings.Contains(decryptedName, "northpole") {
			return room.id
		}
	}
	return 0
}
