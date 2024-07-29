package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Cube struct {
	faces [6]Face
}

func (c *Cube) String() string {
	result := ""
	for side, face := range c.faces {
		result += fmt.Sprintf("Face: %c\n", sideNames[Side(side)])
		result += face.String()
		result += "\n"
	}
	return result
}

func NewCubeSolved() *Cube {
	cube := Cube{}

	for side := FirstSide; side <= LastSide; side++ {
		cube.faces[side] = NewFaceUniform(side)
	}

	return &cube
}

func (c *Cube) IsSolved() bool {
	for side, face := range c.faces {
		if !FaceIsUniform(face, side) {
			return false
		}
	}
	return true
}

type FaceCoord struct {
	line   int
	column int
}

type CubeCoord struct {
	side  Side
	coord FaceCoord
}

func rotateFace(face *Face, rotationCount int) {
	result := Face{}
	// Center never moves
	result.f[1][1] = face.f[1][1]

	cycle := []FaceCoord{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 2},
		{2, 1},
		{2, 0},
		{1, 0},
	}

	for i := range cycle {
		index := len(cycle) + i + 2*rotationCount
		to, from := cycle[index%len(cycle)], cycle[i]
		result.f[to.line][to.column] = face.f[from.line][from.column]
	}
	*face = result
}

func rotateCrown(cube *Cube, move Move) {
	crowns := map[Side][12]CubeCoord{
		Front: {
			{side: Up, coord: FaceCoord{line: 2, column: 0}},
			{side: Up, coord: FaceCoord{line: 2, column: 1}},
			{side: Up, coord: FaceCoord{line: 2, column: 2}},
			{side: Right, coord: FaceCoord{line: 0, column: 0}},
			{side: Right, coord: FaceCoord{line: 1, column: 0}},
			{side: Right, coord: FaceCoord{line: 2, column: 0}},
			{side: Down, coord: FaceCoord{line: 0, column: 0}},
			{side: Down, coord: FaceCoord{line: 0, column: 1}},
			{side: Down, coord: FaceCoord{line: 0, column: 2}},
			{side: Left, coord: FaceCoord{line: 2, column: 2}},
			{side: Left, coord: FaceCoord{line: 1, column: 2}},
			{side: Left, coord: FaceCoord{line: 0, column: 2}},
		},
		Up: {
			{side: Back, coord: FaceCoord{line: 0, column: 2}},
			{side: Back, coord: FaceCoord{line: 0, column: 1}},
			{side: Back, coord: FaceCoord{line: 0, column: 0}},
			{side: Right, coord: FaceCoord{line: 0, column: 2}},
			{side: Right, coord: FaceCoord{line: 0, column: 1}},
			{side: Right, coord: FaceCoord{line: 0, column: 0}},
			{side: Front, coord: FaceCoord{line: 0, column: 2}},
			{side: Front, coord: FaceCoord{line: 0, column: 1}},
			{side: Front, coord: FaceCoord{line: 0, column: 0}},
			{side: Left, coord: FaceCoord{line: 0, column: 2}},
			{side: Left, coord: FaceCoord{line: 0, column: 1}},
			{side: Left, coord: FaceCoord{line: 0, column: 0}},
		},
		Right: {
			{side: Up, coord: FaceCoord{line: 2, column: 2}},
			{side: Up, coord: FaceCoord{line: 1, column: 2}},
			{side: Up, coord: FaceCoord{line: 0, column: 2}},
			{side: Back, coord: FaceCoord{line: 0, column: 0}},
			{side: Back, coord: FaceCoord{line: 1, column: 0}},
			{side: Back, coord: FaceCoord{line: 2, column: 0}},
			{side: Down, coord: FaceCoord{line: 2, column: 0}},
			{side: Down, coord: FaceCoord{line: 1, column: 0}},
			{side: Down, coord: FaceCoord{line: 0, column: 0}},
			{side: Front, coord: FaceCoord{line: 2, column: 2}},
			{side: Front, coord: FaceCoord{line: 1, column: 2}},
			{side: Front, coord: FaceCoord{line: 0, column: 2}},
		},
		Down: {
			{side: Front, coord: FaceCoord{line: 2, column: 2}},
			{side: Front, coord: FaceCoord{line: 2, column: 1}},
			{side: Front, coord: FaceCoord{line: 2, column: 0}},
			{side: Left, coord: FaceCoord{line: 2, column: 2}},
			{side: Left, coord: FaceCoord{line: 2, column: 1}},
			{side: Left, coord: FaceCoord{line: 2, column: 0}},
			{side: Back, coord: FaceCoord{line: 2, column: 2}},
			{side: Back, coord: FaceCoord{line: 2, column: 1}},
			{side: Back, coord: FaceCoord{line: 2, column: 0}},
			{side: Right, coord: FaceCoord{line: 2, column: 2}},
			{side: Right, coord: FaceCoord{line: 2, column: 1}},
			{side: Right, coord: FaceCoord{line: 2, column: 0}},
		},
		Left: {
			{side: Up, coord: FaceCoord{line: 0, column: 0}},
			{side: Up, coord: FaceCoord{line: 1, column: 0}},
			{side: Up, coord: FaceCoord{line: 2, column: 0}},
			{side: Front, coord: FaceCoord{line: 0, column: 0}},
			{side: Front, coord: FaceCoord{line: 1, column: 0}},
			{side: Front, coord: FaceCoord{line: 2, column: 0}},
			{side: Down, coord: FaceCoord{line: 0, column: 2}},
			{side: Down, coord: FaceCoord{line: 1, column: 2}},
			{side: Down, coord: FaceCoord{line: 2, column: 2}},
			{side: Back, coord: FaceCoord{line: 2, column: 2}},
			{side: Back, coord: FaceCoord{line: 1, column: 2}},
			{side: Back, coord: FaceCoord{line: 0, column: 2}},
		},
		Back: {
			{side: Up, coord: FaceCoord{line: 0, column: 2}},
			{side: Up, coord: FaceCoord{line: 0, column: 1}},
			{side: Up, coord: FaceCoord{line: 0, column: 0}},
			{side: Left, coord: FaceCoord{line: 0, column: 0}},
			{side: Left, coord: FaceCoord{line: 1, column: 0}},
			{side: Left, coord: FaceCoord{line: 2, column: 0}},
			{side: Down, coord: FaceCoord{line: 2, column: 2}},
			{side: Down, coord: FaceCoord{line: 2, column: 1}},
			{side: Down, coord: FaceCoord{line: 2, column: 0}},
			{side: Right, coord: FaceCoord{line: 2, column: 2}},
			{side: Right, coord: FaceCoord{line: 1, column: 2}},
			{side: Right, coord: FaceCoord{line: 0, column: 2}},
		},
	}

	crown := crowns[move.Side]

	newCube := *cube

	// TODO: Handle anti-clockwise
	for i := range crown {
		from, to := crown[i], crown[(i+3)%len(crown)]
		*newCube.get(to) = *cube.get(from)
	}

	*cube = newCube
}

