package main

import (
	"fmt"
	"strings"
	"errors"
)

type Face = int 

const (
	Up 		Face = 0
	Down 	Face = iota
	Left 	Face = iota
	Right 	Face = iota
	Front 	Face = iota
	Back 	Face = iota
)

type Move struct {
	Face 			Face
	Clockwise 		bool
	NumRotations 	int
}

func (m Move) String() string {
    return fmt.Sprintf("face: %d, clockwise: %t, rotations: %d", m.Face, m.Clockwise, m.NumRotations)
}

func ParseMove(str string) (move Move, err error) {
	if len(str) == 0 || len(str) > 2 {
		err = fmt.Errorf("invalid move: %v, of length: %v", str, len(str))
		return 
	}
	move.NumRotations = 1
	move.Clockwise = true
	for i, c := range str {
		switch i {
		case 0:
			switch c {
			case 'U':
				move.Face = Up
			case 'D':
				move.Face = Down
			case 'L':
				move.Face = Left
			case 'R':
				move.Face = Right
			case 'F':
				move.Face = Front
			case 'B':
				move.Face = Back
			default:
				err = errors.New("invalid face")
			}
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