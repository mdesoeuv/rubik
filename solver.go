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

func corner(corner Corner) CornerCoords {
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
		b = CubeCoord{Left, FaceCoord{2, 2}}
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

func corner_for(a, b, c Side) (Corner, error) {
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

type MemoizationEntry struct {
	a, b, c Side
	corner  Corner
}

var cornerManhattanDistanceMap = map[MemoizationEntry]int{}

func (cube *Cube) cornerManhattanDistance(id Corner) int {
	cornerCoords := corner(id)
	sideA := *cube.get(cornerCoords.a)
	sideB := *cube.get(cornerCoords.b)
	sideC := *cube.get(cornerCoords.c)

	entry := MemoizationEntry{sideA, sideB, sideC, id}
	result, stored := cornerManhattanDistanceMap[entry]
	if stored {
		return result
	}

	to_explore := []Cube{*cube}
	expectedCorner, _ := corner_for(sideA, sideB, sideC)
	toValidateCoords := corner(expectedCorner)
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

func (c *Cube) Successors() []Cube {
	result := make([]Cube, 0, len(AllMoves))
	for _, move := range AllMoves {
		newCube := *c
		newCube.apply(move)
		result = append(result, newCube)
	}
	return result
}

func heuristic(cube *Cube) int {
	sum := 0
	for corner := FirstCorner; corner <= LastCorner; corner++ {
		sum += cube.cornerManhattanDistance(corner)
	}
	// TODO: Add side_manhattan_distance
	return (sum / 4)
}

func (cube *Cube) solve() *[]Cube {
	bound := heuristic(cube)

	path := []Cube{*cube}

	for {
		t := search(&path, 0, bound)
		if t == FOUND {
			return &path
		}
		if t == math.MaxInt {
			return nil
		}
		bound = t
	}
}

const FOUND int = -1

func search(path *[]Cube, g int, bound int) int {
	cube := &(*path)[len(*path)-1]
	f := g + heuristic(cube)
	if f > bound {
		return f
	}
	if cube.isSolved() {
		return FOUND
	}
	min := math.MaxInt
	for _, move := range AllMoves {
		newCube := *cube
		newCube.apply(move)
		if !slices.Contains(*path, newCube) {
			*path = append(*path, newCube)
			t := search(path, g+1, bound)
			if t == FOUND {
				return FOUND
			}
			if t < min {
				min = t
			}
			*path = (*path)[:len(*path)-1]
		}
	}
	return min
}
