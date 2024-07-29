package main

func (c *Cube) apply(move Move) {
	rotateFace(&c.faces[move.Side], move.NumRotations)
	rotateCrown(c, move)
	return
}

func rotateFace(face *Face, rotationCount int) {
	result := Face{}
	// Center never moves
	result.f[1][1] = face.f[1][1]

	for i := range faceCycle {
		index := len(faceCycle) + i + 2*rotationCount
		to, from := faceCycle[index%len(faceCycle)], faceCycle[i]
		result.f[to.line][to.column] = face.f[from.line][from.column]
	}
	*face = result
}

func rotateCrown(cube *Cube, move Move) {

	crown := crownCycle[move.Side]

	newCube := *cube

	for i := range crown {
		index := len(crown) + i + 3*move.NumRotations
		to, from := crown[index%len(crown)], crown[i]
		*newCube.get(to) = *cube.get(from)
	}

	*cube = newCube
}

var faceCycle = []FaceCoord{
	{0, 0},
	{0, 1},
	{0, 2},
	{1, 2},
	{2, 2},
	{2, 1},
	{2, 0},
	{1, 0},
}

var crownCycle = map[Side][12]CubeCoord{
	Front: {
		{side: Up, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Up, faceCoord: FaceCoord{line: 2, column: 1}},
		{side: Up, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Right, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Right, faceCoord: FaceCoord{line: 1, column: 0}},
		{side: Right, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Down, faceCoord: FaceCoord{line: 0, column: 2}},
		{side: Down, faceCoord: FaceCoord{line: 0, column: 1}},
		{side: Down, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Left, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Left, faceCoord: FaceCoord{line: 1, column: 2}},
		{side: Left, faceCoord: FaceCoord{line: 0, column: 2}},
	},
	Up: {
		{side: Back, faceCoord: FaceCoord{line: 0, column: 2}},
		{side: Back, faceCoord: FaceCoord{line: 0, column: 1}},
		{side: Back, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Right, faceCoord: FaceCoord{line: 0, column: 2}},
		{side: Right, faceCoord: FaceCoord{line: 0, column: 1}},
		{side: Right, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Front, faceCoord: FaceCoord{line: 0, column: 2}},
		{side: Front, faceCoord: FaceCoord{line: 0, column: 1}},
		{side: Front, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Left, faceCoord: FaceCoord{line: 0, column: 2}},
		{side: Left, faceCoord: FaceCoord{line: 0, column: 1}},
		{side: Left, faceCoord: FaceCoord{line: 0, column: 0}},
	},
	Right: {
		{side: Up, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Up, faceCoord: FaceCoord{line: 1, column: 2}},
		{side: Up, faceCoord: FaceCoord{line: 0, column: 2}},
		{side: Back, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Back, faceCoord: FaceCoord{line: 1, column: 0}},
		{side: Back, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Down, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Down, faceCoord: FaceCoord{line: 1, column: 2}},
		{side: Down, faceCoord: FaceCoord{line: 0, column: 2}},
		{side: Front, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Front, faceCoord: FaceCoord{line: 1, column: 2}},
		{side: Front, faceCoord: FaceCoord{line: 0, column: 2}},
	},
	Down: {
		{side: Front, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Front, faceCoord: FaceCoord{line: 2, column: 1}},
		{side: Front, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Left, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Left, faceCoord: FaceCoord{line: 2, column: 1}},
		{side: Left, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Back, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Back, faceCoord: FaceCoord{line: 2, column: 1}},
		{side: Back, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Right, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Right, faceCoord: FaceCoord{line: 2, column: 1}},
		{side: Right, faceCoord: FaceCoord{line: 2, column: 0}},
	},
	Left: {
		{side: Up, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Up, faceCoord: FaceCoord{line: 1, column: 0}},
		{side: Up, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Front, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Front, faceCoord: FaceCoord{line: 1, column: 0}},
		{side: Front, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Down, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Down, faceCoord: FaceCoord{line: 1, column: 0}},
		{side: Down, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Back, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Back, faceCoord: FaceCoord{line: 1, column: 2}},
		{side: Back, faceCoord: FaceCoord{line: 0, column: 2}},
	},
	Back: {
		{side: Up, faceCoord: FaceCoord{line: 0, column: 2}},
		{side: Up, faceCoord: FaceCoord{line: 0, column: 1}},
		{side: Up, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Left, faceCoord: FaceCoord{line: 0, column: 0}},
		{side: Left, faceCoord: FaceCoord{line: 1, column: 0}},
		{side: Left, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Down, faceCoord: FaceCoord{line: 2, column: 0}},
		{side: Down, faceCoord: FaceCoord{line: 2, column: 1}},
		{side: Down, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Right, faceCoord: FaceCoord{line: 2, column: 2}},
		{side: Right, faceCoord: FaceCoord{line: 1, column: 2}},
		{side: Right, faceCoord: FaceCoord{line: 0, column: 2}},
	},
}
