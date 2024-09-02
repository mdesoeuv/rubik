package visual

import (
	"fmt"
	"testing"

	cmn "github.com/mdesoeuv/rubik/common"
)

func sliceIsOfSide(array []cmn.Side, expected cmn.Side) bool {
	for _, side := range array {
		if side != expected {
			return false
		}
	}
	return true
}

func TestSortedCubeMoveUp(t *testing.T) {
	cube := NewCubeSolved()
	leftSide := cube.faces[cmn.Left].f[0][0]
	rightSide := cube.faces[cmn.Right].f[0][0]
	frontSide := cube.faces[cmn.Front].f[0][0]
	backSide := cube.faces[cmn.Back].f[0][0]
	upSide := cube.faces[cmn.Up].f[0][0]
	downSide := cube.faces[cmn.Down].f[0][0]

	cube.Apply(cmn.Move{Side: cmn.Up, Rotation: cmn.RotationClockwise()})

	// Check that the up face is rotated
	if !sliceIsOfSide(cube.faces[cmn.Right].f[0][:], backSide) {
		t.Errorf("Right face did not rotate correctly")
	}
	if !sliceIsOfSide(cube.faces[cmn.Front].f[0][:], rightSide) {
		t.Errorf("Front face did not rotate correctly")
	}
	if !sliceIsOfSide(cube.faces[cmn.Back].f[0][:], leftSide) {
		t.Errorf("Back face did not rotate correctly")
	}
	if !sliceIsOfSide(cube.faces[cmn.Left].f[0][:], frontSide) {
		t.Errorf("Left face did not rotate correctly")
	}

	// Check that the other squares are unchanged
	for _, line := range cube.faces[cmn.Up].f[:] {
		if !sliceIsOfSide(line[:], upSide) {
			t.Errorf("Up face has changed")
		}
	}
	for _, line := range cube.faces[cmn.Left].f[1:] {
		if !sliceIsOfSide(line[:], leftSide) {
			t.Errorf("Left: unexpected square change")
		}
	}
	for _, line := range cube.faces[cmn.Right].f[1:] {
		if !sliceIsOfSide(line[:], rightSide) {
			t.Errorf("Right: unexpected square change")
		}
	}
	for _, line := range cube.faces[cmn.Front].f[1:] {
		if !sliceIsOfSide(line[:], frontSide) {
			t.Errorf("Front: unexpected square change")
		}
	}
	for _, line := range cube.faces[cmn.Back].f[1:] {
		if !sliceIsOfSide(line[:], backSide) {
			t.Errorf("Back: unexpected square change")
		}
	}
	for _, line := range cube.faces[cmn.Down].f[:] {
		if !sliceIsOfSide(line[:], downSide) {
			t.Errorf("Up face has changed")
		}
	}
}

func NewCubeTopSpinned() *Cube {
	cube := NewCubeSolved()

	// Spin top
	frontLine := cube.faces[cmn.Front].f[0]
	rightLine := cube.faces[cmn.Right].f[0]
	backLine := cube.faces[cmn.Back].f[0]
	leftLine := cube.faces[cmn.Left].f[0]
	cube.faces[cmn.Front].f[0] = leftLine
	cube.faces[cmn.Right].f[0] = frontLine
	cube.faces[cmn.Back].f[0] = rightLine
	cube.faces[cmn.Left].f[0] = backLine

	return cube
}

func TestFrontFaceRotation(t *testing.T) {
	cube := NewCubeTopSpinned()

	move := cmn.Move{
		Side:     cmn.Front,
		Rotation: cmn.RotationClockwise(),
	}

	cube.Apply(move)

	expectedFace := Face{
		f: [3][3]cmn.Side{
			{cmn.Front, cmn.Front, cmn.Left},
			{cmn.Front, cmn.Front, cmn.Left},
			{cmn.Front, cmn.Front, cmn.Left},
		},
	}

	if !FaceEqual(cube.faces[cmn.Front], expectedFace) {
		t.Errorf("Wrong front face after rotation")
		t.Errorf("Expected:\n%vFound:\n%v", expectedFace, cube.faces[cmn.Front])
		fmt.Println(cube.Blueprint())
	}
}

func TestCubeString(t *testing.T) {
	cube := NewCubeSolved()
	result := cube.String()
	expected := "Face: U\n" +
		"[U][U][U]\n" +
		"[U][U][U]\n" +
		"[U][U][U]\n\n" +
		"Face: D\n" +
		"[D][D][D]\n" +
		"[D][D][D]\n" +
		"[D][D][D]\n\n" +
		"Face: L\n" +
		"[L][L][L]\n" +
		"[L][L][L]\n" +
		"[L][L][L]\n\n" +
		"Face: R\n" +
		"[R][R][R]\n" +
		"[R][R][R]\n" +
		"[R][R][R]\n\n" +
		"Face: F\n" +
		"[F][F][F]\n" +
		"[F][F][F]\n" +
		"[F][F][F]\n\n" +
		"Face: B\n" +
		"[B][B][B]\n" +
		"[B][B][B]\n" +
		"[B][B][B]\n\n"
	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}

func TestCubeCopy(t *testing.T) {
	cube := NewCubeSolved()

	duplicate := *cube

	cube.faces[cmn.Front].f[0][0] = cmn.Back

	if FaceEqual(cube.faces[cmn.Front], duplicate.faces[cmn.Front]) {
		t.Errorf("Cubes are not supposed to be the same ")
	}
}

func BenchmarkCubeApply(b *testing.B) {
	cube := NewCubeSolved()

	b.StartTimer()
	for i := 0; i <= b.N; i += 1 {
		cube.Apply(cmn.AllMoves[i%len(cmn.AllMoves)])
	}
	b.StopTimer()
}
