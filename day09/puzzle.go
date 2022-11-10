package day09

import (
	"aoc2016/utils"
	"strconv"
	"strings"
)

func isCompressed(message string) bool {
	return strings.Contains(message, "(")
}

func subsequence(message string, length int) string {
	sub := strings.Builder{}
	i := strings.IndexByte(message, ')')
	for j := i + 1; j < i + 1 + length; j++ {
		sub.WriteByte(message[j])
	}
	return sub.String()
}

func decompress(message string) string {
	/*if !isCompressed(message) {
		return message
	}
	b := strings.Builder{}
	j := strings.IndexByte(message, '(')
	k := strings.IndexByte(message, ')')
	marker := message[j+1:k]
	parts := strings.Split(marker, "x")
	subLength, _ := strconv.Atoi(parts[0])
	times, _ := strconv.Atoi(parts[1])

	for i := 0; i < j; i++ {
		b.WriteByte(message[i])
	}

	sub := subsequence(message, subLength)
	for t := 0; t < times; t++ {
		b.WriteString(sub)
	}

	for i := k + 1 + len(sub); i < len(message); i++ {
		b.WriteByte(message[i])
	}
	return decompress(b.String())*/

	b := strings.Builder{}

	insideMarker := false
	j := 0
	k := 0

	for i := 0; i < len(message); i++ {
		if message[i] == '(' {
			insideMarker = true
			j = i
		} else if message[i] == ')' {
			insideMarker = false
			k = i
			marker := message[j+1:k]
			parts := strings.Split(marker, "x")
			subLength, _ := strconv.Atoi(parts[0])
			times, _ := strconv.Atoi(parts[1])

			sub := strings.Builder{}
			for s := 0; s < subLength; s++ {
				sub.WriteByte(message[i + 1 + s])
			}
			for t := 0; t < times; t++ {
				b.WriteString(sub.String())
			}
			 i += subLength
		} else if !insideMarker {
			b.WriteByte(message[i])
		}
	}
	return b.String()
}

func decompress2(message string) string {
	if !isCompressed(message) {
		return message
	}
	b := strings.Builder{}
	j := strings.IndexByte(message, '(')
	k := strings.IndexByte(message, ')')
	marker := message[j+1:k]
	parts := strings.Split(marker, "x")
	subLength, _ := strconv.Atoi(parts[0])
	times, _ := strconv.Atoi(parts[1])

	for i := 0; i < j; i++ {
		b.WriteByte(message[i])
	}

	sub := subsequence(message, subLength)
	for t := 0; t < times; t++ {
		b.WriteString(sub)
	}

	for i := k + 1 + len(sub); i < len(message); i++ {
		b.WriteByte(message[i])
	}
	return decompress(b.String())
}

func Solution1() int {
	input := utils.Read("assets/input09.txt")
	message := decompress(input)
	length := 0
	for _, char := range message {
		if char != ' ' {
			length++
		}
	}
	return length
}

func Solution2() int {
	input := "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"
	message := decompress2(input)
	length := 0
	for _, char := range message {
		if char != ' ' {
			length++
		}
	}
	return length
}