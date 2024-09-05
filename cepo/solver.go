package cepo

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"math"
	"os"
	"slices"

	cmn "github.com/mdesoeuv/rubik/common"
)

type Solver struct {
	G1CornerHeuristicTable map[CornerOrientations]uint8
	G2EdgeHeuristicTable   map[EdgePermutation]uint8
	G2CornerHeuristicTable map[CornerPermutation]uint8
	G3EdgeHeuristicTable   map[EdgePermutation]uint8
	G3CornerHeuristicTable map[CornerPermutation]uint8
	G3HeuristicTable       map[Cube]uint8
}

func (s *Solver) save() error {
	file, createError := os.Create("rubik.cache")
	if createError != nil {
		return fmt.Errorf("could not create cache file: %v", createError)
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	error := encoder.Encode(*s)
	return error
}

func LoadSolver(cacheName string) (*Solver, error) {
	solver := &Solver{}
	file, openError := os.Open(cacheName)
	if openError != nil {
		return solver, openError
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := gob.NewDecoder(reader)
	decodeError := decoder.Decode(solver)
	if decodeError != nil {
		os.Remove(cacheName)
	}
	return solver, decodeError
}

func (s *Solver) PrintStats() {
	fmt.Printf("G1CornerHeuristicTable: %v\n", len(s.G1CornerHeuristicTable))
	fmt.Printf("G2EdgeHeuristicTable:   %v\n", len(s.G2EdgeHeuristicTable))
	fmt.Printf("G2CornerHeuristicTable: %v\n", len(s.G2CornerHeuristicTable))
	fmt.Printf("G3EdgeHeuristicTable:   %v\n", len(s.G3EdgeHeuristicTable))
	fmt.Printf("G3CornerHeuristicTable: %v\n", len(s.G3CornerHeuristicTable))
	fmt.Printf("G3HeuristicTable:       %v\n", len(s.G3HeuristicTable))
}

func NewSolver() *Solver {
	loadedSolver, loadError := LoadSolver("rubik.cache")
	if loadError == nil {
		return loadedSolver
	}
	G3CornerHeuristicTable := MakeG3CornerPermutationTable()
	solver := &Solver{
		G1CornerHeuristicTable: MakeG1CornerOrientationsTable(),
		G2EdgeHeuristicTable:   MakeG2EdgePermutationTable(),
		G2CornerHeuristicTable: MakeG2CornerPermutationTable(G3CornerHeuristicTable),
		G3EdgeHeuristicTable:   MakeG3EdgePermutationTable(),
		G3CornerHeuristicTable: G3CornerHeuristicTable,
		G3HeuristicTable:       MakeG3HeuristicTable(),
	}
	solver.save()
	return solver
}

var maybeNilSolver *Solver

func GetGlobalSolver() *Solver {
	if maybeNilSolver == nil {
		maybeNilSolver = NewSolver()
	}
	return maybeNilSolver
}

// TODO: Improve precision
func (s *Solver) distanceToG2InG1(c *Cube) int {
	coDistance := int(s.G1CornerHeuristicTable[c.CornerOrientations])
	edgeDistance := c.EdgePermutation.FUBDCorrectSliceDistance()
	if coDistance > edgeDistance {
		return coDistance
	} else {
		return edgeDistance
	}
}

func (s *Solver) distanceToG3InG2(c *Cube) int {
	epDistance := s.G2EdgeHeuristicTable[c.EdgePermutation]
	cpDistance := s.G2CornerHeuristicTable[c.CornerPermutation]
	if epDistance > cpDistance {
		return int(epDistance)
	} else {
		return int(cpDistance)
	}
}

func (s *Solver) distanceToG4InG3(c *Cube) int {
	epDistance := s.G3EdgeHeuristicTable[c.EdgePermutation]
	cpDistance := s.G3CornerHeuristicTable[c.CornerPermutation]
	if epDistance > cpDistance {
		return int(epDistance)
	} else {
		return int(cpDistance)
	}
}

func (s *Solver) ToG1(c *Cube) []cmn.Move {
	bound := c.EdgeOrientations.Distance()

	seen := map[Cube]struct{}{}
	seen[*c] = struct{}{}

	for {
		t, solution := s.searchG1(seen, c, nil, 0, bound)
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

func (s *Solver) searchG1(
	seen map[Cube]struct{},
	cube *Cube,
	previousMove *cmn.Move,
	g int,
	bound int,
) (int, []cmn.Move) {
	f := g + cube.EdgeOrientations.Distance()
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
			t, steps := s.searchG1(seen, &newCube, &move, g+1, bound)
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

func (s *Solver) ToG2AssumingG1(c *Cube) []cmn.Move {
	bound := s.distanceToG2InG1(c)

	seen := map[Cube]struct{}{}
	seen[*c] = struct{}{}

	for {
		t, solution := s.searchG2(seen, c, nil, 0, bound)
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

func (s *Solver) ToG2(c *Cube) []cmn.Move {
	cube := *c
	movesToG1 := s.ToG1(&cube)
	if movesToG1 == nil {
		return nil
	}
	cmn.ApplySequence(&cube, movesToG1)
	movesToG2 := s.ToG2AssumingG1(&cube)
	if movesToG2 == nil {
		return nil
	}
	return append(movesToG1, movesToG2...)
}

func (s *Solver) searchG2(
	seen map[Cube]struct{},
	cube *Cube,
	previousMove *cmn.Move,
	g int,
	bound int,
) (int, []cmn.Move) {
	f := g + s.distanceToG2InG1(cube)
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
			t, steps := s.searchG2(seen, &newCube, &move, g+1, bound)
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

func (s *Solver) ToG3AssumingG2(c *Cube) []cmn.Move {
	bound := s.distanceToG3InG2(c)

	seen := map[Cube]struct{}{}
	seen[*c] = struct{}{}

	for {
		t, solution := s.searchG3(seen, c, nil, 0, bound)
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

func (s *Solver) ToG3(c *Cube) []cmn.Move {
	cube := *c
	movesToG2 := s.ToG2(&cube)
	if movesToG2 == nil {
		return nil
	}
	cmn.ApplySequence(&cube, movesToG2)
	movesToG3 := s.ToG3AssumingG2(&cube)
	if movesToG3 == nil {
		return nil
	}
	return append(movesToG2, movesToG3...)
}

func (s *Solver) searchG3(
	seen map[Cube]struct{},
	cube *Cube,
	previousMove *cmn.Move,
	g int,
	bound int,
) (int, []cmn.Move) {
	f := g + s.distanceToG3InG2(cube)
	if f > bound {
		return f, nil
	}
	if s.IsG3AssumingG2(cube) {
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
			t, steps := s.searchG3(seen, &newCube, &move, g+1, bound)
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

func (s *Solver) ToG4AssumingG3(c *Cube) []cmn.Move {
	bound := s.distanceToG4InG3(c)

	seen := map[Cube]struct{}{}
	seen[*c] = struct{}{}

	for {
		t, solution := s.searchG4(seen, c, nil, 0, bound)
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

func (s *Solver) ToG4(c *Cube) []cmn.Move {
	cube := *c
	movesToG3 := s.ToG3(&cube)
	if movesToG3 == nil {
		return nil
	}
	cmn.ApplySequence(&cube, movesToG3)
	movesToG4 := s.ToG4AssumingG3(&cube)
	if movesToG4 == nil {
		return nil
	}
	return append(movesToG3, movesToG4...)
}

func (s *Solver) searchG4(
	seen map[Cube]struct{},
	cube *Cube,
	previousMove *cmn.Move,
	g int,
	bound int,
) (int, []cmn.Move) {
	f := g + s.distanceToG4InG3(cube)
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
			t, steps := s.searchG4(seen, &newCube, &move, g+1, bound)
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
		edge:   c.EdgePermutation,
		corner: c.CornerPermutation,
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
	solvedCube := NewCubeSolved().CornerPermutation
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
	solvedCube := NewCubeSolved().EdgePermutation
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
	solvedCube := NewCubeSolved().EdgePermutation
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
	solvedCube := NewCubeSolved().CornerOrientations
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
