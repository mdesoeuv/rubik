package main

import (
	"math/rand/v2"
	"slices"
	"testing"
)

var edgeCoordList = []FaceCoord{
	{0, 1}, {1, 0}, {1, 2}, {2, 1},
}

func TestCubeV2Get(t *testing.T) {
	cube := NewCubeV2Solved()

	// Test corners
	for side := FirstSide; side <= LastSide; side++ {
		for _, coord := range edgeCoordList {
			squareSide := cube.get(CubeCoord{
				side:      side,
				faceCoord: coord,
			})
			if squareSide != side {
				t.Errorf("On side %c coord %v found %c",
					SideToString(side), coord, SideToString(squareSide),
				)
			}
		}
	}
}

func TestEdgePermutationSolved(t *testing.T) {
	ep := NewEdgePermutationSolved()

	for index := FirstEdgeIndex; index <= LastEdgeIndex; index++ {
		found := ep.get(index)
		if found != index {
			t.Errorf("Edge %v is not at its place found %v instead", index, found)
		}
	}
}

func TestCornerPermutationSolved(t *testing.T) {
	cp := NewCornerPermutationSolved()

	for index := CornerIndexFirst; index <= CornerIndexLast; index++ {
		found := cp.get(index)
		if found != index {
			t.Errorf("Corner %v is not at its place found %v instead", index, found)
		}
	}
}

func TestCubeV2TopSpinEdgeOrientation(t *testing.T) {
	cube := NewCubeV2Solved()

	cube.apply(Move{
		Side:         Up,
		NumRotations: 1,
	})

	for index := FirstEdgeIndex; index <= LastEdgeIndex; index++ {
		orientation := cube.edgeOrientations.get(index)
		if slices.Contains(CrownEdgeUp[:], index) {
			// Check Up side
			if !orientation {
				t.Errorf("Edge Orientation should have changed")
			}
		} else {
			// Check the other sides
			if orientation {
				t.Errorf("Edge Orientation should not have changed")
			}
		}
	}
}

func TestCubeV2TopSpinEdgePermutation(t *testing.T) {
	cube := NewCubeV2Solved()

	cube.apply(Move{
		Side:         Up,
		NumRotations: 1,
	})

	for index := FirstEdgeIndex; index <= LastEdgeIndex; index++ {
		originalIndex := cube.edgePermutation.get(index)
		edgeArrayIndex := slices.Index(CrownEdgeUp[:], index)
		if edgeArrayIndex != -1 {
			// Check Up side
			expected := CrownEdgeUp[(edgeArrayIndex+3)%len(CrownEdgeUp)]
			if originalIndex != expected {
				t.Errorf("Expected edge %v but found %v", expected, originalIndex)
			}
		} else {
			// Check the other sides
			if originalIndex != index {
				t.Errorf("Edge Orientation should not have changed")
			}
		}
	}
}

//
// func TestCubeV2TopSpin(t *testing.T) {
// 	cube := NewCubeV2Solved()
//
// 	cube.apply(Move{
// 		Side:         Up,
// 		NumRotations: 1,
// 	})
//
// 	// Check Up side
// 	for _, coord := range edgeCoordList {
// 		square := cube.get(CubeCoord{
// 			side:      Up,
// 			faceCoord: coord,
// 		})
//
// 		if square != Up {
// 			t.Errorf("Squares shouldn't change on spinning face: %c",
// 				SideToString(square))
// 		}
// 	}
// }

func TestCubeV2Rewindable(t *testing.T) {
	cube := NewCubeV2Solved()
	r := rand.New(rand.NewPCG(0, 0))

	moveSequence := []Move{}

	for i := 0; i < 100; i++ {
		move := AllMoves[r.IntN(len(AllMoves))]
		cube.apply(move)
		moveSequence = append(moveSequence, move)

		cubeCopy := cube
		for j := i; j >= 0; j-- {
			cubeCopy.apply(moveSequence[j].Reverse())
		}

		if !cubeCopy.isSolved() {
			t.Fatalf("Doing %v in reverse didn't reverse the cube", moveSequence)
		}
	}
}

// func TestCubeV2VsCubeV1(t *testing.T) {
// 	cubeV1 := NewCubeSolved()
// 	cubeV2 := NewCubeV2Solved()
//
// 	r := rand.New(rand.NewPCG(0, 0))
// 	for i := 0; i < 100; i++ {
// 		for side := FirstSide; side <= LastSide; side++ {
// 			for _, coord := range edgeCoordList {
// 				squareSideV1 := cubeV1.get(CubeCoord{
// 					side:      side,
// 					faceCoord: coord,
// 				})
// 				squareSideV2 := cubeV2.get(CubeCoord{
// 					side:      side,
// 					faceCoord: coord,
// 				})
// 				if *squareSideV1 != squareSideV2 {
// 					t.Errorf("On side %c coord %v after %v moves, found %c expected %c",
// 						SideToString(side),
// 						coord,
// 						i,
// 						SideToString(squareSideV2),
// 						SideToString(*squareSideV1),
// 					)
// 				}
// 			}
// 		}
//
// 		if t.Failed() {
// 			// Don't test further
// 			return
// 		}
//
// 		move := AllMoves[r.IntN(len(AllMoves))]
// 		cubeV1.apply(move)
// 		cubeV2.apply(move)
// 	}
// }
