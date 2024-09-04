package cepo

import (
	cmn "github.com/mdesoeuv/rubik/common"
)

type Solver struct {
}

func (solver *Solver) Solve(cube cmn.Cube) []cmn.Move {
	switch cube := cube.(type) {
	case *Cube:
		return cube.ToG4()
	default:
		panic("invalid cube")
	}
}

func (c *Cube) NewSolver() cmn.Solver {
	return NewSolver()
}

func NewSolver() *Solver {
	return &Solver{}
}

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

var G3Cubes = MakeG3Cubes()

func (c *Cube) IsG3AssumingG2() bool {
	// return (c.edgePermutation.AllInCorrectSlice() &&
	// 	c.cornerPermutation.AllInCorrectOrbit())
	_, isG3 := G3Cubes[*c]
	return isG3
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

func (c *Cube) ToG2() []cmn.Move {
	cube := *c
	movesToG1 := cube.ToG1()
	if movesToG1 == nil {
		return nil
	}
	cmn.ApplySequence(&cube, movesToG1)
	movesToG2 := cube.ToG2AssumingG1()
	if movesToG2 == nil {
		return nil
	}
	return append(movesToG1, movesToG2...)
}

func (c *Cube) ToG3() []cmn.Move {
	cube := *c
	movesToG2 := cube.ToG2()
	if movesToG2 == nil {
		return nil
	}
	cmn.ApplySequence(&cube, movesToG2)
	movesToG3 := cube.ToG3AssumingG2()
	if movesToG3 == nil {
		return nil
	}
	return append(movesToG2, movesToG3...)
}

func (c *Cube) ToG4() []cmn.Move {
	cube := *c
	movesToG3 := cube.ToG3()
	if movesToG3 == nil {
		return nil
	}
	cmn.ApplySequence(&cube, movesToG3)
	movesToG4 := cube.ToG4AssumingG3()
	if movesToG4 == nil {
		return nil
	}
	return append(movesToG3, movesToG4...)
}

func (c *Cube) Blueprint() string {
	panic("unimplemented")
}

func (c *Cube) Clone() cmn.Cube {
	newCube := *c
	return &newCube
}

var G1CornerHeuristicTable = MakeG1CornerOrientationsTable()

// TODO: Improve precision
func (c *Cube) distanceToG2InG1() int {
	coDistance := int(G1CornerHeuristicTable[c.cornerOrientations])
	// coDistance := c.cornerOrientations.Distance()
	edgeDistance := c.edgePermutation.FUBDCorrectSliceDistance()
	if coDistance > edgeDistance {
		return coDistance
	} else {
		return edgeDistance
	}
}

var G2EdgeHeuristicTable = MakeG2EdgePermutationTable()
var G2CornerHeuristicTable = MakeG2CornerPermutationTable(G3CornerHeuristicTable)

func (c *Cube) distanceToG3InG2() int {
	// epDistance := c.edgePermutation.AllInCorrectSliceDistance()
	// cpDistance := c.cornerPermutation.AllInCorrectOrbitDistance()
	epDistance := G2EdgeHeuristicTable[c.edgePermutation]
	cpDistance := G2CornerHeuristicTable[c.cornerPermutation]
	if epDistance > cpDistance {
		return int(epDistance)
	} else {
		return int(cpDistance)
	}
}

var G3CornerHeuristicTable = MakeG3CornerPermutationTable()
var G3HeuristicTable = MakeG3HeuristicTable()

func (c *Cube) distanceToG4InG3() int {
	// epDistance := c.edgePermutation.Distance()
	// cpDistance := c.cornerPermutation.Distance()
	// if epDistance > cpDistance {
	// 	return epDistance
	// } else {
	// 	return cpDistance
	// }
	return int(G3HeuristicTable[*c])
}

func (c Cube) Equal(o Cube) bool {
	return (c.edgeOrientations == o.edgeOrientations &&
		c.edgePermutation == o.edgePermutation &&
		c.cornerOrientations == o.cornerOrientations &&
		c.cornerPermutation == o.cornerPermutation)
}
