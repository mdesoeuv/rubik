package main

import (
	"fmt"
	"testing"
)

func TestFaceString(t *testing.T) {
	cube := NewCubeSolved()
	result := fmt.Sprint(cube.faces[Front])
	expected := "[F][F][F]\n[F][F][F]\n[F][F][F]\n"
	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}