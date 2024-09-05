package common

import "errors"

type Side uint8

// WARN: Do not change values
// It is required for `corner_for` to work properly
const (
	Up    Side = 0
	Down  Side = 1
	Left  Side = 2
	Right Side = 3
	Front Side = 4
	Back  Side = 5

	// Short form
	U = Up
	D = Down
	L = Left
	R = Right
	F = Front
	B = Back

	FirstSide Side = Up
	LastSide  Side = Back
	SideCount      = int(LastSide) + 1
)

func (s Side) Letter() (letter rune) {
	switch s {
	case Up:
		letter = 'U'
	case Down:
		letter = 'D'
	case Left:
		letter = 'L'
	case Right:
		letter = 'R'
	case Front:
		letter = 'F'
	case Back:
		letter = 'B'
	}
	return
}

func (s Side) String() string {
	return string(s.Letter())
}

func SideFromLetter(letter rune) (face Side, err error) {
	switch letter {
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
		err = errors.New("invalid face letter")
	}
	return
}
