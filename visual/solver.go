package visual

import (
	"fmt"
	"math"
	"slices"

	cmn "github.com/mdesoeuv/rubik/common"
)

type Corner = int

const (
	CornerUpLeftFront    = 0
	CornerUpRightFront   = 1
	CornerUpLeftBack     = 2
	CornerUpRightBack    = 3
	CornerDownLeftFront  = 4
	CornerDownRightFront = 5
	CornerDownLeftBack   = 6
	CornerDownRightBack  = 7
	FirstCorner          = CornerUpLeftFront
	LastCorner           = CornerDownRightBack
	CornerCount          = LastCorner + 1
)

type CornerCoords struct {
	a, b, c cmn.CubeCoord
}

type Edge = int

const (
	EdgeUpLeft     Edge = 0
	EdgeUpRight    Edge = 1
	EdgeUpFront    Edge = 2
	EdgeUpBack     Edge = 3
	EdgeDownLeft   Edge = 4
	EdgeDownRight  Edge = 5
	EdgeDownFront  Edge = 6
	EdgeDownBack   Edge = 7
	EdgeLeftFront  Edge = 8
	EdgeLeftBack   Edge = 9
	EdgeRightFront Edge = 10
	EdgeRightBack  Edge = 11
	FirstEdge      Edge = EdgeUpLeft
	LastEdge       Edge = EdgeRightBack
	EdgeCount      int  = LastEdge + 1
)

type EdgeCoords struct {
	a, b cmn.CubeCoord
}

func cornerCoords(corner Corner) CornerCoords {
	var a, b, c cmn.CubeCoord
	switch corner {
	case CornerUpLeftFront:
		a = cmn.CubeCoord{Side: cmn.Up, FaceCoord: cmn.FaceCoord20}
		b = cmn.CubeCoord{Side: cmn.Left, FaceCoord: cmn.FaceCoord02}
		c = cmn.CubeCoord{Side: cmn.Front, FaceCoord: cmn.FaceCoord00}
	case CornerUpRightFront:
		a = cmn.CubeCoord{Side: cmn.Up, FaceCoord: cmn.FaceCoord22}
		b = cmn.CubeCoord{Side: cmn.Right, FaceCoord: cmn.FaceCoord00}
		c = cmn.CubeCoord{Side: cmn.Front, FaceCoord: cmn.FaceCoord02}
	case CornerUpLeftBack:
		a = cmn.CubeCoord{Side: cmn.Up, FaceCoord: cmn.FaceCoord00}
		b = cmn.CubeCoord{Side: cmn.Left, FaceCoord: cmn.FaceCoord00}
		c = cmn.CubeCoord{Side: cmn.Back, FaceCoord: cmn.FaceCoord02}
	case CornerUpRightBack:
		a = cmn.CubeCoord{Side: cmn.Up, FaceCoord: cmn.FaceCoord02}
		b = cmn.CubeCoord{Side: cmn.Right, FaceCoord: cmn.FaceCoord02}
		c = cmn.CubeCoord{Side: cmn.Back, FaceCoord: cmn.FaceCoord00}
	case CornerDownLeftFront:
		a = cmn.CubeCoord{Side: cmn.Down, FaceCoord: cmn.FaceCoord00}
		b = cmn.CubeCoord{Side: cmn.Left, FaceCoord: cmn.FaceCoord22}
		c = cmn.CubeCoord{Side: cmn.Front, FaceCoord: cmn.FaceCoord20}
	case CornerDownRightFront:
		a = cmn.CubeCoord{Side: cmn.Down, FaceCoord: cmn.FaceCoord02}
		b = cmn.CubeCoord{Side: cmn.Right, FaceCoord: cmn.FaceCoord20}
		c = cmn.CubeCoord{Side: cmn.Front, FaceCoord: cmn.FaceCoord22}
	case CornerDownLeftBack:
		a = cmn.CubeCoord{Side: cmn.Down, FaceCoord: cmn.FaceCoord20}
		b = cmn.CubeCoord{Side: cmn.Left, FaceCoord: cmn.FaceCoord20}
		c = cmn.CubeCoord{Side: cmn.Back, FaceCoord: cmn.FaceCoord22}
	case CornerDownRightBack:
		a = cmn.CubeCoord{Side: cmn.Down, FaceCoord: cmn.FaceCoord22}
		b = cmn.CubeCoord{Side: cmn.Right, FaceCoord: cmn.FaceCoord22}
		c = cmn.CubeCoord{Side: cmn.Back, FaceCoord: cmn.FaceCoord20}
	}
	return CornerCoords{a, b, c}
}

