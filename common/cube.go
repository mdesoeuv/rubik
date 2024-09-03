package common

import "math/rand/v2"

type Cube interface {
	IsSolved() bool
	Apply(Move)
	Get(CubeCoord) Side
	Solve() []Move
	Clone() Cube
	Blueprint() string
}

type CubeCoord struct {
	Side      Side
	FaceCoord FaceCoord
}

type CornerCoords struct {
	A, B, C CubeCoord
}

type EdgeCoords struct {
	A, B CubeCoord
}

func Shuffle(cube Cube, r *rand.Rand, move_count int) {
	var previouSide *Side = nil
	for move := 0; move < move_count; move++ {
		m := AllMoves[r.IntN(len(AllMoves))]
		for previouSide != nil && m.Side == *previouSide {
			m = AllMoves[r.IntN(len(AllMoves))]
		}
		cube.Apply(m)
		previouSide = &m.Side
	}
}

func ApplySequence(cube Cube, moves []Move) {
	for _, move := range moves {
		cube.Apply(move)
	}
}
