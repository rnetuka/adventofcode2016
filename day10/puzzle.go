package day10

import (
	"aoc2016/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Print() {
	solution1, solution2 := Solution()
	fmt.Println("Day 10")
	fmt.Println(" - What is the number of the bot?", solution1)
	fmt.Println(" - What do you get if you multiply together the values of one chip in each of outputs?", solution2)
}

func Solution() (int, int) {
	instructions := utils.ReadLines("assets/input10.txt")
	comparingBotId := -1

	var listener ComparisonListener = func(bot *Bot, low int, high int) {
		if low == 17 && high == 61 {
			comparingBotId = bot.id
		}
	}

	factory := NewFactory()
	factory.listener = &listener

	for _, instruction := range instructions {
		re := regexp.MustCompile("value ([0-9]+) goes to bot ([0-9]+)")
		find := re.FindStringSubmatch(instruction)
		if find != nil {
			microchip, _ := strconv.Atoi(find[1])
			id, _ := strconv.Atoi(find[2])
			factory.Bot(id).PickUp(microchip)
		}

		re = regexp.MustCompile("bot ([0-9]+) gives low to (bot|output) [0-9]+ and high to (bot|output) [0-9]+")
		find = re.FindStringSubmatch(instruction)
		if find != nil {
			id, _ := strconv.Atoi(find[1])
			bot := factory.Bot(id)
			if len(bot.microchips) == 2 {
				// Bot does have enough microchips, process the instruction immediately
				bot.Process(instruction, factory)
			} else {
				// Bot doesn't have enough microchips, queue this instruction
				bot.Queue(instruction)
			}
		}
	}

	// Proceed with all waiting tasks
	for true {
		if factory.WorkFinished() {
			break
		}
		for _, bot := range factory.bots {
			bot.ResumeWork(factory)
		}
	}

	outputResult := 1
	outputResult *= factory.Output(0).microchips[0]
	outputResult *= factory.Output(1).microchips[0]
	outputResult *= factory.Output(2).microchips[0]

	return comparingBotId, outputResult
}