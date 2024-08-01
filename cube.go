package main

import (
	"fmt"
	"math/rand/v2"
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

func (cube *Cube) Shuffle(r *rand.Rand, move_count int) {
	for move := 0; move < move_count; move++ {
		cube.apply(AllMoves[rand.IntN(len(AllMoves))])
	}
}

func (c *Cube) get(coord CubeCoord) *Side {
	return &c.faces[coord.side].f[coord.faceCoord.line][coord.faceCoord.column]
}

func (c *Cube) isSolved() bool {
	for face_index, face := range c.faces {
		for _, line := range face.f {
			for _, side := range line {
				if side != Side(face_index) {
					return false
				}
			}
		}
	}
	return true
}
