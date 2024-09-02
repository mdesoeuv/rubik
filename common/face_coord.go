package common

import (
	"fmt"
)

type FaceCoord struct {
	line   int
	column int
}

var FaceCoord00 = FaceCoord{line: 0, column: 0}
var FaceCoord01 = FaceCoord{line: 0, column: 1}
var FaceCoord02 = FaceCoord{line: 0, column: 2}
var FaceCoord10 = FaceCoord{line: 1, column: 0}
var FaceCoord11 = FaceCoord{line: 1, column: 1}
var FaceCoord12 = FaceCoord{line: 1, column: 2}
var FaceCoord20 = FaceCoord{line: 2, column: 0}
var FaceCoord21 = FaceCoord{line: 2, column: 1}
var FaceCoord22 = FaceCoord{line: 2, column: 2}

func (fc FaceCoord) isValid() bool {
	return 0 <= fc.line && fc.line < 3 && 0 <= fc.column && fc.column < 3
}

func (fc FaceCoord) String() string {
	return fmt.Sprintf("{%v, %v}", fc.line, fc.column)
}

func NewFaceCoord(line, column int) (fc FaceCoord, err error) {
	fc = FaceCoord{line: line, column: column}
	if !fc.isValid() {
		err = fmt.Errorf("Invalid coordinates: %v", fc)
		fc = FaceCoord{}
	}
	return
}

func (fc FaceCoord) Line() int {
	return fc.line
}

func (fc FaceCoord) Column() int {
	return fc.column
}

func (fc FaceCoord) IsEdge() bool {
	return fc.line == 1 || fc.column == 1
}

func (fc FaceCoord) IsCorner() bool {
	return !fc.IsEdge()
}
