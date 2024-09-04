package main

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/mdesoeuv/rubik/cepo"
	"github.com/mdesoeuv/rubik/common"
	visual "github.com/mdesoeuv/rubik/visual"
)

func TestCubeToG3(t *testing.T) {
	r := rand.New(rand.NewPCG(0, 0))
	maxStepCount := 0

	G3Cubes := cepo.MakeG3Cubes()

	for move_count := 0; move_count <= 100; move_count += 10 {
		newCepo := cepo.NewCubeSolved()
		cube := VisualCepo{Cepo: newCepo, Visual: visual.NewCubeSolved()}

		common.Shuffle(&cube, r, move_count)

		fmt.Printf("Solving #%v\n", move_count)
		steps := cube.Cepo.ToG3()
		if steps == nil {
			t.Fatalf("There should be a solution")
		}

		if len(steps) > maxStepCount {
			maxStepCount = len(steps)
		}

		common.ApplySequence(&cube, steps)

		if !cube.Cepo.IsG3() {
			t.Fatalf("Cube should be G3 after applying the steps")
		}

		if _, isG3 := G3Cubes[*cube.Cepo]; !isG3 {
			t.Fatalf("Cube.IsG3 and G3Cubes disagree:\n%v", cube.Blueprint())
		}
	}
	fmt.Printf("Max step count taken: %v\n", maxStepCount)
}

func BenchmarkSolveCube(b *testing.B) {
	r := rand.New(rand.NewPCG(0, 0))
	maxStepCount := 0
	for n := 0; n < b.N; n++ {
		for move_count := 0; move_count <= 100; move_count += 10 {
			newCepo := cepo.NewCubeSolved()
			cube := VisualCepo{Cepo: newCepo, Visual: visual.NewCubeSolved()}

			common.Shuffle(&cube, r, move_count)

			b.Logf("Solving #%v\n", move_count)
			steps := cube.Solve()
			if steps == nil {
				b.Fatalf("There should be a solution")
			}

			if len(steps) > maxStepCount {
				maxStepCount = len(steps)
			}

			common.ApplySequence(&cube, steps)

			if !cube.IsSolved() {
				b.Fatalf("Cube should be solved after applying the steps")
			}
		}
		b.Logf("Max step count taken: %v\n", maxStepCount)
	}
}
