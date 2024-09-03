package visual

import (
	"github.com/fatih/color"

	cmn "github.com/mdesoeuv/rubik/common"
)

func (c *Cube) Blueprint() string {
	lines := [9]string{}
	emptyLine := "         "
	lines[0] = emptyLine + " " + c.faces[cmn.Up].FaceGetLineString(0)
	lines[1] = emptyLine + " " + c.faces[cmn.Up].FaceGetLineString(1)
	lines[2] = emptyLine + " " + c.faces[cmn.Up].FaceGetLineString(2) + "\n"

	lines[3] = c.faces[cmn.Left].FaceGetLineString(0) + " " + c.faces[cmn.Front].FaceGetLineString(0) + " " + c.faces[cmn.Right].FaceGetLineString(0) + " " + c.faces[cmn.Back].FaceGetLineString(0)
	lines[4] = c.faces[cmn.Left].FaceGetLineString(1) + " " + c.faces[cmn.Front].FaceGetLineString(1) + " " + c.faces[cmn.Right].FaceGetLineString(1) + " " + c.faces[cmn.Back].FaceGetLineString(1)
	lines[5] = c.faces[cmn.Left].FaceGetLineString(2) + " " + c.faces[cmn.Front].FaceGetLineString(2) + " " + c.faces[cmn.Right].FaceGetLineString(2) + " " + c.faces[cmn.Back].FaceGetLineString(2) + "\n"

	lines[6] = emptyLine + " " + c.faces[cmn.Down].FaceGetLineString(0)
	lines[7] = emptyLine + " " + c.faces[cmn.Down].FaceGetLineString(1)
	lines[8] = emptyLine + " " + c.faces[cmn.Down].FaceGetLineString(2)

	s := ""
	for i, line := range lines {
		s += line
		if i < len(lines)-1 {
			s += "\n"
		}
	}
	return s
}

var theme = map[cmn.Side]func(a ...interface{}) string{
	cmn.Right: color.New(color.BgRed).SprintFunc(),
	cmn.Front: color.New(color.BgBlue).SprintFunc(),
	cmn.Back:  color.New(color.BgGreen).SprintFunc(),
	cmn.Up:    color.New(color.BgWhite).SprintFunc(),
	cmn.Down:  color.New(color.BgYellow).SprintFunc(),
	cmn.Left:  color.New(color.BgMagenta).SprintFunc(),
}
