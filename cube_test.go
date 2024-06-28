package main

import (
	"fmt"
	"testing"
)

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
	leftColor := cube.faces[Left].f[0][0]
	rightColor := cube.faces[Right].f[0][0]
	frontColor := cube.faces[Front].f[0][0]
	backColor := cube.faces[Back].f[0][0]
	upColor := cube.faces[Up].f[0][0]
	downColor := cube.faces[Down].f[0][0]

	cube.apply(Move{Up, true, 1})

	// Check that the up face is rotated
	if !sliceIsOfColor(cube.faces[Right].f[0][:], backColor) {
		t.Errorf("Right face did not rotate correctly")
	}
	if !sliceIsOfColor(cube.faces[Front].f[0][:], rightColor) {
		t.Errorf("Front face did not rotate correctly")
	}
	if !sliceIsOfColor(cube.faces[Back].f[0][:], leftColor) {
		t.Errorf("Back face did not rotate correctly")
	}
	if !sliceIsOfColor(cube.faces[Left].f[0][:], frontColor) {
		t.Errorf("Left face did not rotate correctly")
	}

	// Check that the other squares are unchanged
	for _, line := range cube.faces[Up].f[:] {
		if !sliceIsOfColor(line[:], upColor) {
			t.Errorf("Up face has changed")	
		}
	}
	for _, line := range cube.faces[Left].f[1:] {
		if !sliceIsOfColor(line[:], leftColor) {
			t.Errorf("Left: unexpected square change")	
		}
	}
	for _, line := range cube.faces[Right].f[1:] {
		if !sliceIsOfColor(line[:], rightColor) {
			t.Errorf("Right: unexpected square change")	
		}
	}
	for _, line := range cube.faces[Front].f[1:] {
		if !sliceIsOfColor(line[:], frontColor) {
			t.Errorf("Front: unexpected square change")	
		}
	}
	for _, line := range cube.faces[Back].f[1:] {
		if !sliceIsOfColor(line[:], backColor) {
			t.Errorf("Back: unexpected square change")	
		}
	}
	for _, line := range cube.faces[Down].f[:] {
		if !sliceIsOfColor(line[:], downColor) {
			t.Errorf("Up face has changed")	
		}
	}
}

func NewCubeTopSpinned() *Cube {
	cube := NewCubeSolved()

	// Spin top
	frontLine := cube.faces[Front].f[0]
	rightLine := cube.faces[Right].f[0]
	backLine := cube.faces[Back].f[0]
	leftLine := cube.faces[Left].f[0]
	cube.faces[Front].f[0] = leftLine
	cube.faces[Right].f[0] = frontLine
	cube.faces[Back].f[0] = rightLine
	cube.faces[Left].f[0] = backLine
	
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
		f: [3][3]Color {
		{Front, Front, Left},
		{Front, Front, Left},
		{Front, Front, Left},
		},
	}

	if !FaceEqual(cube.faces[Front], expectedFace) {
		t.Errorf("Wrong front face after rotation")
	}
}

func TestCubeString(t *testing.T) {
	cube := NewCubeSolved()
	result := fmt.Sprint(cube)
	expected := "Face: U\n[U][U][U]\n[U][U][U]\n[U][U][U]\n\nFace: D\n[D][D][D]\n[D][D][D]\n[D][D][D]\n\nFace: L\n[L][L][L]\n[L][L][L]\n[L][L][L]\n\nFace: R\n[R][R][R]\n[R][R][R]\n[R][R][R]\n\nFace: F\n[F][F][F]\n[F][F][F]\n[F][F][F]\n\nFace: B\n[B][B][B]\n[B][B][B]\n[B][B][B]\n\n"
	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
