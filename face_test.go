package main

import (
	"testing"
)

func TestFaceString(t *testing.T) {
	cube := NewCubeSolved()
	result := cube.faces[Front].String()
	expected := "[F][F][F]\n" +
				"[F][F][F]\n" +
				"[F][F][F]\n"
	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}