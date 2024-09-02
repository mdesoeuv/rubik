package main

import (
	"fmt"

	cepo "github.com/mdesoeuv/rubik/cepo"
	cmn "github.com/mdesoeuv/rubik/common"
	visual "github.com/mdesoeuv/rubik/visual"
)

type PairCube struct {
	Visual *visual.Cube
	Cepo   *cepo.Cube
}

func (pc *PairCube) IsSolved() bool {
	visualIsSolve := pc.Visual.IsSolved()
	cepoIsSolved := pc.Cepo.IsSolved()

	if visualIsSolve != cepoIsSolved {
		panic(fmt.Sprintf("Incoherence found:\nVisual: %v, Cepo: %v\n", visualIsSolve, cepoIsSolved))
	}
	return cepoIsSolved
}

func (pc *PairCube) Get(coord cmn.CubeCoord) cmn.Side {
	return pc.Visual.Get(coord)
}

func (pc *PairCube) Apply(move cmn.Move) {
	pc.Visual.Apply(move)
	pc.Cepo.Apply(move)
}
