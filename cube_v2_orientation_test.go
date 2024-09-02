package main

import (
	"slices"
	"testing"
)

func TestCubeV2TopSpinCornerOrientation(t *testing.T) {
	orientations := NewCornerOrientationsSolved()

	orientations.apply(Move{
		Side:         Up,
		NumRotations: 1,
	})

	for corner := CornerIndexFirst; corner <= CornerIndexLast; corner++ {
		orientation := orientations.get(corner)
		if slices.Contains(CrownCornerUp[:], corner) {
			// Check Up side
			if orientation != CornerOrientationFrontBack {
				t.Errorf("Invalid orientation found: %v", orientation)
			}
		} else {
			// Check the other sides
			if orientation != CornerOrientationLeftRight {
				t.Errorf("Invalid orientation found: %v", orientation)
			}
		}
	}
}