func edgeCoords(e Edge) EdgeCoords {
	var a, b cmn.CubeCoord
	switch e {
	case EdgeUpLeft:
		a = cmn.CubeCoord{Side: cmn.Up, FaceCoord: cmn.FaceCoord10}
		b = cmn.CubeCoord{Side: cmn.Left, FaceCoord: cmn.FaceCoord01}
	case EdgeUpRight:
		a = cmn.CubeCoord{Side: cmn.Up, FaceCoord: cmn.FaceCoord12}
		b = cmn.CubeCoord{Side: cmn.Right, FaceCoord: cmn.FaceCoord01}
	case EdgeUpFront:
		a = cmn.CubeCoord{Side: cmn.Up, FaceCoord: cmn.FaceCoord21}
		b = cmn.CubeCoord{Side: cmn.Front, FaceCoord: cmn.FaceCoord01}
	case EdgeUpBack:
		a = cmn.CubeCoord{Side: cmn.Up, FaceCoord: cmn.FaceCoord01}
		b = cmn.CubeCoord{Side: cmn.Back, FaceCoord: cmn.FaceCoord01}
	case EdgeDownLeft:
		a = cmn.CubeCoord{Side: cmn.Down, FaceCoord: cmn.FaceCoord10}
		b = cmn.CubeCoord{Side: cmn.Left, FaceCoord: cmn.FaceCoord21}
	case EdgeDownRight:
		a = cmn.CubeCoord{Side: cmn.Down, FaceCoord: cmn.FaceCoord12}
		b = cmn.CubeCoord{Side: cmn.Right, FaceCoord: cmn.FaceCoord21}
	case EdgeDownFront:
		a = cmn.CubeCoord{Side: cmn.Down, FaceCoord: cmn.FaceCoord01}
		b = cmn.CubeCoord{Side: cmn.Front, FaceCoord: cmn.FaceCoord21}
	case EdgeDownBack:
		a = cmn.CubeCoord{Side: cmn.Down, FaceCoord: cmn.FaceCoord21}
		b = cmn.CubeCoord{Side: cmn.Back, FaceCoord: cmn.FaceCoord21}
	case EdgeLeftFront:
		a = cmn.CubeCoord{Side: cmn.Left, FaceCoord: cmn.FaceCoord12}
		b = cmn.CubeCoord{Side: cmn.Front, FaceCoord: cmn.FaceCoord10}
	case EdgeLeftBack:
		a = cmn.CubeCoord{Side: cmn.Left, FaceCoord: cmn.FaceCoord10}
		b = cmn.CubeCoord{Side: cmn.Back, FaceCoord: cmn.FaceCoord12}
	case EdgeRightFront:
		a = cmn.CubeCoord{Side: cmn.Right, FaceCoord: cmn.FaceCoord10}
		b = cmn.CubeCoord{Side: cmn.Front, FaceCoord: cmn.FaceCoord12}
	case EdgeRightBack:
		a = cmn.CubeCoord{Side: cmn.Right, FaceCoord: cmn.FaceCoord12}
		b = cmn.CubeCoord{Side: cmn.Back, FaceCoord: cmn.FaceCoord10}
	}
	return EdgeCoords{a, b}
}

// TODO: Create CornerPiece type
func cornerFor(a, b, c cmn.Side) (Corner, error) {
	sides := []cmn.Side{a, b, c}
	slices.Sort(sides)
	switch sides[0] {
	case cmn.Up:
		switch sides[1] {
		case cmn.Left:
			switch sides[2] {
			case cmn.Back:
				return CornerUpLeftBack, nil
			case cmn.Front:
				return CornerUpLeftFront, nil
			}
		case cmn.Right:
			switch sides[2] {
			case cmn.Back:
				return CornerUpRightBack, nil
			case cmn.Front:
				return CornerUpRightFront, nil
			}
		}
	case cmn.Down:
		switch sides[1] {
		case cmn.Left:
			switch sides[2] {
			case cmn.Back:
				return CornerDownLeftBack, nil
			case cmn.Front:
				return CornerDownLeftFront, nil
			}
		case cmn.Right:
			switch sides[2] {
			case cmn.Back:
				return CornerDownRightBack, nil
			case cmn.Front:
				return CornerDownRightFront, nil
			}
		}
	}
	return 0, fmt.Errorf("Impossible side combination")
}

