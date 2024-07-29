package main

import (
	"fmt"
)

type Cube struct {
	faces [6]Face
}

type CubeCoord struct {
	side      Side
	faceCoord FaceCoord
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

func (c *Cube) get(coord CubeCoord) *Side {
	return &c.faces[coord.side].f[coord.faceCoord.line][coord.faceCoord.column]
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
