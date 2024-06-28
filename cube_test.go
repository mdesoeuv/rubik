package main

import "testing"

var colors = []Color{Red, Blue, Green, White, Yellow, Orange}


func sliceIsOfColor(array [] Color, expected Color) bool {
	for _, color := range array {
		if color != expected {
			return false
		}
	}
	return true
}

func TestSortedCubeMoveUp(t *testing.T) {
	cube := NewCubeSolved()
	leftColor := cube.faces[Left][0][0]
	rightColor := cube.faces[Right][0][0]
	frontColor := cube.faces[Front][0][0]
	backColor := cube.faces[Back][0][0]
	upColor := cube.faces[Up][0][0]
	downColor := cube.faces[Down][0][0]

	cube.apply(Move{Up, true, 1})

	// Check that the up face is rotated
	if !sliceIsOfColor(cube.faces[Right][0][:], backColor) {
		t.Errorf("Right face did not rotate correctly")
	}
	if !sliceIsOfColor(cube.faces[Front][0][:], rightColor) {
		t.Errorf("Front face did not rotate correctly")
	}
	if !sliceIsOfColor(cube.faces[Back][0][:], leftColor) {
		t.Errorf("Back face did not rotate correctly")
	}
	if !sliceIsOfColor(cube.faces[Left][0][:], frontColor) {
		t.Errorf("Left face did not rotate correctly")
	}

	// Check that the other squares are unchanged
	for _, line := range cube.faces[Up][:] {
		if !sliceIsOfColor(line[:], upColor) {
			t.Errorf("Up face has changed")	
		}
	}
	for _, line := range cube.faces[Left][1:] {
		if !sliceIsOfColor(line[:], leftColor) {
			t.Errorf("Left: unexpected square change")	
		}
	}
	for _, line := range cube.faces[Right][1:] {
		if !sliceIsOfColor(line[:], rightColor) {
			t.Errorf("Right: unexpected square change")	
		}
	}
	for _, line := range cube.faces[Front][1:] {
		if !sliceIsOfColor(line[:], frontColor) {
			t.Errorf("Front: unexpected square change")	
		}
	}
	for _, line := range cube.faces[Back][1:] {
		if !sliceIsOfColor(line[:], backColor) {
			t.Errorf("Back: unexpected square change")	
		}
	}
	for _, line := range cube.faces[Down][:] {
		if !sliceIsOfColor(line[:], downColor) {
			t.Errorf("Up face has changed")	
		}
	}
}

func NewCubeTopSpinned() *Cube {
	cube := NewCubeSolved()

	// Spin top
	frontLine := cube.faces[Front][0]
	rightLine := cube.faces[Right][0]
	backLine := cube.faces[Back][0]
	leftLine := cube.faces[Left][0]
	cube.faces[Front][0] = leftLine
	cube.faces[Right][0] = frontLine
	cube.faces[Back][0] = rightLine
	cube.faces[Left][0] = backLine
	
	return cube
}

func TestFrontFaceRotation(t *testing.T) {
	cube := NewCubeTopSpinned()

	move := Move {
		Side:         Front,
		Clockwise:    true,
		NumRotations: 1,
	}

	cube.apply(move)

	expectedFace := Face {
		{Front, Front, Left},
		{Front, Front, Left},
		{Front, Front, Left},
	}

	if !FaceEqual(cube.faces[Front], expectedFace) {
		t.Errorf("Wrong front face after rotation")
	}
}
