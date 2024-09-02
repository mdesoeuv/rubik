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

func (c *Cube) ToG2() []cmn.Move {
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

var G1Moves = makeG1Moves()

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
