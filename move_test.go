package main

import "testing"

func TestParseMove(t *testing.T) {
	move, err := ParseMove("D2")
	expectedMove := Move{
		Face:         Down,
		Clockwise:    true,
		NumRotations: 2,
	}
	if err != nil || move != expectedMove {
		t.FailNow()
	}
}
