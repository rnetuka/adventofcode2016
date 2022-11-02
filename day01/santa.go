package day01

type Santa struct {
	position        Coords
	facingDirection int
}

func (santa *Santa) turn(turnDirection int) {
	direction := santa.facingDirection + turnDirection
	if direction == -1 {
		direction = 3
	}
	if direction == 4 {
		direction = 0
	}
	santa.facingDirection = direction
}

func (santa *Santa) move(distance int) []Coords {
	var visited []Coords
	vector := moveVectors[santa.facingDirection]
	for i := 0; i < distance; i++ {
		santa.position.add(vector)
		visited = append(visited, santa.position)
	}
	return visited
}

func (santa *Santa) process(instruction Instruction) []Coords {
	santa.turn(instruction.turnDirection)
	return santa.move(instruction.distance)
}
