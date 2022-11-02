package day03

import (
	"aoc2016/utils"
	"regexp"
	"strconv"
)

func checkTriangle(a, b, c int) int {
	if (a + b) <= c {
		return 0
	}
	if (a + c) <= b {
		return 0
	}
	if (b + c) <= a {
		return 0
	}
	return 1
}

func getSides(line string) (int, int, int) {
	re, _ := regexp.Compile("[0-9]+")
	parts := re.FindAllString(line, -1)
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	c, _ := strconv.Atoi(parts[2])
	return a, b, c
}

func Solution1() int {
	triangles := 0
	for _, line := range utils.ReadLines("assets/input03.txt") {
		a, b, c := getSides(line)
		triangles += checkTriangle(a, b, c)
	}
	return triangles
}

func Solution2() int {
	triangles := 0
	lines := utils.ReadLines("assets/input03.txt")
	for i := 0; i < len(lines); i += 3 {
		a1, a2, a3 := getSides(lines[i])
		b1, b2, b3 := getSides(lines[i + 1])
		c1, c2, c3 := getSides(lines[i + 2])
		triangles += checkTriangle(a1, b1, c1)
		triangles += checkTriangle(a2, b2, c2)
		triangles += checkTriangle(a3, b3, c3)
	}
	return triangles
}
