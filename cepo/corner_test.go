package cepo_test

import (
	"math/rand/v2"
	"testing"

	"github.com/mdesoeuv/rubik/cepo"
	"github.com/mdesoeuv/rubik/common"
)

func TestCornerOrientationRewind(t *testing.T) {
	orientations := cepo.NewCornerOrientationsSolved()
	r := rand.New(rand.NewPCG(0, 0))

	moveSequence := []common.Move{}

	for i := 0; i < 100; i++ {
		move := common.AllMoves[r.IntN(len(common.AllMoves))]
		orientations.Apply(move)
		moveSequence = append(moveSequence, move)

		orientations := orientations
		for j := i; j >= 0; j-- {
			orientations.Apply(moveSequence[j].Reverse())
		}

		if !orientations.IsSolved() {
			t.Fatalf("Doing %v in reverse didn't reverse the corner orientations", moveSequence)
		}
	}
}

func TestCornerPermutationRewind(t *testing.T) {
	permutation := cepo.NewCornerPermutationSolved()
	r := rand.New(rand.NewPCG(0, 0))

	moveSequence := []common.Move{}

	for i := 0; i < 100; i++ {
		move := common.AllMoves[r.IntN(len(common.AllMoves))]
		permutation.Apply(move)
		moveSequence = append(moveSequence, move)

		orientations := permutation
		for j := i; j >= 0; j-- {
			orientations.Apply(moveSequence[j].Reverse())
		}

		if !orientations.IsSolved() {
			t.Fatalf("Doing %v in reverse didn't reverse the corner permutation", moveSequence)
		}
	}
}

func BenchmarkCornerOrientationsApply(b *testing.B) {
	orientations := cepo.NewCornerOrientationsSolved()

	b.StartTimer()
	for i := 0; i <= b.N; i += 1 {
		orientations.Apply(common.AllMoves[i%len(common.AllMoves)])
	}
	b.StopTimer()
}

func BenchmarkCornerPermutationApply(b *testing.B) {
	permutation := cepo.NewCornerPermutationSolved()

	b.StartTimer()
	for i := 0; i <= b.N; i += 1 {
		permutation.Apply(common.AllMoves[i%len(common.AllMoves)])
	}
	b.StopTimer()
}
