package main

import (
	"aoc2016/day01"
	"aoc2016/day02"
	"aoc2016/day03"
	"aoc2016/day04"
	"aoc2016/day05"
	"aoc2016/day06"
	"aoc2016/day07"
	"fmt"
)

func main() {
	fmt.Println("Day 1")
	fmt.Println(" - How many blocks away is Easter Bunny HQ?", day01.Solution1())
	fmt.Println(" - How many blocks away is the first location you visit twice?", day01.Solution2())
	fmt.Println("Day 2")
	fmt.Println(" - What is the bathroom code?", day02.Solution1())
	fmt.Println(" - What is the correct bathroom code?", day02.Solution2())
	fmt.Println("Day 3")
	fmt.Println(" - How many of the listed triangles are possible?", day03.Solution1())
	fmt.Println(" - Instead reading by columns, how many of the listed triangles are possible?", day03.Solution2())
	fmt.Println("Day 4")
	fmt.Println(" - What is the sum of the sector IDs of the real rooms?", day04.Solution1())
	fmt.Println(" - What is the sector ID of the room where North Pole objects are stored?", day04.Solution2())
	fmt.Println("Day 5")
	fmt.Println(" - Given the actual Door ID, what is the password?", day05.Solution1())
	fmt.Println(" - Given the actual Door ID and this new method, what is the password?", day05.Solution2())
	fmt.Println("Day 6")
	fmt.Println(" - What is the error-corrected version of the message being sent?", day06.Solution1())
	fmt.Println(" - What is the original message?", day06.Solution2())
	fmt.Println("Day 7")
	fmt.Println(" - How many IPs in your puzzle input support TLS?", day07.Solution1())
	fmt.Println(" - How many IPs in your puzzle input support SSL?", day07.Solution2())
}
