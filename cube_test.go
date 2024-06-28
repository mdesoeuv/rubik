package main

// import "testing"

var colors = []Color{Red, Blue, Green, White, Yellow, Orange}

func GetUniformFace(color Color) (face Face) {
	for i, line := range face {
		for j, _ := range line {
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
	for i, line := range cube.faces[Left][0] {
		if line[i] != frontColor {
			t.FailNow()
		}
	}
	for i, line := range cube.faces[Right][0] {
		if line[i] != backColor {
			t.FailNow()
		}
	}
	for i, line := range cube.faces[Front][0] {
		if line[i] != rightColor {
			t.FailNow()
		}
	}
	for i, line := range cube.faces[Back][0] {
		if line[i] != leftColor {
			t.FailNow()
		}
	}

	// Check that the other squares are unchanged
	for i, line := range cube.faces[Up][1:2] {
		for j, color := range line {
			if color != upColor {
				t.FailNow()
			}
		}
	}
	for i, line := range cube.faces[Left][1:2] {
		for j, color := range line {
			if color != leftColor {
				t.FailNow()
			}
		}
	}
	for i, line := range cube.faces[Right][1:2] {
		for j, color := range line {
			if color != rightColor {
				t.FailNow()
			}
		}
	}
	for i, line := range cube.faces[Front][1:2] {
		for j, color := range line {
			if color != frontColor {
				t.FailNow()
			}
		}
	}
	for i, line := range cube.faces[Back][1:2] {
		for j, color := range line {
			if color != backColor {
				t.FailNow()
			}
		}
	}
	for i, line := range cube.faces[Down][0:3] {
		for j, color := range line {
			if color != downColor {
				t.FailNow()
			}
		}
	}
}

