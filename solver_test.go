package main

import "testing"

func TestManhattanDistanceSolved(t *testing.T) {
	cube := NewCubeSolved()

	for corner := FirstCorner; corner <= LastCorner; corner++ {
		distance := cube.corner_manhattan_distance(CornerDownLeftFront)
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

	if cube.corner_manhattan_distance(CornerUpLeftFront) != 1 {
		t.Fatalf("Top spinned cube should have up left front corner at distance 1")
	}
}

func TestManhattanDistanceDoubleMove(t *testing.T) {
	cube := NewCubeSolved()

	cube.apply(Move{Front, -1})
	cube.apply(Move{Down, 1})

	distance := cube.corner_manhattan_distance(CornerDownRightFront)
	if distance != 2 {
		t.Fatalf(
			"Cube with 2 moves have distance of r, "+
				"but got distance %v", distance,
		)
	}
}

func TestSolveSolved(t *testing.T) {
	cube := NewCubeSolved()

	solution := cube.solve()
	if len(*solution) != 1 {
		t.Fatalf("Solved cube should require 0 steps")
	}
}

func TestSolveSingleMove(t *testing.T) {
	cube := NewCubeSolved()

	cube.apply(Move{Front, 1})

	solution := cube.solve()
	if solution == nil {
		t.Fatalf("Should find solution")
	}
	steps := len(*solution) - 1
	if steps != 1 {
		t.Fatalf("Should take 1 step, but took %v steps", steps)
	}
}
