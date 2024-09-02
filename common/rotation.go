package common

type Rotation struct {
	// TODO: Make it a positive number
	amount int8
}

var AllRotations = []Rotation{
	RotationAntiClockwise(),
	RotationClockwise(),
	RotationHalfTurn(),
}

func RotationAntiClockwise() Rotation {
	return Rotation{-1}
}

func RotationNone() Rotation {
	return Rotation{0}
}

func RotationClockwise() Rotation {
	return Rotation{1}
}

func RotationHalfTurn() Rotation {
	return Rotation{2}
}

func (r Rotation) String() string {
	switch r.amount {
	case -1:
		return "anti-clockwise"
	case 0:
		return "none"
	case 1:
		return "clockwise"
	case 2:
		return "half-turn"
	default:
		return "INVALID"
	}
}

// TODO: Use math instead
func (r Rotation) Reverse() Rotation {
	switch r.amount {
	case -1:
		r.amount = 1
	case 1:
		r.amount = -1
	}
	return r
}

func (r Rotation) Int() int {
	return int(r.amount)
}

func (r Rotation) PositiveInt() int {
	return int(r.amount+4) & 3
}

func (r Rotation) IsQuaterTurn() bool {
	return r.amount&1 != 0
}
