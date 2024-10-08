package visual

import (
	"math/rand/v2"
	"testing"

	cmn "github.com/mdesoeuv/rubik/common"
)

var solver = NewSolver()

func TestManhattanDistanceSolved(t *testing.T) {
	cube := NewCubeSolved()

	for corner := FirstCorner; corner <= LastCorner; corner++ {
		distance := solver.cornerManhattanDistance(CornerDownLeftFront, cube)
		if distance != 0 {
			t.Fatalf(
				"All corners in solved cube should be at distance 0, "+
					"but got distance %v", distance,
			)
		}
	}
}

func TestManhattanDistanceBasic(t *testing.T) {
	cube := NewCubeTopSpinned()

	if solver.cornerManhattanDistance(CornerUpLeftFront, cube) != 1 {
		t.Fatalf("Top spinned cube should have up left front corner at distance 1")
	}
}

func TestManhattanDistanceDoubleMove(t *testing.T) {
	cube := NewCubeSolved()

	cube.Apply(cmn.Move{Side: cmn.Front, Rotation: cmn.RotationAntiClockwise()})
	cube.Apply(cmn.Move{Side: cmn.Down, Rotation: cmn.RotationClockwise()})

	distance := solver.cornerManhattanDistance(CornerDownRightFront, cube)
	if distance != 2 {
		t.Fatalf(
			"Cube with 2 moves have distance of r, "+
				"but got distance %v", distance,
		)
	}
}

func TestSolveSolved(t *testing.T) {
	cube := NewCubeSolved()

	steps := solver.Solve(cube)
	if steps == nil || len(steps) != 0 {
		t.Fatalf("Solved cube should require 0 steps")
	}
}

func TestSolveSingleMove(t *testing.T) {
	cube := NewCubeSolved()

	cube.Apply(cmn.Move{Side: cmn.Front, Rotation: cmn.RotationClockwise()})

	steps := solver.Solve(cube)
	if steps == nil {
		t.Fatalf("Should find solution")
	}
	if len(steps) != 1 {
		t.Fatalf("Should take 1 step, but took %v steps", len(steps))
	}
}

func TestShuffledCube(t *testing.T) {
	r := rand.New(rand.NewPCG(0, 0))
	for move_count := 0; move_count <= 10; move_count++ {
		cube := NewCubeSolved()
		cube.Shuffle(r, move_count)
		steps := solver.Solve(cube)
		t.Logf("Cube %v solved!\n", move_count)
		if steps == nil {
			t.Fatalf("There should be a solution")
		}

		t.Logf("Took %v steps", len(steps))
		if len(steps) > move_count {
			t.Fatalf("Solution shouldn't have more than %v moves but got %v", move_count, steps)
		}
	}
}
