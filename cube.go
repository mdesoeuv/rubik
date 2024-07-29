package main

import (
	"fmt"

	"github.com/fatih/color"
)


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

func (c Cube) String() string {
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
	result.f[1][1] = face.f[1][1]

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
		result.f[0][0] = face.f[0][1]
		result.f[0][1] = face.f[0][2]
		result.f[0][2] = face.f[1][2]
		result.f[1][2] = face.f[2][2]
		result.f[2][2] = face.f[2][1]
		result.f[2][1] = face.f[2][0]
		result.f[2][0] = face.f[1][0]
		result.f[1][0] = face.f[0][0]
	} else {
		result.f[0][0] = face.f[0][0]
		result.f[0][1] = face.f[0][1]
		result.f[0][2] = face.f[0][2]
		result.f[1][2] = face.f[1][2]
		result.f[2][2] = face.f[2][2]
		result.f[2][1] = face.f[2][1]
		result.f[2][0] = face.f[2][0]
		result.f[1][0] = face.f[0][0]
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
		for _, line := range face.f {
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

func (c *Cube)print() {
	
	lines := [9]string{}
	emptyLine := "          "
	lines[0] = emptyLine + c.faces[Up].FaceGetLineString(0) + emptyLine + emptyLine
	lines[1] = emptyLine + c.faces[Up].FaceGetLineString(1) + emptyLine + emptyLine
	lines[2] = emptyLine + c.faces[Up].FaceGetLineString(2) + emptyLine + emptyLine + "\n"

	lines[3] = c.faces[Left].FaceGetLineString(0) + " " + c.faces[Front].FaceGetLineString(0) + " " + c.faces[Right].FaceGetLineString(0) + " " + c.faces[Back].FaceGetLineString(0)
	lines[4] = c.faces[Left].FaceGetLineString(1) + " " + c.faces[Front].FaceGetLineString(1) + " " + c.faces[Right].FaceGetLineString(1) + " " + c.faces[Back].FaceGetLineString(1)
	lines[5] = c.faces[Left].FaceGetLineString(2) + " " + c.faces[Front].FaceGetLineString(2) + " " + c.faces[Right].FaceGetLineString(2) + " " + c.faces[Back].FaceGetLineString(2) + "\n"
	
	lines[6] = emptyLine + c.faces[Down].FaceGetLineString(0) + emptyLine + emptyLine
	lines[7] = emptyLine + c.faces[Down].FaceGetLineString(1) + emptyLine + emptyLine
	lines[8] = emptyLine + c.faces[Down].FaceGetLineString(2) + emptyLine + emptyLine

	for _, line := range lines {
		fmt.Println(line)
	} 
}

var theme = map[Side]func(a ...interface{}) string{
	Right: color.New(color.BgRed).SprintFunc(),
	Front: color.New(color.BgBlue).SprintFunc(),
	Back: color.New(color.BgGreen).SprintFunc(),
	Up: color.New(color.BgWhite).SprintFunc(),
	Down: color.New(color.BgYellow).SprintFunc(),
	Left:color.New(color.BgMagenta).SprintFunc(),
}
