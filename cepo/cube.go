package cepo

import (
	cmn "github.com/mdesoeuv/rubik/common"
)

type Cube struct {
	// TODO: Conbine edge description in a single uint64
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

func (c *Cube) IsG3AssumingG2() bool {
	return (c.edgePermutation.AllInCorrectSlice() &&
		c.cornerPermutation.AllInCorrectOrbit())
}

func (c *Cube) IsG3() bool {
	return c.IsG2() && c.IsG3AssumingG2()
}

func (c *Cube) IsG4AssumingG3() bool {
	return (c.edgePermutation.IsSolved() &&
		c.cornerPermutation.IsSolved())
}

func (c *Cube) IsG4() bool {
	return c.IsSolved()
}

// TODO: Improve precision
func (c *Cube) distanceToG2InG1() int {
	coDistance := c.cornerOrientations.Distance()
	edgeDistance := c.edgePermutation.FUBDCorrectSliceDistance()
	if coDistance > edgeDistance {
		return coDistance
	} else {
		return edgeDistance
	}
}

func (c *Cube) distanceToG3InG2() int {
	// TODO: the code \(^o^)/
	return 0
}

func (c *Cube) distanceToG4InG3() int {
	epDistance := c.edgePermutation.Distance()
	cpDistance := c.cornerPermutation.Distance()
	if epDistance > cpDistance {
		return epDistance
	} else {
		return cpDistance
	}
}

func (c Cube) Equal(o Cube) bool {
	return (c.edgeOrientations == o.edgeOrientations &&
		c.edgePermutation == o.edgePermutation &&
		c.cornerOrientations == o.cornerOrientations &&
		c.cornerPermutation == o.cornerPermutation)
}
