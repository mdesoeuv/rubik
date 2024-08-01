package main

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func TestManhattanDistanceSolved(t *testing.T) {
	cube := NewCubeSolved()

	for corner := FirstCorner; corner <= LastCorner; corner++ {
		distance := cube.cornerManhattanDistance(CornerDownLeftFront)
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

	if cube.cornerManhattanDistance(CornerUpLeftFront) != 1 {
		t.Fatalf("Top spinned cube should have up left front corner at distance 1")
	}
}

func TestManhattanDistanceDoubleMove(t *testing.T) {
	cube := NewCubeSolved()

	cube.apply(Move{Front, -1})
	cube.apply(Move{Down, 1})

	distance := cube.cornerManhattanDistance(CornerDownRightFront)
	if distance != 2 {
		t.Fatalf(
			"Cube with 2 moves have distance of r, "+
				"but got distance %v", distance,
		)
	}
}

func TestSolveSolved(t *testing.T) {
	cube := NewCubeSolved()

	steps := cube.solve()
	if steps == nil || len(steps) != 0 {
		t.Fatalf("Solved cube should require 0 steps")
	}
}

func TestSolveSingleMove(t *testing.T) {
	cube := NewCubeSolved()

	cube.apply(Move{Front, 1})

	steps := cube.solve()
	if steps == nil {
		t.Fatalf("Should find solution")
	}
	if len(steps) != 1 {
		t.Fatalf("Should take 1 step, but took %v steps", len(steps))
	}
}

func TestShuffledCube(t *testing.T) {
	r := rand.New(rand.NewPCG(0, 0))
	for move_count := 0; move_count <= 5; move_count++ {
		cube := NewCubeSolved()
		cube.Shuffle(r, move_count)
		steps := cube.solve()
		fmt.Printf("Cube %v solved!\n", move_count)
		if steps == nil {
			t.Fatalf("There should be a solution")
		}

		fmt.Printf("Took %v steps\n", len(steps))
		if len(steps) > move_count {
			t.Fatalf("Solution shouldn't have more than %v moves but got %v", move_count, steps)
		}
	}
}
