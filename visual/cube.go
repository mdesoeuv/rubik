package visual

import (
	"fmt"
	"math/rand/v2"

	cmn "github.com/mdesoeuv/rubik/common"
)

var sideNames = map[cmn.Side]rune{cmn.Front: 'F', cmn.Back: 'B', cmn.Up: 'U', cmn.Down: 'D', cmn.Left: 'L', cmn.Right: 'R'}

type Cube struct {
	faces [6]Face
}

func (c *Cube) String() string {
	result := ""
	for side, face := range c.faces {
		result += fmt.Sprintf("Face: %c\n", sideNames[cmn.Side(side)])
		result += face.String()
		result += "\n"
	}
	return result
}

func NewCubeSolved() *Cube {
	cube := Cube{}

	for side := cmn.FirstSide; side <= cmn.LastSide; side++ {
		cube.faces[side] = NewFaceUniform(side)
	}

	return &cube
}

func (cube *Cube) Shuffle(r *rand.Rand, move_count int) {
	var previouSide *cmn.Side = nil
	for move := 0; move < move_count; move++ {
		m := cmn.AllMoves[r.IntN(len(cmn.AllMoves))]
		for previouSide != nil && m.Side == *previouSide {
			m = cmn.AllMoves[r.IntN(len(cmn.AllMoves))]
		}
		cube.Apply(m)
		previouSide = &m.Side
	}
}

func (c *Cube) Get(coord cmn.CubeCoord) cmn.Side {
	return c.faces[coord.Side].f[coord.FaceCoord.Line()][coord.FaceCoord.Column()]
}

func (c *Cube) Set(coord cmn.CubeCoord, side cmn.Side) {
	c.faces[coord.Side].f[coord.FaceCoord.Line()][coord.FaceCoord.Column()] = side
}

func (c *Cube) IsSolved() bool {
	for face_index, face := range c.faces {
		for _, line := range face.f {
			for _, side := range line {
				if side != cmn.Side(face_index) {
					return false
				}
			}
		}
	}
	return true
}

func (c *Cube) Clone() cmn.Cube {
	newCube := *c
	return &newCube
}
