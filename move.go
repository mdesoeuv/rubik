package main

import (
	"fmt"
	"strings"
	"errors"
)

type Side = int 

const (
	Up 		Side = 0
	Down 	Side = iota
	Left 	Side = iota
	Right 	Side = iota
	Front 	Side = iota
	Back 	Side = iota
	FirstSide Side = Up
	LastSide = Back
    SideCount = LastSide + 1
)

var sideNames = map[Side]rune{Front: 'F', Back: 'B', Up: 'U', Down: 'D', Left: 'L', Right: 'R'}

type Move struct {
	Side 			Side
	Clockwise 		bool
	NumRotations 	int
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
    return fmt.Sprintf("face: %c, clockwise: %t, rotations: %d", SideToString(m.Side), m.Clockwise, m.NumRotations)
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
	move.Clockwise = true
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
				move.Clockwise = false
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
