package day06

import "aoc2016/utils"

func maxFrequentChar(characters map[rune]int) rune {
	max := 0
	var result rune
	for char, occurrences := range characters {
		if occurrences > max {
			max = occurrences
			result = char
		}
	}
	return result
}

func minFrequentChar(characters map[rune]int) rune {
	min := 10000
	var result rune
	for char, occurrences := range characters {
		if occurrences < min {
			min = occurrences
			result = char
		}
	}
	return result
}

func Solution1() string {
	messages := utils.ReadLines("assets/input06.txt")
	messageLength := len(messages[0])
	result := make([]rune, messageLength)
	for i := 0; i < messageLength; i++ {
		characters := make(map[rune]int)
		for _, message := range messages {
			char := rune(message[i])
			characters[char]++
		}
		result[i] = maxFrequentChar(characters)
	}
	return string(result)
}

func Solution2() string {
	messages := utils.ReadLines("assets/input06.txt")
	messageLength := len(messages[0])
	result := make([]rune, messageLength)
	for i := 0; i < messageLength; i++ {
		characters := make(map[rune]int)
		for _, message := range messages {
			char := rune(message[i])
			characters[char]++
		}
		result[i] = minFrequentChar(characters)
	}
	return string(result)
}
