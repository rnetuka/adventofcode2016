package day02

type Keypad struct {
	size    int
	buttons [][]rune
}

func MakeKeypad(size int) Keypad {
	keypad := Keypad{}
	keypad.size = size
	keypad.buttons = make([][]rune, size)
	for i := range keypad.buttons {
		keypad.buttons[i] = make([]rune, size)
	}
	return keypad
}

func (keypad *Keypad) SetButton(row int, column int, button rune) {
	keypad.buttons[row][column] = button
}

func (keypad *Keypad) buttonAt(row, column int) rune {
	return keypad.buttons[row][column]
}

func (keypad *Keypad) coordinatesOf(button rune) (int, int) {
	for row := 0; row < keypad.size; row++ {
		for column := 0; column < keypad.size; column++ {
			if keypad.buttons[row][column] == button {
				return row, column
			}
		}
	}
	return -1, -1
}

func (keypad *Keypad) hasButtonAt(row, column int) bool {
	if row < 0 || row >= keypad.size || column < 0 || column >= keypad.size {
		return false
	}
	return keypad.buttons[row][column] != rune(0)
}

func (keypad *Keypad) FindButton(start rune, instruction string) rune {
	row, column := keypad.coordinatesOf(start)
	operations := make(map[rune]func() (int, int))
	operations['U'] = func() (int, int) { return row - 1, column }
	operations['D'] = func() (int, int) { return row + 1, column }
	operations['L'] = func() (int, int) { return row, column - 1 }
	operations['R'] = func() (int, int) { return row, column + 1 }
	for _, char := range instruction {
		operation := operations[char]
		newRow, newColumn := operation()
		if keypad.hasButtonAt(newRow, newColumn) {
			row, column = newRow, newColumn
		}
	}
	return keypad.buttonAt(row, column)
}
