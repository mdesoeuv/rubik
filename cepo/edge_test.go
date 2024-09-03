package cepo_test

import (
	"math/rand/v2"
	"testing"

	"github.com/mdesoeuv/rubik/cepo"
	"github.com/mdesoeuv/rubik/common"
)

func TestEdgeOrientationRewind(t *testing.T) {
	orientations := cepo.NewEdgeOrientationsSolved()
	r := rand.New(rand.NewPCG(0, 0))

	moveSequence := []common.Move{}
	orientationsSequence := []cepo.EdgeOrientations{}

	for i := 0; i < 100; i++ {
		move := common.AllMoves[r.IntN(len(common.AllMoves))]
		orientationsSequence = append(orientationsSequence, orientations)
		orientations.Apply(move)
		moveSequence = append(moveSequence, move)

		o := orientations
		for j := i; j >= 0; j-- {
			move := moveSequence[j]
			o.Apply(move.Reverse())
			if !orientationsSequence[j].Equal(o) {
				t.Fatalf("%v: Doing %v in reverse didn't get the right steps back", i, move)
			}
		}

		if !o.IsSolved() {
			t.Fatalf("%v: Doing %v in reverse didn't reverse the edge orientations", i, moveSequence)
		}
	}
}

func TestEdgePermutationRewind(t *testing.T) {
	permutation := cepo.NewEdgePermutationSolved()
	r := rand.New(rand.NewPCG(0, 0))

	moveSequence := []common.Move{}
	permutationSequence := []cepo.EdgePermutation{}

	for i := 0; i < 100; i++ {
		move := common.AllMoves[r.IntN(len(common.AllMoves))]
		permutationSequence = append(permutationSequence, permutation)
		permutation.Apply(move)
		moveSequence = append(moveSequence, move)

		p := permutation
		for j := i; j >= 0; j-- {
			p.Apply(moveSequence[j].Reverse())
			if !permutationSequence[j].Equal(p) {
				t.Fatalf("%v: Doing %v in reverse didn't get the right steps back", i, move)
			}
		}

		if !p.IsSolved() {
			t.Fatalf("%v: Doing %v in reverse didn't reverse the edge permutation", i, moveSequence)
		}
	}
}

func TestEdgeOrientationDistance(t *testing.T) {
	eo := cepo.NewEdgeOrientationsSolved()

	upTurn := common.Move{
		Side:     common.Up,
		Rotation: common.RotationClockwise(),
	}
	rightTurn := common.Move{
		Side:     common.Right,
		Rotation: common.RotationClockwise(),
	}

	eo.Apply(upTurn)

	if eo.Distance() != 1 {
		t.Errorf("Up quater turned solved cube should have a distance of 1")
	}

	eo.Apply(upTurn)

	if eo.Distance() != 0 {
		t.Errorf("Up half turned solved cube should have a distance of 0")
	}

	eo.Apply(upTurn)
	eo.Apply(rightTurn)

	if eo.Distance() != 2 {
		t.Errorf("Up anti-clockwise rotation + Right rotation should be distance of 2")
	}
}

func TestEdgePermutationCorrectSlice(t *testing.T) {
	ep := cepo.NewEdgePermutationSolved()

	if !ep.URBLInCorrectSlice() {
		t.Errorf("Solved edge permutation should have URBL in correct slice")
	}

	if !ep.FRBLInCorrectSlice() {
		t.Errorf("Solved edge permutation should have FRBL in correct slice")
	}

	if !ep.FUBDInCorrectSlice() {
		t.Errorf("Solved edge permutation should have FUBD in correct slice")
	}
}

func BenchmarkEdgeOrientationsApply(b *testing.B) {
	orientations := cepo.NewEdgeOrientationsSolved()

	b.StartTimer()
	for i := 0; i <= b.N; i += 1 {
		orientations.Apply(common.AllMoves[i%len(common.AllMoves)])
	}
	b.StopTimer()
}

func BenchmarkEdgePermutationApply(b *testing.B) {
	permutation := cepo.NewEdgePermutationSolved()

	b.StartTimer()
	for i := 0; i <= b.N; i += 1 {
		permutation.Apply(common.AllMoves[i%len(common.AllMoves)])
	}
	b.StopTimer()
}
