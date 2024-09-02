package cepo_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/mdesoeuv/rubik/cepo"
	"github.com/mdesoeuv/rubik/common"
)

func TestCubeToG1(t *testing.T) {
	r := rand.New(rand.NewPCG(0, 0))
	maxStepCount := 0
	for move_count := 0; move_count <= 100; move_count++ {
		cube := cepo.NewCubeSolved()
		common.Shuffle(&cube, r, move_count)
		steps := cube.ToG1()
		if steps == nil {
			t.Fatalf("There should be a solution")
		}

		if len(steps) > maxStepCount {
			maxStepCount = len(steps)
		}
		if len(steps) > move_count {
			t.Fatalf("ToG1 shouldn't have more than %v moves but got %v", move_count, steps)
		}

		for _, step := range steps {
			cube.Apply(step)
		}

		if !cube.IsG1() {
			t.Fatalf("Cube should be G1 after applying the steps")
		}
	}
	fmt.Printf("Max step count taken: %v\n", maxStepCount)
}

func TestCubeToG2(t *testing.T) {
	r := rand.New(rand.NewPCG(0, 0))
	maxStepCount := 0
	for move_count := 0; move_count <= 100; move_count++ {
		fmt.Printf("Move count: %v\n", move_count)
		cube := cepo.NewCubeSolved()
		common.Shuffle(&cube, r, move_count)

		steps := cube.ToG1()
		if steps == nil {
			t.Fatalf("There should be a solution")
		}

		if len(steps) > move_count {
			t.Fatalf("ToG1 shouldn't have more than %v moves but got %v", move_count, steps)
		}

		for _, step := range steps {
			cube.Apply(step)
		}

		if !cube.IsG1() {
			t.Fatalf("Cube should be G1 after applying the steps")
		}

		steps = cube.ToG2()
		if steps == nil {
			t.Fatalf("There should be a solution")
		}

		if len(steps) > maxStepCount {
			maxStepCount = len(steps)
		}

		for _, step := range steps {
			cube.Apply(step)
		}

		if !cube.IsG2() {
			t.Fatalf("Cube should be G2 after applying the steps")
		}
	}
	fmt.Printf("Max step count taken: %v\n", maxStepCount)
}

func TestCubeToG1FromArticleShuffle(t *testing.T) {
	shuffle := common.ArticleExampleShuffleMoveList()
	cube := cepo.NewCubeSolved()

	for _, move := range shuffle {
		cube.Apply(move)
	}

	solution := cube.ToG1()

	for _, move := range solution {
		cube.Apply(move)
	}

	if !cube.IsG1() {
		t.Fatalf("Could should be G1")
	}

	fmt.Printf("Solution: %v\n", solution)
}

func BenchmarkCubeApply(b *testing.B) {
	cube := cepo.NewCubeSolved()

	b.StartTimer()
	for i := 0; i <= b.N; i += 1 {
		cube.Apply(common.AllMoves[i%len(common.AllMoves)])
	}
	b.StopTimer()
}