// TODO: Create EdgePiece type
func edgeFor(a, b cmn.Side) (Edge, error) {
	if b < a {
		a, b = b, a
	}
	switch a {
	case cmn.Up:
		switch b {
		case cmn.Left:
			return EdgeUpLeft, nil
		case cmn.Right:
			return EdgeUpRight, nil
		case cmn.Front:
			return EdgeUpFront, nil
		case cmn.Back:
			return EdgeUpBack, nil
		}
	case cmn.Down:
		switch b {
		case cmn.Left:
			return EdgeDownLeft, nil
		case cmn.Right:
			return EdgeDownRight, nil
		case cmn.Front:
			return EdgeDownFront, nil
		case cmn.Back:
			return EdgeDownBack, nil
		}
	case cmn.Left:
		switch b {
		case cmn.Front:
			return EdgeLeftFront, nil
		case cmn.Back:
			return EdgeLeftBack, nil
		}
	case cmn.Right:
		switch b {
		case cmn.Front:
			return EdgeRightFront, nil
		case cmn.Back:
			return EdgeRightBack, nil
		}
	}
	return -1, fmt.Errorf("Impossible side combination")
}

type CornerMemoizationEntry struct {
	a, b, c cmn.Side
	corner  Corner
}

type EdgeMemoizationEntry struct {
	a, b cmn.Side
	edge Edge
}

// var cornerManhattanDistanceMap = map[CornerMemoizationEntry]int{}
// var edgeManhattanDistanceMap = map[EdgeMemoizationEntry]int{}
var cornerManhattanDistanceMap = makeCornerManhattanDistanceMap()
var edgeManhattanDistanceMap = makeEdgeManhattanDistanceMap()

