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
	Bits uint32
}
type CornerPermutationBits = uint32
type CornerPermutation struct {
	Bits CornerPermutationBits
}

func NewCornerOrientationsSolved() CornerOrientations {
	return CornerOrientations{Bits: 0}
}

func (co CornerOrientations) IsSolved() bool {
	return co.Bits == 0
}

func (co CornerOrientations) Get(index cmn.CornerIndex) CornerOrientation {
	var i = index - 1
	return CornerOrientation((co.Bits >> (i * 2)) & 3)
}

func (co *CornerOrientations) Set(index cmn.CornerIndex, orientation CornerOrientation) {
	var i = index - 1
	co.Bits &= ^(3 << (i * 2))                  // Clear bits
	co.Bits |= (uint32(orientation) << (i * 2)) // Set value
}

func (co *CornerOrientations) Apply(move cmn.Move) {
	nco := *co
	crown := &cmn.CrownCornerList[move.Side]
	rotation := ((move.Side >> 1) + 2) << 1
	for i := range crown {
		from, to := crown[i], crown[(i+move.Rotation.PositiveInt())%len(crown)]
		orientation := co.Get(from)
		if move.Rotation.IsQuaterTurn() {
			orientation = ((orientation << 1) + uint8(rotation<<1)) % CornerOrientationCount
		}
		nco.Set(cmn.CornerIndex(to), orientation)
	}
	*co = nco
}

func (co CornerOrientations) Distance() int {
	// There is only 1 one for each mis oriented corner
	return (bits.OnesCount32(co.Bits) + 3) / 4
}

func NewCornerPermutationSolved() CornerPermutation {
	cp := CornerPermutation{Bits: 0}
	for index := cmn.FirstCornerIndex; index <= cmn.LastCornerIndex; index++ {
		cp.Set(index, index)
	}
	return cp
}

func (cp CornerPermutation) Get(index cmn.CornerIndex) cmn.CornerIndex {
	var i = index - 1
	return cmn.CornerIndex((cp.Bits>>(i*3))&0x7 + 1)
}

func (cp *CornerPermutation) Set(index, value cmn.CornerIndex) {
	var i = index - 1
	cp.Bits &= ^CornerPermutationBits(0x7 << (i * 3))    // Clear edge
	cp.Bits |= CornerPermutationBits(value-1) << (i * 3) // Set value
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

func (cp CornerPermutation) AllInCorrectOrbit() bool {
	orbit_1 := 1 << cp.Get(cmn.ULB)
	orbit_1 += 1 << cp.Get(cmn.DLF)
	orbit_1 += 1 << cp.Get(cmn.DRB)
	orbit_1 += 1 << cp.Get(cmn.URF)
	orbit_2 := 1 << cp.Get(cmn.ULF)
	orbit_2 += 1 << cp.Get(cmn.DLB)
	orbit_2 += 1 << cp.Get(cmn.DRF)
	orbit_2 += 1 << cp.Get(cmn.URB)
	return orbit_1 == 0b0000_1111_0 && orbit_2 == 0b1111_0000_0
}

func (cp CornerPermutation) AllInCorrectOrbitDistance() int {
	count := 0
	if cp.Get(cmn.ULB) > cmn.URF {
		count += 1
	}
	if cp.Get(cmn.DLF) > cmn.URF {
		count += 1
	}
	if cp.Get(cmn.DRB) > cmn.URF {
		count += 1
	}
	if cp.Get(cmn.URF) > cmn.URF {
		count += 1
	}
	if cp.Get(cmn.ULF) < cmn.ULF {
		count += 1
	}
	if cp.Get(cmn.DLB) < cmn.ULF {
		count += 1
	}
	if cp.Get(cmn.DRF) < cmn.ULF {
		count += 1
	}
	if cp.Get(cmn.URB) < cmn.ULF {
		count += 1
	}
	return (count + 3) / 4
}

func (cp CornerPermutation) Distance() int {
	count := 0
	for corner := cmn.FirstCornerIndex; corner <= cmn.LastCornerIndex; corner++ {
		if cp.Get(corner) != corner {
			count += 1
		}
	}
	return (count + 3) / 4
}
