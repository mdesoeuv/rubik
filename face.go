package main

import (
	"fmt"
)

type Face struct {
	f [3][3]Side
}

func (f Face) String() string {
	result := ""
	for _, line := range f.f {
		for _, side := range line {
			result += fmt.Sprintf("[%c]", sideNames[side])
		}
		result += "\n"
	}
	return result
}

func (f Face) FaceGetLineString(line int) string {
	result := ""
	for _, side := range f.f[line] {
		result += fmt.Sprintf("[%c]", sideNames[side])
	}
	return result
}

func NewFaceUniform(side Side) (face Face) {
	for i, line := range face.f {
		for j := range line {
			face.f[i][j] = side
		}
	}
	return
}

func FaceEqual(a, b Face) bool {
	for line, _ := range a.f {
		for column, _ := range a.f[line] {
			if a.f[line][column] != b.f[line][column] {
				return false
			}
		}
	}
	return true
}

func FaceIsUniform(face Face, side Side) bool {
	for line := range face.f {
		for column := range face.f[line] {
			if face.f[line][column] != side {
				return false
			}
		}
	}
	return true
}
