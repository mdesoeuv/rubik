package main

type Color = int
const (
	Red Color = 0
	Blue Color = iota
	Green Color = iota
	White Color = iota
	Yellow Color = iota
	Orange Color = iota
	ColorCount int = iota
)

type Cube struct {
	faces [6]Face
}

type Face = [3][3]Color

func NewCube() *Cube {
	return &Cube{}
}

func (c *Cube)apply(move Move) {
	panic("Todo")
}

func (c *Cube)isSolved() bool {
	seenColors := [ColorCount]bool{}
	for _, face := range c.faces {
		var currentColor *Color = nil
		for _, line := range face {
			for _, color := range line {
				if currentColor != nil  && color != *currentColor {
					return false
				}
				currentColor = &color
			}
		}
		if seenColors[*currentColor] {
			return false
		}
		seenColors[*currentColor] = true
	}
	return true
}
