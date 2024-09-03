package visual

import (
	"testing"

	cmn "github.com/mdesoeuv/rubik/common"
)

func TestFaceString(t *testing.T) {
	cube := NewCubeSolved()
	result := cube.faces[cmn.Front].String()
	expected := "[F][F][F]\n" +
		"[F][F][F]\n" +
		"[F][F][F]\n"
	if result != expected {
		t.Errorf("Expected: %v, got: %v", expected, result)
	}
}
