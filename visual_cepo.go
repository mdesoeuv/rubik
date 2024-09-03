package main

import (
	"fmt"

	cepo "github.com/mdesoeuv/rubik/cepo"
	cmn "github.com/mdesoeuv/rubik/common"
	visual "github.com/mdesoeuv/rubik/visual"
)

type VisualCepo struct {
	Visual *visual.Cube
	Cepo   *cepo.Cube
}

func (pc *VisualCepo) IsSolved() bool {
	visualIsSolve := pc.Visual.IsSolved()
	cepoIsSolved := pc.Cepo.IsSolved()

	if visualIsSolve != cepoIsSolved {
		panic(fmt.Sprintf("Incoherence found:\nVisual: %v, Cepo: %v\n", visualIsSolve, cepoIsSolved))
	}
	return cepoIsSolved
}

func (pc *VisualCepo) Get(coord cmn.CubeCoord) cmn.Side {
	return pc.Visual.Get(coord)
}

func (pc *VisualCepo) Apply(move cmn.Move) {
	pc.Visual.Apply(move)
	pc.Cepo.Apply(move)
}

// TODO: implement real solver for CEPO
func (pc *VisualCepo) Solve() []cmn.Move {
	return pc.Cepo.ToG1()
}

func (pc *VisualCepo) Clone() cmn.Cube {
	newCepo := *pc.Cepo
	newVisual := *pc.Visual
	newVisualCepo := VisualCepo{Cepo: &newCepo, Visual: &newVisual}

	return &newVisualCepo
}

func (pc *VisualCepo) Blueprint() string {
	return pc.Visual.Blueprint()
}