func makeCornerManhattanDistanceMap() map[CornerMemoizationEntry]int {
	result := map[CornerMemoizationEntry]int{}

	expectedSize := 8 * 8 * 3

	solvedCube := NewCubeSolved()

	seenCubes := map[Cube]struct{}{
		*solvedCube: {},
	}
	toExplore := []Cube{*solvedCube}
	toExploreNext := []Cube{}
	for distance := 0; len(result) < expectedSize; distance++ {
		for _, cube := range toExplore {
			for c := FirstCorner; c <= LastCorner; c++ {
				cc := cornerCoords(c)
				entry := CornerMemoizationEntry{
					a:      *cube.Get(cc.a),
					b:      *cube.Get(cc.b),
					c:      *cube.Get(cc.c),
					corner: c,
				}
				_, configurationSeen := result[entry]
				if !configurationSeen {
					result[entry] = distance
				}
			}
			for _, next := range cube.Successors() {
				_, seen := seenCubes[next]
				if seen {
					continue
				}
				seenCubes[next] = struct{}{}
				toExploreNext = append(toExploreNext, next)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore
		// Reuse allocated storage
		toExploreNext = toExploreNext[:0]
	}

	return result
}

func (cube *Cube) cornerManhattanDistance(id Corner) int {
	coords := cornerCoords(id)
	sideA := *cube.Get(coords.a)
	sideB := *cube.Get(coords.b)
	sideC := *cube.Get(coords.c)

	entry := CornerMemoizationEntry{sideA, sideB, sideC, id}
	result, stored := cornerManhattanDistanceMap[entry]
	if stored {
		return result
	}

	panic("Unreachable")

	to_explore := []Cube{*cube}
	expectedCorner, _ := cornerFor(sideA, sideB, sideC)
	toValidateCoords := cornerCoords(expectedCorner)
	for move_count := 0; move_count < 10; move_count++ {
		to_explore_next := []Cube{}
		for _, c := range to_explore {
			aIsValid := *c.Get(toValidateCoords.a) == toValidateCoords.a.Side
			bIsValid := *c.Get(toValidateCoords.b) == toValidateCoords.b.Side
			cIsValid := *c.Get(toValidateCoords.c) == toValidateCoords.c.Side
			if aIsValid && bIsValid && cIsValid {
				cornerManhattanDistanceMap[entry] = move_count
				return move_count
			}
			to_explore_next = append(to_explore_next, c.Successors()...)
		}
		to_explore = to_explore_next
	}
	panic("Could not find distance")
}

func makeEdgeManhattanDistanceMap() map[EdgeMemoizationEntry]int {
	result := map[EdgeMemoizationEntry]int{}

	expectedSize := 12 * 12 * 2

	solvedCube := NewCubeSolved()

	seenCubes := map[Cube]struct{}{
		*solvedCube: {},
	}
	toExplore := []Cube{*solvedCube}
	toExploreNext := []Cube{}
	for distance := 0; len(result) < expectedSize; distance++ {
		for _, cube := range toExplore {
			for e := FirstEdge; e <= LastEdge; e++ {
				ec := edgeCoords(e)
				entry := EdgeMemoizationEntry{
					a:    *cube.Get(ec.a),
					b:    *cube.Get(ec.b),
					edge: e,
				}
				_, configurationSeen := result[entry]
				if !configurationSeen {
					result[entry] = distance
				}
			}
			for _, next := range cube.Successors() {
				_, seen := seenCubes[next]
				if seen {
					continue
				}
				seenCubes[next] = struct{}{}
				toExploreNext = append(toExploreNext, next)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore
		// Reuse allocated storage
		toExploreNext = toExploreNext[:0]
	}

	return result
}

func (cube *Cube) edgeManhattanDistance(id Edge) int {
	coords := edgeCoords(id)
	sideA := *cube.Get(coords.a)
	sideB := *cube.Get(coords.b)

	entry := EdgeMemoizationEntry{sideA, sideB, id}
	result, stored := edgeManhattanDistanceMap[entry]
	if stored {
		return result
	}

	panic("Unreachable")

	to_explore := []Cube{*cube}
	expectedEdge, _ := edgeFor(sideA, sideB)
	toValidateCoords := edgeCoords(expectedEdge)
	for move_count := 0; move_count < 10; move_count++ {
		to_explore_next := []Cube{}
		for _, c := range to_explore {
			aIsValid := *c.Get(toValidateCoords.a) == toValidateCoords.a.Side
			bIsValid := *c.Get(toValidateCoords.b) == toValidateCoords.b.Side
			if aIsValid && bIsValid {
				edgeManhattanDistanceMap[entry] = move_count
				return move_count
			}
			to_explore_next = append(to_explore_next, c.Successors()...)
		}
		to_explore = to_explore_next
	}
	panic("Could not find distance")
}

func (c *Cube) Successors() []Cube {
	result := make([]Cube, 0, len(cmn.AllMoves))
	for _, move := range cmn.AllMoves {
		newCube := *c
		newCube.Apply(move)
		result = append(result, newCube)
	}
	return result
}

func (cube *Cube) edgeDistanceSum() int {
	sum := 0
	for edge := FirstEdge; edge <= LastEdge; edge++ {
		sum += cube.edgeManhattanDistance(edge)
	}
	return sum
}

func (cube *Cube) cornerDistanceSum() int {
	sum := 0
	for corner := FirstCorner; corner <= LastCorner; corner++ {
		sum += cube.cornerManhattanDistance(corner)
	}
	return sum
}

func heuristic(cube *Cube) int {
	edgeDistanceSum := cube.edgeDistanceSum()
	cornerDistanceSum := cube.cornerDistanceSum()
	if edgeDistanceSum > cornerDistanceSum {
		return (edgeDistanceSum + 3) / 4
	} else {
		return (cornerDistanceSum + 3) / 4
	}
}

func (cube *Cube) Solve() []cmn.Move {
	bound := heuristic(cube)

	seen := map[Cube]struct{}{}
	seen[*cube] = struct{}{}

	for {
		t, solution := search(seen, *cube, nil, 0, bound)
		if solution != nil {
			slices.Reverse(solution)
			return solution
		}
		if t == math.MaxInt {
			return nil
		}
		bound = t
	}
}

func (c *Cube) goodCorners() bool {
	for side := cmn.FirstSide; side <= cmn.LastSide; side++ {
		f := c.faces[side]
		good := f.f[0][0] == side && f.f[0][2] == side && f.f[2][0] == side && f.f[2][2] == side
		if !good {
			return false
		}
	}
	return true
}

func (c *Cube) goodEdges() bool {
	for side := cmn.FirstSide; side <= cmn.LastSide; side++ {
		f := c.faces[side]
		good := f.f[0][1] == side && f.f[1][0] == side && f.f[1][2] == side && f.f[2][1] == side
		if !good {
			return false
		}
	}
	return true
}

func search(seen map[Cube]struct{}, cube Cube, previousMove *cmn.Move, g int, bound int) (int, []cmn.Move) {
	f := g + heuristic(&cube)
	if f > bound {
		return f, nil
	}
	if cube.IsSolved() {
		return 0, []cmn.Move{}
	}
	min := math.MaxInt
	for _, move := range cmn.AllMoves {
		if previousMove != nil {
			if previousMove.Side == move.Side {
				continue
			}
			// Enforce opperation order for independant operations
			if previousMove.Side == cmn.Right && move.Side == cmn.Left {
				continue
			}
			if previousMove.Side == cmn.Up && move.Side == cmn.Down {
				continue
			}
			if previousMove.Side == cmn.Front && move.Side == cmn.Back {
				continue
			}
		}
		newCube := cube
		newCube.Apply(move)
		_, wasSeen := seen[newCube]
		if !wasSeen {
			seen[newCube] = struct{}{}
			t, steps := search(seen, newCube, &move, g+1, bound)
			if steps != nil {
				return t, append(steps, move)
			}
			if t < min {
				min = t
			}
			delete(seen, newCube)
		}
	}
	return min, nil
}
