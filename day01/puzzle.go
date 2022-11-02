/* --- Day 1: No Time for a Taxicab ---
 * Santa's sleigh uses a very high-precision clock to guide its movements, and the clock's oscillator is regulated by
 * stars. Unfortunately, the stars have been stolen... by the Easter Bunny. To save Christmas, Santa needs you to
 * retrieve all fifty stars by December 25th.
 *
 * Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second
 * puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!
 *
 * You're airdropped near Easter Bunny Headquarters in a city somewhere. "Near", unfortunately, is as close as you can
 * get - the instructions on the Easter Bunny Recruiting Document the Elves intercepted start here, and nobody had time
 * to work them out further.
 *
 * The Document indicates that you should start at the given coordinates (where you just landed) and face North. Then,
 * follow the provided sequence: either turn left (L) or right (R) 90 degrees, then walk forward the given number of
 * blocks, ending at a new intersection.
 *
 * There's no time to follow such ridiculous instructions on foot, though, so you take a moment and work out the
 * destination. Given that you can only walk on the street grid of the city, how far is the shortest path to the
 * destination?
 *
 * For example:
 * Following R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away.
 * R2, R2, R2 leaves you 2 blocks due South of your starting position, which is 2 blocks away.
 * R5, L5, R5, R3 leaves you 12 blocks away.
 * How many blocks away is Easter Bunny HQ?
 *
 * --- Part Two ---
 * Then, you notice the instructions continue on the back of the Recruiting Document. Easter Bunny HQ is actually at the
 * first location you visit twice.
 *
 * For example, if your instructions are R8, R4, R4, R8, the first location you visit twice is 4 blocks away, due East.
 *
 * How many blocks away is the first location you visit twice?
 */

package day01

import (
	"aoc2016/utils"
	"strconv"
	"strings"
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

const (
	RIGHT = 1
	LEFT  = -1
)

type Instruction struct {
	turnDirection int
	distance      int
}

var turnDirections = map[uint8]int{
	'R': RIGHT,
	'L': LEFT,
}

func parseInstruction(text string) Instruction {
	turnDirection := turnDirections[text[0]]
	distance, err := strconv.Atoi(text[1:])
	if err != nil {
		panic("Direction is not a number")
	}
	return Instruction{turnDirection, distance}
}

func readInstructions() []Instruction {
	line := utils.ReadLines("assets/input01.txt")[0]
	var instructions []Instruction
	for _, str := range strings.Split(line, ", ") {
		instructions = append(instructions, parseInstruction(str))
	}
	return instructions
}

func Solution1() int {
	santa := new(Santa)
	for _, instruction := range readInstructions() {
		santa.process(instruction)
	}
	return santa.position.distanceFrom(0, 0)
}

func Solution2() int {
	visited := make(map[Coords]bool)
	santa := new(Santa)
	for _, instruction := range readInstructions() {
		for _, square := range santa.process(instruction) {
			if alreadyVisited, _ := visited[square]; alreadyVisited {
				return square.distanceFrom(0, 0)
			}
			visited[square] = true
		}
	}
	return -1
}
