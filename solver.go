package main

import (
	"fmt"
	"math"
	"slices"
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
	a, b, c CubeCoord
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
	a, b CubeCoord
}

func cornerCoords(corner Corner) CornerCoords {
	var a, b, c CubeCoord
	switch corner {
	case CornerUpLeftFront:
		a = CubeCoord{Up, FaceCoord{2, 0}}
		b = CubeCoord{Left, FaceCoord{0, 2}}
		c = CubeCoord{Front, FaceCoord{0, 0}}
	case CornerUpRightFront:
		a = CubeCoord{Up, FaceCoord{2, 2}}
		b = CubeCoord{Right, FaceCoord{0, 0}}
		c = CubeCoord{Front, FaceCoord{0, 2}}
	case CornerUpLeftBack:
		a = CubeCoord{Up, FaceCoord{0, 0}}
		b = CubeCoord{Left, FaceCoord{0, 0}}
		c = CubeCoord{Back, FaceCoord{0, 2}}
	case CornerUpRightBack:
		a = CubeCoord{Up, FaceCoord{0, 2}}
		b = CubeCoord{Right, FaceCoord{0, 2}}
		c = CubeCoord{Back, FaceCoord{0, 0}}
	case CornerDownLeftFront:
		a = CubeCoord{Down, FaceCoord{0, 0}}
		b = CubeCoord{Left, FaceCoord{2, 2}}
		c = CubeCoord{Front, FaceCoord{2, 0}}
	case CornerDownRightFront:
		a = CubeCoord{Down, FaceCoord{0, 2}}
		b = CubeCoord{Right, FaceCoord{2, 0}}
		c = CubeCoord{Front, FaceCoord{2, 2}}
	case CornerDownLeftBack:
		a = CubeCoord{Down, FaceCoord{2, 0}}
		b = CubeCoord{Left, FaceCoord{2, 0}}
		c = CubeCoord{Back, FaceCoord{2, 2}}
	case CornerDownRightBack:
		a = CubeCoord{Down, FaceCoord{2, 2}}
		b = CubeCoord{Right, FaceCoord{2, 2}}
		c = CubeCoord{Back, FaceCoord{2, 0}}
	}
	return CornerCoords{a, b, c}
}

func edgeCoords(e Edge) EdgeCoords {
	var a, b CubeCoord
	switch e {
	case EdgeUpLeft:
		a = CubeCoord{Up, FaceCoord{1, 0}}
		b = CubeCoord{Left, FaceCoord{0, 1}}
	case EdgeUpRight:
		a = CubeCoord{Up, FaceCoord{1, 2}}
		b = CubeCoord{Right, FaceCoord{0, 1}}
	case EdgeUpFront:
		a = CubeCoord{Up, FaceCoord{2, 1}}
		b = CubeCoord{Front, FaceCoord{0, 1}}
	case EdgeUpBack:
		a = CubeCoord{Up, FaceCoord{0, 1}}
		b = CubeCoord{Back, FaceCoord{0, 1}}
	case EdgeDownLeft:
		a = CubeCoord{Down, FaceCoord{1, 0}}
		b = CubeCoord{Left, FaceCoord{2, 1}}
	case EdgeDownRight:
		a = CubeCoord{Down, FaceCoord{1, 2}}
		b = CubeCoord{Right, FaceCoord{2, 1}}
	case EdgeDownFront:
		a = CubeCoord{Down, FaceCoord{0, 1}}
		b = CubeCoord{Front, FaceCoord{2, 1}}
	case EdgeDownBack:
		a = CubeCoord{Down, FaceCoord{2, 1}}
		b = CubeCoord{Back, FaceCoord{2, 1}}
	case EdgeLeftFront:
		a = CubeCoord{Left, FaceCoord{1, 2}}
		b = CubeCoord{Front, FaceCoord{1, 0}}
	case EdgeLeftBack:
		a = CubeCoord{Left, FaceCoord{1, 0}}
		b = CubeCoord{Back, FaceCoord{1, 2}}
	case EdgeRightFront:
		a = CubeCoord{Right, FaceCoord{1, 0}}
		b = CubeCoord{Front, FaceCoord{1, 2}}
	case EdgeRightBack:
		a = CubeCoord{Right, FaceCoord{1, 2}}
		b = CubeCoord{Back, FaceCoord{1, 0}}
	}
	return EdgeCoords{a, b}
}

