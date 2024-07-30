package main

import (
	"github.com/fatih/color"
)

func (c *Cube) blueprint() string {
	lines := [9]string{}
	emptyLine := "         "
	lines[0] = emptyLine + " " + c.faces[Up].FaceGetLineString(0)
	lines[1] = emptyLine + " " + c.faces[Up].FaceGetLineString(1)
	lines[2] = emptyLine + " " + c.faces[Up].FaceGetLineString(2) + "\n"

	lines[3] = c.faces[Left].FaceGetLineString(0) + " " + c.faces[Front].FaceGetLineString(0) + " " + c.faces[Right].FaceGetLineString(0) + " " + c.faces[Back].FaceGetLineString(0)
	lines[4] = c.faces[Left].FaceGetLineString(1) + " " + c.faces[Front].FaceGetLineString(1) + " " + c.faces[Right].FaceGetLineString(1) + " " + c.faces[Back].FaceGetLineString(1)
	lines[5] = c.faces[Left].FaceGetLineString(2) + " " + c.faces[Front].FaceGetLineString(2) + " " + c.faces[Right].FaceGetLineString(2) + " " + c.faces[Back].FaceGetLineString(2) + "\n"

	lines[6] = emptyLine + " " + c.faces[Down].FaceGetLineString(0)
	lines[7] = emptyLine + " " + c.faces[Down].FaceGetLineString(1)
	lines[8] = emptyLine + " " + c.faces[Down].FaceGetLineString(2)

	s := ""
	for _, line := range lines {
		s += line + "\n"
	}
	return s
}

var theme = map[Side]func(a ...interface{}) string{
	Right: color.New(color.BgRed).SprintFunc(),
	Front: color.New(color.BgBlue).SprintFunc(),
	Back:  color.New(color.BgGreen).SprintFunc(),
	Up:    color.New(color.BgWhite).SprintFunc(),
	Down:  color.New(color.BgYellow).SprintFunc(),
	Left:  color.New(color.BgMagenta).SprintFunc(),
}
