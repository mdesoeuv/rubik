package common_test

import (
	"testing"

	"github.com/mdesoeuv/rubik/common"
)

func TestRotationReverse(t *testing.T) {
	reverseTable := []struct{ a, b common.Rotation }{
		{common.RotationNone(), common.RotationNone()},
		{common.RotationClockwise(), common.RotationAntiClockwise()},
		{common.RotationAntiClockwise(), common.RotationClockwise()},
		{common.RotationHalfTurn(), common.RotationHalfTurn()},
	}

	for _, entry := range reverseTable {
		if entry.a.Reverse() != entry.b {
			t.Errorf("Reverse of %v should be %v", entry.a, entry.b)
		}
	}
}

func TestRotationInt(t *testing.T) {
	reverseTable := []struct {
		a common.Rotation
		b int
	}{
		{common.RotationAntiClockwise(), -1},
		{common.RotationNone(), 0},
		{common.RotationClockwise(), 1},
		{common.RotationHalfTurn(), 2},
	}

	for _, entry := range reverseTable {
		if entry.a.Int() != entry.b {
			t.Errorf("Int of %v should be %v", entry.a, entry.b)
		}
	}
}

func TestRotationPositiveInt(t *testing.T) {
	reverseTable := []struct {
		a common.Rotation
		b int
	}{
		{common.RotationNone(), 0},
		{common.RotationClockwise(), 1},
		{common.RotationHalfTurn(), 2},
		{common.RotationAntiClockwise(), 3},
	}

	for _, entry := range reverseTable {
		if entry.a.PositiveInt() != entry.b {
			t.Errorf("Positive int of %v should be %v", entry.a, entry.b)
		}
	}
}
