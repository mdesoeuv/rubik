package cepo

import (
	"math"
	"slices"

	cmn "github.com/mdesoeuv/rubik/common"
)

func (c *Cube) ToG1() []cmn.Move {
	bound := c.edgeOrientations.Distance()

	seen := map[Cube]struct{}{}
	seen[*c] = struct{}{}

	for {
		t, solution := searchG1(seen, c, nil, 0, bound)
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

func searchG1(
	seen map[Cube]struct{},
	cube *Cube,
	previousMove *cmn.Move,
	g int,
	bound int,
) (int, []cmn.Move) {
	f := g + cube.edgeOrientations.Distance()
	if f > bound {
		return f, nil
	}
	if cube.IsG1() {
		return 0, []cmn.Move{}
	}
	min := math.MaxInt
	for _, move := range cmn.AllMoves {
		if previousMove != nil && previousMove.IsRedudantWith(move) {
			continue
		}
		newCube := *cube
		newCube.Apply(move)
		_, wasSeen := seen[newCube]
		if !wasSeen {
			seen[newCube] = struct{}{}
			t, steps := searchG1(seen, &newCube, &move, g+1, bound)
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

func makeG1Moves() (result []cmn.Move) {
	for side := cmn.FirstSide; side <= cmn.LastSide; side++ {
		if side == cmn.Up || side == cmn.Down {
			result = append(result, cmn.Move{
				Side:     side,
				Rotation: cmn.RotationHalfTurn(),
			})
		} else {
			for _, rotation := range cmn.AllRotations {
				result = append(result, cmn.Move{
					Side:     side,
					Rotation: rotation,
				})
			}
		}
	}
	return
}

var G1Moves = makeG1Moves()

func (c *Cube) ToG2AssumingG1() []cmn.Move {
	bound := c.distanceToG2InG1()

	seen := map[Cube]struct{}{}
	seen[*c] = struct{}{}

	for {
		t, solution := searchG2(seen, c, nil, 0, bound)
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

func searchG2(
	seen map[Cube]struct{},
	cube *Cube,
	previousMove *cmn.Move,
	g int,
	bound int,
) (int, []cmn.Move) {
	f := g + cube.distanceToG2InG1()
	if f > bound {
		return f, nil
	}
	if cube.IsG2AssumingG1() {
		return 0, []cmn.Move{}
	}
	min := math.MaxInt
	for _, move := range G1Moves {
		if previousMove != nil && previousMove.IsRedudantWith(move) {
			continue
		}
		newCube := *cube
		newCube.Apply(move)
		_, wasSeen := seen[newCube]
		if !wasSeen {
			seen[newCube] = struct{}{}
			t, steps := searchG2(seen, &newCube, &move, g+1, bound)
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

func makeG2Moves() (result []cmn.Move) {
	for side := cmn.FirstSide; side <= cmn.LastSide; side++ {
		if side == cmn.Left || side == cmn.Right {
			for _, rotation := range cmn.AllRotations {
				result = append(result, cmn.Move{
					Side:     side,
					Rotation: rotation,
				})
			}
		} else {
			result = append(result, cmn.Move{
				Side:     side,
				Rotation: cmn.RotationHalfTurn(),
			})
		}
	}
	return
}

var G2Moves = makeG2Moves()

func (c *Cube) ToG3AssumingG2() []cmn.Move {
	bound := c.distanceToG3InG2()

	seen := map[Cube]struct{}{}
	seen[*c] = struct{}{}

	for {
		t, solution := searchG3(seen, c, nil, 0, bound)
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

func searchG3(
	seen map[Cube]struct{},
	cube *Cube,
	previousMove *cmn.Move,
	g int,
	bound int,
) (int, []cmn.Move) {
	f := g + cube.distanceToG3InG2()
	if f > bound {
		return f, nil
	}
	if cube.IsG3AssumingG2() {
		return 0, []cmn.Move{}
	}
	min := math.MaxInt
	for _, move := range G2Moves {
		if previousMove != nil && previousMove.IsRedudantWith(move) {
			continue
		}
		newCube := *cube
		newCube.Apply(move)
		_, wasSeen := seen[newCube]
		if !wasSeen {
			seen[newCube] = struct{}{}
			t, steps := searchG3(seen, &newCube, &move, g+1, bound)
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

func makeG3Moves() (result []cmn.Move) {
	for side := cmn.FirstSide; side <= cmn.LastSide; side++ {
		result = append(result, cmn.Move{
			Side:     side,
			Rotation: cmn.RotationHalfTurn(),
		})
	}
	return
}

var G3Moves = makeG3Moves()
var G3Heuristic = MakeG3HeuristicTable()

func (c *Cube) ToG4AssumingG3() []cmn.Move {
	bound := c.distanceToG4InG3()

	seen := map[Cube]struct{}{}
	seen[*c] = struct{}{}

	for {
		t, solution := searchG4(seen, c, nil, 0, bound)
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

func searchG4(
	seen map[Cube]struct{},
	cube *Cube,
	previousMove *cmn.Move,
	g int,
	bound int,
) (int, []cmn.Move) {
	f := g + cube.distanceToG4InG3()
	if f > bound {
		return f, nil
	}
	if cube.IsG4AssumingG3() {
		return 0, []cmn.Move{}
	}
	min := math.MaxInt
	for _, move := range G3Moves {
		if previousMove != nil && previousMove.IsRedudantWith(move) {
			continue
		}
		newCube := *cube
		newCube.Apply(move)
		_, wasSeen := seen[newCube]
		if !wasSeen {
			seen[newCube] = struct{}{}
			t, steps := searchG4(seen, &newCube, &move, g+1, bound)
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

func MakeG3Cubes() map[Cube]struct{} {
	solvedCube := NewCubeSolved()
	toExplore := []Cube{*solvedCube}
	seen := map[Cube]struct{}{*solvedCube: {}}

	for len(toExplore) != 0 {
		// TODO: a better queue
		cube := toExplore[0]
		toExplore = toExplore[1:]

		for _, move := range G3Moves {
			newCube := cube
			newCube.Apply(move)

			if _, x := seen[newCube]; x {
				continue
			}
			seen[newCube] = struct{}{}

			toExplore = append(toExplore, newCube)
		}
	}
	return seen
}

func MakeG3HeuristicTable() map[Cube]uint8 {
	solvedCube := *NewCubeSolved()
	toExplore := []Cube{solvedCube}
	toExploreNext := []Cube{}
	result := map[Cube]uint8{solvedCube: 0}

	distance := uint8(1)
	for len(toExplore) != 0 {
		for _, cube := range toExplore {
			for _, move := range G3Moves {
				newCube := cube
				newCube.Apply(move)

				if _, x := result[newCube]; x {
					continue
				}
				result[newCube] = distance
				toExploreNext = append(toExploreNext, newCube)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore[:0]
		distance += 1

	}
	return result
}

type G3Cube struct {
	edge   EdgePermutation
	corner CornerPermutation
}

func ToG3Cube(c Cube) G3Cube {
	return G3Cube{
		edge:   c.edgePermutation,
		corner: c.cornerPermutation,
	}
}

func (c *G3Cube) Apply(move cmn.Move) {
	c.edge.Apply(move)
	c.corner.Apply(move)
}

func MakeG3BetterHeuristicTable() map[G3Cube]uint8 {
	solvedCube := ToG3Cube(*NewCubeSolved())
	toExplore := []G3Cube{solvedCube}
	toExploreNext := []G3Cube{}
	result := map[G3Cube]uint8{solvedCube: 0}

	distance := uint8(1)
	for len(toExplore) != 0 {
		for _, cube := range toExplore {
			for _, move := range G3Moves {
				newCube := cube
				newCube.Apply(move)

				if _, x := result[newCube]; x {
					continue
				}
				result[newCube] = distance
				toExploreNext = append(toExploreNext, newCube)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore[:0]
		distance += 1
	}
	return result
}

func MakeG3CornerPermutationTable() map[CornerPermutation]uint8 {
	solvedCube := NewCubeSolved().cornerPermutation
	toExplore := []CornerPermutation{solvedCube}
	toExploreNext := []CornerPermutation{}
	result := map[CornerPermutation]uint8{solvedCube: 0}

	distance := uint8(1)
	for len(toExplore) != 0 {
		for _, cube := range toExplore {
			for _, move := range G3Moves {
				newCube := cube
				newCube.Apply(move)

				if _, x := result[newCube]; x {
					continue
				}
				result[newCube] = distance
				toExploreNext = append(toExploreNext, newCube)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore[:0]
		distance += 1
	}
	return result
}

func MakeG3EdgePermutationTable() map[EdgePermutation]uint8 {
	solvedCube := NewCubeSolved().edgePermutation
	toExplore := []EdgePermutation{solvedCube}
	toExploreNext := []EdgePermutation{}
	result := map[EdgePermutation]uint8{solvedCube: 0}

	distance := uint8(1)
	for len(toExplore) != 0 {
		for _, cube := range toExplore {
			for _, move := range G3Moves {
				newCube := cube
				newCube.Apply(move)

				if _, x := result[newCube]; x {
					continue
				}
				result[newCube] = distance
				toExploreNext = append(toExploreNext, newCube)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore[:0]
		distance += 1
	}
	return result
}

func MakeG2CornerPermutationTable(g3 map[CornerPermutation]uint8) map[CornerPermutation]uint8 {
	toExplore := []CornerPermutation{}
	toExploreNext := []CornerPermutation{}
	result := map[CornerPermutation]uint8{}

	for corner := range g3 {
		toExplore = append(toExplore, corner)
		result[corner] = 0
	}

	distance := uint8(1)
	for len(toExplore) != 0 {
		for _, cube := range toExplore {
			for _, move := range G2Moves {
				newCube := cube
				newCube.Apply(move)

				if _, x := result[newCube]; x {
					continue
				}
				result[newCube] = distance
				toExploreNext = append(toExploreNext, newCube)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore[:0]
		distance += 1
	}
	return result
}

func MakeG2EdgePermutationTable() map[EdgePermutation]uint8 {
	solvedCube := NewCubeSolved().edgePermutation
	toExplore := []EdgePermutation{solvedCube}
	toExploreNext := []EdgePermutation{}
	result := map[EdgePermutation]uint8{solvedCube: 0}

	distance := uint8(1)
	for len(toExplore) != 0 {
		for _, cube := range toExplore {
			for _, move := range G2Moves {
				newCube := cube
				newCube.Apply(move)

				if _, x := result[newCube]; x {
					continue
				}
				result[newCube] = distance
				toExploreNext = append(toExploreNext, newCube)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore[:0]
		distance += 1
	}
	return result
}

func MakeG1CornerOrientationsTable() map[CornerOrientations]uint8 {
	solvedCube := NewCubeSolved().cornerOrientations
	toExplore := []CornerOrientations{solvedCube}
	toExploreNext := []CornerOrientations{}
	result := map[CornerOrientations]uint8{solvedCube: 0}

	distance := uint8(1)
	for len(toExplore) != 0 {
		for _, cube := range toExplore {
			for _, move := range G1Moves {
				newCube := cube
				newCube.Apply(move)

				if _, x := result[newCube]; x {
					continue
				}
				result[newCube] = distance
				toExploreNext = append(toExploreNext, newCube)
			}
		}
		toExplore, toExploreNext = toExploreNext, toExplore[:0]
		distance += 1
	}
	return result
}

// TODO: Only store the interresting values
// Not every orientation is valuable

// func MakeG1EdgePermutationTable() map[EdgePermutation]uint8 {
// 	solvedCube := NewCubeSolved().edgePermutation
// 	toExplore := []EdgePermutation{solvedCube}
// 	toExploreNext := []EdgePermutation{}
// 	result := map[EdgePermutation]uint8{solvedCube: 0}
//
// 	distance := uint8(1)
// 	for len(toExplore) != 0 {
// 		for _, cube := range toExplore {
// 			for _, move := range G1Moves {
// 				newCube := cube
// 				newCube.Apply(move)
//
// 				if _, x := result[newCube]; x {
// 					continue
// 				}
// 				result[newCube] = distance
// 				toExploreNext = append(toExploreNext, newCube)
// 			}
// 		}
// 		toExplore, toExploreNext = toExploreNext, toExplore[:0]
// 		distance += 1
// 	}
// 	return result
// }
