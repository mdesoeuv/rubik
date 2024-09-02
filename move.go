package main

import (
	"errors"
	"fmt"
	"strings"
)

type Side = uint8

// WARN: Do not change values
// It is required for `corner_for` to work properly
const (
	Up        Side = 0
	Down      Side = 1
	Left      Side = 2
	Right     Side = 3
	Front     Side = 4
	Back      Side = 5
	FirstSide Side = Up
	LastSide  Side = Back
	SideCount      = int(LastSide) + 1
)

var sideNames = map[Side]rune{Front: 'F', Back: 'B', Up: 'U', Down: 'D', Left: 'L', Right: 'R'}

type Move struct {
	Side Side
	// TODO: Make it a positive number
	NumRotations int
}

var AllMoves []Move = makeAllMoves()

func makeAllMoves() (result []Move) {
	rotations := []int{-1, 1, 2}
	for side := FirstSide; side <= LastSide; side++ {
		for _, numRotation := range rotations {
			result = append(result, Move{side, numRotation})
		}
	}
	return result
}

func (m Move) Reverse() Move {
	switch m.NumRotations {
	case -1:
		m.NumRotations = 1
	case 1:
		m.NumRotations = -1
	}
	return m
}

func ParseSide(c rune) (face Side, err error) {
	switch c {
	case 'U':
		face = Up
	case 'D':
		face = Down
	case 'L':
		face = Left
	case 'R':
		face = Right
	case 'F':
		face = Front
	case 'B':
		face = Back
	default:
		err = errors.New("invalid face")
	}
	return
}

func (m Move) String() string {
	return fmt.Sprintf("face: %c, rotations: %d", SideToString(m.Side), m.NumRotations)
}

func (m Move) CompactString() string {
	s := string(SideToString(m.Side))
	if m.NumRotations == -1 {
		s += "'"
	} else if m.NumRotations == 2 {
		s += "2"
	}
	return s
}

func SideToString(s Side) rune {
	return sideNames[s]
}

func ParseMove(str string) (move Move, err error) {
	if len(str) == 0 || len(str) > 2 {
		err = fmt.Errorf("invalid move: %v, of length: %v", str, len(str))
		return
	}
	var face Side
	move.NumRotations = 1
	for i, c := range str {
		switch i {
		case 0:
			face, err = ParseSide(c)
			if err != nil {
				return
			}
			move.Side = face
		case 1:
			switch c {
			case '\'':
				move.NumRotations = -1
			case '2':
				move.NumRotations = 2
			default:
				err = errors.New("invalid rotation")
				return
			}
		}
	}
	return
}

func ParseMoveList(str string) (moveList []Move, err error) {
	strArray := strings.Split(str, " ")
	for _, moveStr := range strArray {
		move, e := ParseMove(moveStr)
		if e != nil {
			return nil, e
		}
		moveList = append(moveList, move)
	}
	return
}
