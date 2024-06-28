package main

import "testing"

var colors = []Color{Red, Blue, Green, White, Yellow, Orange}

func GetUniformFace(color Color) (face Face) {
	for i, line := range face {
		for j := range line {
			face[i][j] = color
		}
	}
	return
}


func GetSortedCube() (cube *Cube) {

	cube = NewCube()
	cube.faces[Front] = GetUniformFace(Red)
	cube.faces[Back] = GetUniformFace(Blue)
	cube.faces[Left] = GetUniformFace(Green)
	cube.faces[Right] = GetUniformFace(White)
	cube.faces[Up] = GetUniformFace(Yellow)
	cube.faces[Down] = GetUniformFace(Orange)
	return
}

func sliceIsOfColor(array [] Color, expected Color) bool {
	for _, color := range array {
		if color != expected {
			return false
		}
	}
	return true
}

func TestSortedCubeMoveUp(t *testing.T) {
	cube := GetSortedCube()
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
