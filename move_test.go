package main

import "testing"

// Move Parsing Tests
func TestParseMove(t *testing.T) {

	move, err := ParseMove("D2")
	expectedMove := Move{
		Side:         Down,
		NumRotations: 2,
	}
	if err != nil || move != expectedMove {
		t.FailNow()
	}
}

func TestParseMoveSingle(t *testing.T) {
	move, err := ParseMove("D")
	expectedMove := Move{
		Side:         Down,
		NumRotations: 1,
	}
	if err != nil || move != expectedMove {
		t.FailNow()
	}
}

func TestParseMoveCounterClockwise(t *testing.T) {
	move, err := ParseMove("D'")
	expectedMove := Move{
		Side:         Down,
		NumRotations: -1,
	}
	if err != nil || move != expectedMove {
		t.FailNow()
	}
}

func TestParseMoveInvalidLength(t *testing.T) {
	_, err := ParseMove("D2'")
	if err == nil {
		t.FailNow()
	}
}

func TestParseMoveInvalidRotation(t *testing.T) {
	_, err := ParseMove("D3")
	if err == nil {
		t.FailNow()
	}
}
