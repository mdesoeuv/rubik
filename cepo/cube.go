package cepo

import (
	cmn "github.com/mdesoeuv/rubik/common"
)

func (s *Solver) Solve(cube cmn.Cube) []cmn.Move {
	switch cube := cube.(type) {
	case *Cube:
		return s.ToG4(cube)
	default:
		panic("invalid cube")
	}
}

func (c *Cube) NewSolver() cmn.Solver {
	return NewSolver()
}

type Cube struct {
	edgeOrientations   EdgeOrientations
	edgePermutation    EdgePermutation
	cornerOrientations CornerOrientations
	cornerPermutation  CornerPermutation
}

func NewCubeSolved() *Cube {
	return &Cube{
		edgeOrientations:   NewEdgeOrientationsSolved(),
		edgePermutation:    NewEdgePermutationSolved(),
		cornerOrientations: NewCornerOrientationsSolved(),
		cornerPermutation:  NewCornerPermutationSolved(),
	}
}

func (c *Cube) Apply(m cmn.Move) {
	c.edgeOrientations.Apply(m)
	c.edgePermutation.Apply(m)
	c.cornerOrientations.Apply(m)
	c.cornerPermutation.Apply(m)
}

func (c *Cube) Get(coord cmn.CubeCoord) cmn.Side {
	if coord.FaceCoord.IsEdge() {
		edgeFace := cmn.EdgeFaceMap[coord]
		originalIndex := c.edgePermutation.Get(edgeFace.Index)
		originalCoords := cmn.EdgeIndexMap[originalIndex]
		turned := c.edgeOrientations.Get(originalIndex)
		if (edgeFace.FaceNmb == 0) != turned {
			return originalCoords.A.Side
		} else {
			return originalCoords.B.Side
		}
	} else {
		// cornerIndex := CornerCoordMap[coord]
		// originalIndex := c.cornerPermutation.get(cornerIndex)
		// originalCoords := CornerIndexMap[originalIndex]
		// TODO: Implement
		panic("TODO")
	}
}

func (c *Cube) IsSolved() bool {
	return (c.edgeOrientations.IsSolved() &&
		c.edgePermutation.IsSolved() &&
		c.cornerOrientations.IsSolved() &&
		c.cornerPermutation.IsSolved())
}

func (c *Cube) IsG1() bool {
	return c.edgeOrientations.IsSolved()
}

func (c *Cube) IsG2AssumingG1() bool {
	return (c.cornerOrientations.IsSolved() &&
		c.edgePermutation.FUBDInCorrectSlice())
}

func (c *Cube) IsG2() bool {
	return c.IsG1() && c.IsG2AssumingG1()
}

func (s *Solver) IsG3AssumingG2(c *Cube) bool {
	_, isG3CornerPermutation := s.G3CornerHeuristicTable[c.cornerPermutation]
	_, isG3EdgePermutation := s.G3EdgeHeuristicTable[c.edgePermutation]
	return isG3CornerPermutation && isG3EdgePermutation
}

func (s *Solver) IsG3(c *Cube) bool {
	return c.IsG2() && s.IsG3AssumingG2(c)
}

func (c *Cube) IsG4AssumingG3() bool {
	return (c.edgePermutation.IsSolved() &&
		c.cornerPermutation.IsSolved())
}

func (c *Cube) IsG4() bool {
	return c.IsSolved()
}

func (c *Cube) Blueprint() string {
	panic("unimplemented")
}

func (c *Cube) Clone() cmn.Cube {
	newCube := *c
	return &newCube
}

func (c Cube) Equal(o Cube) bool {
	return (c.edgeOrientations == o.edgeOrientations &&
		c.edgePermutation == o.edgePermutation &&
		c.cornerOrientations == o.cornerOrientations &&
		c.cornerPermutation == o.cornerPermutation)
}