// TODO: Create CornerPiece type
func cornerFor(a, b, c Side) (Corner, error) {
	sides := []Side{a, b, c}
	slices.Sort(sides)
	switch sides[0] {
	case Up:
		switch sides[1] {
		case Left:
			switch sides[2] {
			case Back:
				return CornerUpLeftBack, nil
			case Front:
				return CornerUpLeftFront, nil
			}
		case Right:
			switch sides[2] {
			case Back:
				return CornerUpRightBack, nil
			case Front:
				return CornerUpRightFront, nil
			}
		}
	case Down:
		switch sides[1] {
		case Left:
			switch sides[2] {
			case Back:
				return CornerDownLeftBack, nil
			case Front:
				return CornerDownLeftFront, nil
			}
		case Right:
			switch sides[2] {
			case Back:
				return CornerDownRightBack, nil
			case Front:
				return CornerDownRightFront, nil
			}
		}
	}
	return 0, fmt.Errorf("Impossible side combination")
}

// TODO: Create EdgePiece type
func edgeFor(a, b Side) (Edge, error) {
	if b < a {
		a, b = b, a
	}
	switch a {
	case Up:
		switch b {
		case Left:
			return EdgeUpLeft, nil
		case Right:
			return EdgeUpRight, nil
		case Front:
			return EdgeUpFront, nil
		case Back:
			return EdgeUpBack, nil
		}
	case Down:
		switch b {
		case Left:
			return EdgeDownLeft, nil
		case Right:
			return EdgeDownRight, nil
		case Front:
			return EdgeDownFront, nil
		case Back:
			return EdgeDownBack, nil
		}
	case Left:
		switch b {
		case Front:
			return EdgeLeftFront, nil
		case Back:
			return EdgeLeftBack, nil
		}
	case Right:
		switch b {
		case Front:
			return EdgeRightFront, nil
		case Back:
			return EdgeRightBack, nil
		}
	}
	return -1, fmt.Errorf("Impossible side combination")
}

type CornerMemoizationEntry struct {
	a, b, c Side
	corner  Corner
}

type EdgeMemoizationEntry struct {
	a, b Side
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
					a:      *cube.get(cc.a),
					b:      *cube.get(cc.b),
					c:      *cube.get(cc.c),
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
	sideA := *cube.get(coords.a)
	sideB := *cube.get(coords.b)
	sideC := *cube.get(coords.c)

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
			aIsValid := *c.get(toValidateCoords.a) == toValidateCoords.a.side
			bIsValid := *c.get(toValidateCoords.b) == toValidateCoords.b.side
			cIsValid := *c.get(toValidateCoords.c) == toValidateCoords.c.side
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
					a:    *cube.get(ec.a),
					b:    *cube.get(ec.b),
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
	sideA := *cube.get(coords.a)
	sideB := *cube.get(coords.b)

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
			aIsValid := *c.get(toValidateCoords.a) == toValidateCoords.a.side
			bIsValid := *c.get(toValidateCoords.b) == toValidateCoords.b.side
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
	result := make([]Cube, 0, len(AllMoves))
	for _, move := range AllMoves {
		newCube := *c
		newCube.apply(move)
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

func (cube *Cube) solve() []Move {
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
	for side := FirstSide; side <= LastSide; side++ {
		f := c.faces[side]
		good := f.f[0][0] == side && f.f[0][2] == side && f.f[2][0] == side && f.f[2][2] == side
		if !good {
			return false
		}
	}
	return true
}

func (c *Cube) goodEdges() bool {
	for side := FirstSide; side <= LastSide; side++ {
		f := c.faces[side]
		good := f.f[0][1] == side && f.f[1][0] == side && f.f[1][2] == side && f.f[2][1] == side
		if !good {
			return false
		}
	}
	return true
}

func search(seen map[Cube]struct{}, cube Cube, previousMove *Move, g int, bound int) (int, []Move) {
	f := g + heuristic(&cube)
	if f > bound {
		return f, nil
	}
	if cube.isSolved() {
		return 0, []Move{}
	}
	min := math.MaxInt
	for _, move := range AllMoves {
		if previousMove != nil {
			if previousMove.Side == move.Side {
				continue
			}
			// Enforce opperation order for independant operations
			if previousMove.Side == Right && move.Side == Left {
				continue
			}
			if previousMove.Side == Up && move.Side == Down {
				continue
			}
			if previousMove.Side == Front && move.Side == Back {
				continue
			}
		}
		newCube := cube
		newCube.apply(move)
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
