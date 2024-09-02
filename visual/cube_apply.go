package visual

import (
	cmn "github.com/mdesoeuv/rubik/common"
)

func (c *Cube) Apply(move cmn.Move) {
	rotateFace(&c.faces[move.Side], move.Rotation.Int())
	rotateCrown(c, move)
}

func rotateFace(face *Face, rotationCount int) {
	result := Face{}
	// Center never moves
	result.f[1][1] = face.f[1][1]

	for i := range faceCycle {
		index := len(faceCycle) + i + 2*rotationCount
		to, from := faceCycle[index%len(faceCycle)], faceCycle[i]
		result.f[to.Line()][to.Column()] = face.f[from.Line()][from.Column()]
	}
	*face = result
}

func rotateCrown(cube *Cube, move cmn.Move) {
	crown := crownCycle[move.Side]

	newCube := *cube

	for i := range crown {
		index := len(crown) + i + 3*move.Rotation.Int()
		to, from := crown[index%len(crown)], crown[i]
		newCube.Set(to, cube.Get(from))
	}

	*cube = newCube
}

var faceCycle = []cmn.FaceCoord{
	cmn.FaceCoord00,
	cmn.FaceCoord01,
	cmn.FaceCoord02,
	cmn.FaceCoord12,
	cmn.FaceCoord22,
	cmn.FaceCoord21,
	cmn.FaceCoord20,
	cmn.FaceCoord10,
}

var crownCycle = map[cmn.Side][12]cmn.CubeCoord{
	cmn.Front: {
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord21},
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord10},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord02},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord01},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord12},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord02},
	},
	cmn.Up: {
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord02},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord01},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord02},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord01},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord02},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord01},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord02},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord01},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord00},
	},
	cmn.Right: {
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord12},
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord02},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord10},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord12},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord02},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord12},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord02},
	},
	cmn.Down: {
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord21},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord21},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord21},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord21},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord22},
	},
	cmn.Left: {
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord10},
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord10},
		{Side: cmn.Front, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord10},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord12},
		{Side: cmn.Back, FaceCoord: cmn.FaceCoord02},
	},
	cmn.Back: {
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord02},
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord01},
		{Side: cmn.Up, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord00},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord10},
		{Side: cmn.Left, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord20},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord21},
		{Side: cmn.Down, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord22},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord12},
		{Side: cmn.Right, FaceCoord: cmn.FaceCoord02},
	},
}