func (c *Cube) get(coord CubeCoord) *Side {
	return &c.faces[coord.side].f[coord.coord.line][coord.coord.column]
}

func (c *Cube) apply(move Move) {
	// TODO: Handle count an Clockwise

	// rotate face it self
	rotateFace(&c.faces[move.Side], move.NumRotations)

	rotateCrown(c, move)
	return
}

func (c *Cube) isSolved() bool {
	seenSides := [SideCount]bool{}
	for _, face := range c.faces {
		var currentSide *Side = nil
		for _, line := range face.f {
			for _, side := range line {
				if currentSide != nil && side != *currentSide {
					return false
				}
				currentSide = &side
			}
		}
		if seenSides[*currentSide] {
			return false
		}
		seenSides[*currentSide] = true
	}
	return true
}

func (c *Cube) print() {

	lines := [9]string{}
	emptyLine := "         "
	lines[0] = emptyLine + " " + c.faces[Up].FaceGetLineString(0)
	lines[1] = emptyLine + " " + c.faces[Up].FaceGetLineString(1)
	lines[2] = emptyLine + " " + c.faces[Up].FaceGetLineString(2) + "\n"

	lines[3] = c.faces[Left].FaceGetLineString(0) + " " + c.faces[Front].FaceGetLineString(0) + " " + c.faces[Right].FaceGetLineString(0) + " " + c.faces[Back].FaceGetLineString(0)
	lines[4] = c.faces[Left].FaceGetLineString(1) + " " + c.faces[Front].FaceGetLineString(1) + " " + c.faces[Right].FaceGetLineString(1) + " " + c.faces[Back].FaceGetLineString(1)
	lines[5] = c.faces[Left].FaceGetLineString(2) + " " + c.faces[Front].FaceGetLineString(2) + " " + c.faces[Right].FaceGetLineString(2) + " " + c.faces[Back].FaceGetLineString(2) + "\n"

	lines[6] = emptyLine + " " + c.faces[Down].FaceGetLineString(0)
	lines[7] = emptyLine + " " + c.faces[Down].FaceGetLineString(1)
	lines[8] = emptyLine + " " + c.faces[Down].FaceGetLineString(2)

	for _, line := range lines {
		fmt.Println(line)
	}
}

var theme = map[Side]func(a ...interface{}) string{
	Right: color.New(color.BgRed).SprintFunc(),
	Front: color.New(color.BgBlue).SprintFunc(),
	Back:  color.New(color.BgGreen).SprintFunc(),
	Up:    color.New(color.BgWhite).SprintFunc(),
	Down:  color.New(color.BgYellow).SprintFunc(),
	Left:  color.New(color.BgMagenta).SprintFunc(),
}
