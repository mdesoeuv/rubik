package visual_cepo

import (
	"fmt"

	cepo "github.com/mdesoeuv/rubik/cepo"
	cmn "github.com/mdesoeuv/rubik/common"
	visual "github.com/mdesoeuv/rubik/visual"
)

type Cube struct {
	Visual *visual.Cube
	Cepo   *cepo.Cube
}

type Solver struct {
	CepoSolver *cepo.Solver
}

func (solver *Solver) Solve(cube cmn.Cube) []cmn.Move {
	switch cube := cube.(type) {
	case *Cube:
		return solver.CepoSolver.Solve(cube.Cepo)
	default:
		panic("invalid cube")
	}
}

func (c *Cube) NewSolver() cmn.Solver {
	return NewSolver()
}

func NewSolver() *Solver {
	return &Solver{
		CepoSolver: cepo.NewSolver(),
	}
}

func NewCubeSolved() *Cube {
	return &Cube{
		Cepo:   cepo.NewCubeSolved(),
		Visual: visual.NewCubeSolved(),
	}
}

func (c *Cube) IsSolved() bool {
	visualIsSolve := c.Visual.IsSolved()
	cepoIsSolved := c.Cepo.IsSolved()

	if visualIsSolve != cepoIsSolved {
		panic(fmt.Sprintf("Incoherence found:\nVisual: %v, Cepo: %v\n", visualIsSolve, cepoIsSolved))
	}
	return cepoIsSolved
}

func (c *Cube) Get(coord cmn.CubeCoord) cmn.Side {
	return c.Visual.Get(coord)
}

func (c *Cube) Apply(move cmn.Move) {
	c.Visual.Apply(move)
	c.Cepo.Apply(move)
}

func (c *Cube) Clone() cmn.Cube {
	newCepo := *c.Cepo
	newVisual := *c.Visual
	newVisualCepo := Cube{Cepo: &newCepo, Visual: &newVisual}

	return &newVisualCepo
}

func (c *Cube) Blueprint() string {
	return c.Visual.Blueprint()
}
