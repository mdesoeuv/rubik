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



