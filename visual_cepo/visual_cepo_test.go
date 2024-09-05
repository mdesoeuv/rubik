package visual_cepo

import (
	"math/rand/v2"
	"testing"

	"github.com/mdesoeuv/rubik/cepo"
	"github.com/mdesoeuv/rubik/common"
	visual "github.com/mdesoeuv/rubik/visual"
)

// TODO: Remove this
// It doesn't much sens anymore
func TestCubeToG3(t *testing.T) {
	r := rand.New(rand.NewPCG(0, 0))
	maxStepCount := 0

	solver := cepo.GetGlobalSolver()

	for move_count := 0; move_count <= 100; move_count++ {
		newCepo := cepo.NewCubeSolved()
		cube := Cube{Cepo: newCepo, Visual: visual.NewCubeSolved()}

		common.Shuffle(&cube, r, move_count)

		t.Logf("Solving #%v", move_count)
		steps := solver.ToG3(cube.Cepo)
		if steps == nil {
			t.Fatalf("There should be a solution")
		}

		if len(steps) > maxStepCount {
			maxStepCount = len(steps)
		}

		common.ApplySequence(&cube, steps)

		if !solver.IsG3(cube.Cepo) {
			t.Fatalf("Cube should be G3 after applying the steps")
		}

		// if _, isG3 := G3Cubes[*cube.Cepo]; !isG3 {
		// 	t.Fatalf("Cube.IsG3 and G3Cubes disagree:\n%v", cube.Blueprint())
		// }
	}
	t.Logf("Max step count taken: %v", maxStepCount)
}
