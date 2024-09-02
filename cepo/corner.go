package cepo

import (
	"math/bits"

	cmn "github.com/mdesoeuv/rubik/common"
)

type CornerOrientation = uint8

const (
	LeftRight CornerOrientation = 0
	UpDown    CornerOrientation = 1
	FrontBack CornerOrientation = 2

	// Short forms
	LF = LeftRight
	UD = UpDown
	FB = FrontBack

	FirstCornerOrientation = LeftRight
	LastCornerOrientation  = FrontBack
	CornerOrientationCount = 3
)

type CornerOrientations struct {
	bits uint32
}
type CornerPermutationBits = uint32
type CornerPermutation struct {
	bits CornerPermutationBits
}

func NewCornerOrientationsSolved() CornerOrientations {
	return CornerOrientations{bits: 0}
}

func (co CornerOrientations) IsSolved() bool {
	return co.bits == 0
}

func (co CornerOrientations) Get(index cmn.CornerIndex) CornerOrientation {
	var i = index - 1
	return CornerOrientation((co.bits >> (i * 2)) & 3)
}

func (co *CornerOrientations) Set(index cmn.CornerIndex, orientation CornerOrientation) {
	var i = index - 1
	co.bits &= ^(3 << (i * 2))                  // Clear bits
	co.bits |= (uint32(orientation) << (i * 2)) // Set value
}

func (co *CornerOrientations) Apply(move cmn.Move) {
	nco := *co
	crown := &cmn.CrownCornerList[move.Side]
	rotation := ((move.Side >> 1) + 2) << 1
	for i := range crown {
		from, to := crown[i], crown[(i+move.Rotation.PositiveInt())%len(crown)]
		orientation := co.Get(from)
		// TODO: Remove condition
		if move.Rotation.IsQuaterTurn() {
			orientation = ((orientation << 1) + uint8(rotation<<1)) % CornerOrientationCount
		}
		nco.Set(cmn.CornerIndex(to), orientation)
	}
	*co = nco
}

func (co CornerOrientations) Distance() int {
	// There is only 1 one for each mis oriented corner
	return (bits.OnesCount32(co.bits) + 3) / 4
}

func NewCornerPermutationSolved() CornerPermutation {
	cp := CornerPermutation{bits: 0}
	for index := cmn.FirstCornerIndex; index <= cmn.LastCornerIndex; index++ {
		cp.Set(index, index)
	}
	return cp
}

func (cp CornerPermutation) Get(index cmn.CornerIndex) cmn.CornerIndex {
	var i = index - 1
	return cmn.CornerIndex((cp.bits>>(i*3))&0x7 + 1)
}

func (cp *CornerPermutation) Set(index, value cmn.CornerIndex) {
	var i = index - 1
	cp.bits &= ^CornerPermutationBits(0x7 << (i * 3))    // Clear edge
	cp.bits |= CornerPermutationBits(value-1) << (i * 3) // Set value
}

func (cp *CornerPermutation) Apply(move cmn.Move) {
	ncp := *cp
	crown := &cmn.CrownCornerList[move.Side]
	for i := range crown {
		from, to := crown[i], crown[(i+move.Rotation.PositiveInt())%len(crown)]
		ncp.Set(cmn.CornerIndex(to), cp.Get(from))
	}
	*cp = ncp
}

func (cp CornerPermutation) IsSolved() bool {
	return cp == NewCornerPermutationSolved()
}
