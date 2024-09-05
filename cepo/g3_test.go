package cepo

import (
	"testing"
)

func TestMakeG3Table(t *testing.T) {
	table := MakeG3HeuristicTable()

	maxDistance := uint8(0)
	for _, distance := range table {
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	if maxDistance != 15 {
		t.Errorf("Max distance should be 15 but found: %v", maxDistance)
	}
	foundSize := len(table)
	expectedSize := 663_552
	if foundSize != expectedSize {
		t.Fatal("G3Cubes should have size from article")
	}
}
