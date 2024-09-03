package cepo

import (
	"testing"
)

func TestMakeG3CubesHasSizeFromArticle(t *testing.T) {
	G3Cubes := MakeG3Cubes()

	foundSize := len(G3Cubes)
	expectedSize := 663_552
	if foundSize != expectedSize {
		t.Fatal("G3Cubes should have size from article")
	}
}
