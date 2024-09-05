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
	EdgeOrientations   EdgeOrientations
	EdgePermutation    EdgePermutation
	CornerOrientations CornerOrientations
	CornerPermutation  CornerPermutation
}

func NewCubeSolved() *Cube {
	return &Cube{
		EdgeOrientations:   NewEdgeOrientationsSolved(),
		EdgePermutation:    NewEdgePermutationSolved(),
		CornerOrientations: NewCornerOrientationsSolved(),
		CornerPermutation:  NewCornerPermutationSolved(),
	}
}

func (c *Cube) Apply(m cmn.Move) {
	c.EdgeOrientations.Apply(m)
	c.EdgePermutation.Apply(m)
	c.CornerOrientations.Apply(m)
	c.CornerPermutation.Apply(m)
}

func (c *Cube) Get(coord cmn.CubeCoord) cmn.Side {
	if coord.FaceCoord.IsEdge() {
		edgeFace := cmn.EdgeFaceMap[coord]
		originalIndex := c.EdgePermutation.Get(edgeFace.Index)
		originalCoords := cmn.EdgeIndexMap[originalIndex]
		turned := c.EdgeOrientations.Get(originalIndex)
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
	return (c.EdgeOrientations.IsSolved() &&
		c.EdgePermutation.IsSolved() &&
		c.CornerOrientations.IsSolved() &&
		c.CornerPermutation.IsSolved())
}

func (c *Cube) IsG1() bool {
	return c.EdgeOrientations.IsSolved()
}

func (c *Cube) IsG2AssumingG1() bool {
	return (c.CornerOrientations.IsSolved() &&
		c.EdgePermutation.FUBDInCorrectSlice())
}

func (c *Cube) IsG2() bool {
	return c.IsG1() && c.IsG2AssumingG1()
}

func (s *Solver) IsG3AssumingG2(c *Cube) bool {
	_, isG3 := s.G3HeuristicTable[*c]
	return isG3
	// _, isG3CornerPermutation := s.G3CornerHeuristicTable[c.cornerPermutation]
	// _, isG3EdgePermutation := s.G3EdgeHeuristicTable[c.edgePermutation]
	// return isG3CornerPermutation && isG3EdgePermutation
}

func (s *Solver) IsG3(c *Cube) bool {
	return c.IsG2() && s.IsG3AssumingG2(c)
}

func (c *Cube) IsG4AssumingG3() bool {
	return (c.EdgePermutation.IsSolved() &&
		c.CornerPermutation.IsSolved())
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
	return (c.EdgeOrientations == o.EdgeOrientations &&
		c.EdgePermutation == o.EdgePermutation &&
		c.CornerOrientations == o.CornerOrientations &&
		c.CornerPermutation == o.CornerPermutation)
}
