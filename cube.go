package main

type Color = int
const (
	Red Color = 0
	Blue Color = iota
	Green Color = iota
	White Color = iota
	Yellow Color = iota
	Orange Color = iota
	ColorCount int = iota
)

type Cube struct {
	faces [6]Face
}

type Face = [3][3]Side

func NewFaceUniform(side Side) (face Face) {
	for i, line := range face {
		for j := range line {
			face[i][j] = side
		}
	}
	return
}

func FaceEqual(a, b Face) bool {
	for line, _ := range a {
		for column, _ := range a[line] {
			if a[line][column] != b[line][column] {
				return false
			}
		}
	}
	return true
}

func NewCubeSolved() *Cube {
	cube := Cube{}

	for side := FirstSide; side <= LastSide; side++ {
		cube.faces[side] = NewFaceUniform(side)
	}

	return &cube
}

func FaceIsUniform(face Face, side Side) bool {
	for line := range face {
		for column := range face[line] {
			if face[line][column] != side {
				return false
			}
		}
	}
	return true
}

func (c *Cube)IsSolved() bool {
	for side, face := range c.faces {
		if !FaceIsUniform(face, side) {
			return false
		}
	}
	return true
}

type Coord struct {
	line int
	column int
}

func rotateFace(face Face, clockWise bool) (result Face) {
	// Center never moves
	result[1][1] = face[1][1]

	// cycle := []Coord {
	// 	{0, 0},
	// 	{0, 1},
	// 	{0, 2},
	// 	{1, 2},
	// 	{2, 2},
	// 	{2, 1},
	// 	{2, 0},
	// 	{1, 0},
	// }

	if clockWise {
		// for i, _ := range cycle {
		// 	from, to := cycle[i], cycle[(i + 1) % len(cycle)];
		// 	result[from.line][from.column] = face[to.line][to.column]
		// }
		result[0][0] = face[0][1]
		result[0][1] = face[0][2]
		result[0][2] = face[1][2]
		result[1][2] = face[2][2]
		result[2][2] = face[2][1]
		result[2][1] = face[2][0]
		result[2][0] = face[1][0]
		result[1][0] = face[0][0]
	} else {
		result[0][0] = face[0][0]
		result[0][1] = face[0][1]
		result[0][2] = face[0][2]
		result[1][2] = face[1][2]
		result[2][2] = face[2][2]
		result[2][1] = face[2][1]
		result[2][0] = face[2][0]
		result[1][0] = face[0][0]
	}
	return
}

func (c *Cube)apply(move Move) {
	face := &c.faces[move.Side]

	// rotate face it self
	*face = rotateFace(*face, move.Clockwise)

	// TODO: rotate crown
	return 
}

func (c *Cube)isSolved() bool {
	seenColors := [ColorCount]bool{}
	for _, face := range c.faces {
		var currentColor *Color = nil
		for _, line := range face {
			for _, color := range line {
				if currentColor != nil  && color != *currentColor {
					return false
				}
				currentColor = &color
			}
		}
		if seenColors[*currentColor] {
			return false
		}
		seenColors[*currentColor] = true
	}
	return true
}
