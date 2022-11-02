package day01

type Vector struct {
	dx, dy int
}

var moveVectors = map[int]Vector{
	NORTH: {0, -1},
	EAST:  {1, 0},
	SOUTH: {0, 1},
	WEST:  {-1, 0},
}
