package common

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type Move struct {
	Side     Side
	Rotation Rotation
}

var AllMoves []Move = makeAllMoves()
var sideNames = map[Side]rune{Front: 'F', Back: 'B', Up: 'U', Down: 'D', Left: 'L', Right: 'R'}

func makeAllMoves() (result []Move) {
	for side := FirstSide; side <= LastSide; side++ {
		for _, rotation := range AllRotations {
			result = append(result, Move{side, rotation})
		}
	}
	return
}

func (m Move) Reverse() Move {
	m.Rotation = m.Rotation.Reverse()
	return m
}

func (m Move) IsRedudantWith(o Move) bool {
	return (o.Side == m.Side ||
		// Enforce opperation order for independant operations
		(o.Side == Right && m.Side == Left) ||
		(o.Side == Up && m.Side == Down) ||
		(o.Side == Front && m.Side == Back))
}

func ParseMove(str string) (move Move, err error) {
	if len(str) == 0 || len(str) > 2 {
		err = fmt.Errorf("invalid move: %v, of length: %v", str, len(str))
		return
	}
	var face Side
	move.Rotation = RotationClockwise()
	for i, c := range str {
		switch i {
		case 0:
			face, err = SideFromLetter(c)
			if err != nil {
				return
			}
			move.Side = face
		case 1:
			switch c {
			case '\'':
				move.Rotation = RotationAntiClockwise()
			case '2':
				move.Rotation = RotationHalfTurn()
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

const ArticleExampleSolution = "F L R' D2 B2 U " +
	"F2 D2 L R' F R2 F B2 R B' R2 B R B " +
	"L' U2 R' U2 D2 R2 D2 L' F2 L D2 L' " +
	"L2 R2 B2 R2 F2 R2 D2 U2 B2 U2 R2 U2 R2 U2"

func ArticleExampleSolutionMoveList() []Move {
	moves, _ := ParseMoveList(ArticleExampleSolution)
	return moves
}

func ArticleExampleShuffleMoveList() []Move {
	moves := ArticleExampleSolutionMoveList()
	slices.Reverse(moves)
	return moves
}

func (m Move) String() string {
	s := string(SideToString(m.Side))
	if m.Rotation.amount == -1 {
		s += "'"
	} else if m.Rotation.amount == 2 {
		s += "2"
	}
	return s
}

func SideToString(s Side) rune {
	return sideNames[s]
}
