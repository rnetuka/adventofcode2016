package day02

import (
	"aoc2016/utils"
	"strings"
)

func Solution1() string {
	keypad := MakeKeypad(3)
	keypad.SetButton(0, 0, '1')
	keypad.SetButton(0, 1, '2')
	keypad.SetButton(0, 2, '3')
	keypad.SetButton(1, 0, '4')
	keypad.SetButton(1, 1, '5')
	keypad.SetButton(1, 2, '6')
	keypad.SetButton(2, 0, '7')
	keypad.SetButton(2, 1, '8')
	keypad.SetButton(2, 2, '9')
	var sequence strings.Builder
	button := '5'
	instructions := utils.ReadLines("assets/input02.txt")
	for _, instruction := range instructions {
		button = keypad.FindButton(button, instruction)
		sequence.WriteRune(button)
	}
	return sequence.String()
}

func Solution2() string {
	keypad := MakeKeypad(5)
	keypad.SetButton(0, 2, '1')
	keypad.SetButton(1, 1, '2')
	keypad.SetButton(1, 2, '3')
	keypad.SetButton(1, 3, '4')
	keypad.SetButton(2, 0, '5')
	keypad.SetButton(2, 1, '6')
	keypad.SetButton(2, 2, '7')
	keypad.SetButton(2, 3, '8')
	keypad.SetButton(2, 4, '9')
	keypad.SetButton(3, 1, 'A')
	keypad.SetButton(3, 2, 'B')
	keypad.SetButton(3, 3, 'C')
	keypad.SetButton(4, 2, 'D')
	var sequence strings.Builder
	button := '5'
	instructions := utils.ReadLines("assets/input02.txt")
	for _, instruction := range instructions {
		button = keypad.FindButton(button, instruction)
		sequence.WriteRune(button)
	}
	return sequence.String()
}
