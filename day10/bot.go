package day10

import (
	"aoc2016/utils"
	"regexp"
	"strconv"
)

type ComparisonListener func(*Bot, int, int)

type Bot struct {
	id			 int
	microchips 	 []int
	instructions []string
	listener     *ComparisonListener
}

func (bot *Bot) Queue(instruction string) {
	bot.instructions = append(bot.instructions, instruction)
}

func (bot *Bot) PickUp(microchip int) {
	bot.microchips = append(bot.microchips, microchip)
}

func (bot *Bot) Process(instruction string, factory *Factory) {
	if len(bot.microchips) == 2 {
		re := regexp.MustCompile("bot [0-9]+ gives low to (bot|output) ([0-9]+) and high to (bot|output) ([0-9]+)")
		find := re.FindStringSubmatch(instruction)
		if find != nil {
			targetA := find[1]
			idA, _ := strconv.Atoi(find[2])
			targetB := find[3]
			idB, _ := strconv.Atoi(find[4])

			microchipA := bot.microchips[0]
			microchipB := bot.microchips[1]
			bot.microchips = make([]int, 0)

			low := utils.Min(microchipA, microchipB)
			high := utils.Max(microchipA, microchipB)

			if bot.listener != nil {
				(*bot.listener)(bot, low, high)
			}

			if targetA == "output" {
				factory.Output(idA).Add(low)
			}
			if targetB == "output" {
				factory.Output(idB).Add(high)
			}

			if targetA == "bot" {
				factory.Bot(idA).PickUp(low)
				factory.Bot(idA).ResumeWork(factory)
			}
			if targetB == "bot" {
				factory.Bot(idB).PickUp(high)
				factory.Bot(idB).ResumeWork(factory)
			}
			return
		}
	}
}

func (bot *Bot) ResumeWork(factory *Factory) {
	if len(bot.instructions) > 0 && len(bot.microchips) == 2 {
		instruction := bot.instructions[0]
		bot.instructions = bot.instructions[1:]
		bot.Process(instruction, factory)
	}
}