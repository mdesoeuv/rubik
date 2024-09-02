package cepo

import (
	"math/bits"

	cmn "github.com/mdesoeuv/rubik/common"
)

type EdgeOrientationBits = uint16

const (
	EdgeOrientationMaskUp    EdgeOrientationBits = 0b1001_0000_0101
	EdgeOrientationMaskDown  EdgeOrientationBits = 0b0110_0000_1010
	EdgeOrientationMaskLeft  EdgeOrientationBits = 0b0000_0011_0011
	EdgeOrientationMaskRight EdgeOrientationBits = 0b0000_1100_1100
	EdgeOrientationMaskFront EdgeOrientationBits = 0b0011_0110_0000
	EdgeOrientationMaskBack  EdgeOrientationBits = 0b1100_1001_0000
)

var EdgeOrientationMaskList = [6]EdgeOrientationBits{
	EdgeOrientationMaskUp,
	EdgeOrientationMaskDown,
	EdgeOrientationMaskLeft,
	EdgeOrientationMaskRight,
	EdgeOrientationMaskFront,
	EdgeOrientationMaskBack,
}

type EdgeOrientations struct {
	bits EdgeOrientationBits
}

func NewEdgeOrientationsSolved() EdgeOrientations {
	return EdgeOrientations{bits: 0}
}

func (eo EdgeOrientations) IsSolved() bool {
	return eo.bits == 0
}

func (eo EdgeOrientations) Get(index cmn.EdgeIndex) cmn.EdgeOrientation {
	var i = index - 1
	return (eo.bits>>i)&1 == 1
}

func (eo *EdgeOrientations) Set(index cmn.EdgeIndex, orientation cmn.EdgeOrientation) {
	var i = index - 1
	eo.bits &= ^(1 << i) // Clear bitb
	n := EdgeOrientationBits(0)
	if orientation {
		n = EdgeOrientationBits(1)
	}
	eo.bits |= (n << i) // Set value
}

func (eo *EdgeOrientations) Apply(move cmn.Move) {
	neo := *eo

	crown := &cmn.CrownEdgeList[move.Side]
	for i := range crown {
		from, to := crown[i], crown[(i+move.Rotation.PositiveInt())%len(crown)]
		neo.Set(cmn.EdgeIndex(to), eo.Get(from))
	}
	*eo = neo

	if !move.Rotation.IsQuaterTurn() {
		return
	}
	if move.Side == cmn.Up {
		eo.bits ^= EdgeOrientationMaskUp
	} else if move.Side == cmn.Down {
		eo.bits ^= EdgeOrientationMaskDown
	}
}

func (eo EdgeOrientations) Distance() int {
	const UpDownMask = EdgeOrientationMaskUp | EdgeOrientationMaskDown
	upDownDistance := bits.OnesCount16(eo.bits & UpDownMask)
	middleDistance := 2 * bits.OnesCount16(eo.bits & ^UpDownMask)
	return (upDownDistance + middleDistance + 3) / 4
}

func (eo EdgeOrientations) Equal(other EdgeOrientations) bool {
	return eo.bits == other.bits
}

type EdgePermutationBits = uint64
type EdgePermutation struct {
	bits EdgePermutationBits
}

func NewEdgePermutationSolved() EdgePermutation {
	ep := EdgePermutation{bits: 0}
	for index := cmn.FirstEdgeIndex; index <= cmn.LastEdgeIndex; index++ {
		ep.Set(index, index)
	}
	return ep
}

func (ep EdgePermutation) IsSolved() bool {
	for index := cmn.FirstEdgeIndex; index <= cmn.LastEdgeIndex; index++ {
		if ep.Get(index) != index {
			return false
		}
	}
	return true
}

func (ep EdgePermutation) FUBDInCorrectSlice() bool {
	found := 1 << ep.Get(cmn.UF)
	found |= 1 << ep.Get(cmn.UB)
	found |= 1 << ep.Get(cmn.DB)
	found |= 1 << ep.Get(cmn.DF)
	// Extra 0 as edge start at index 1
	return found != 0b1111_0000_0000_0
}

func (ep EdgePermutation) FUBDCorrectSliceDistance() int {
	distance := (cmn.EdgeIndexCount - int(ep.Get(cmn.UF))) >> 2
	distance += (cmn.EdgeIndexCount - int(ep.Get(cmn.UB))) >> 2
	distance += (cmn.EdgeIndexCount - int(ep.Get(cmn.DB))) >> 2
	distance += (cmn.EdgeIndexCount - int(ep.Get(cmn.DF))) >> 2
	return (distance + 1) >> 1
}

const (
	EdgePermutationMaskUp    EdgePermutationBits = 0b111_000_000_111_000_000_000_000_000_111_000_111
	EdgePermutationMaskDown  EdgePermutationBits = 0b000_111_111_000_000_000_000_000_111_000_111_000
	EdgePermutationMaskLeft  EdgePermutationBits = 0b000_000_000_000_000_000_111_111_000_000_111_111
	EdgePermutationMaskRight EdgePermutationBits = 0b000_000_000_000_111_111_000_000_111_111_000_000
	EdgePermutationMaskFront EdgePermutationBits = 0b000_000_111_111_000_111_001_000_000_000_000_000
	EdgePermutationMaskBack  EdgePermutationBits = 0b111_111_000_000_111_000_000_111_000_000_000_000
)

func (ep EdgePermutation) Get(index cmn.EdgeIndex) cmn.EdgeIndex {
	var i = index - 1
	return cmn.EdgeIndex(((ep.bits >> (i * 4)) & 0xf) + 1)
}

func (ep *EdgePermutation) Set(index, value cmn.EdgeIndex) {
	var i = index - 1
	ep.bits &= ^(0xf << (i * 4))                       // Clear edge
	ep.bits |= EdgePermutationBits(value-1) << (i * 4) // Set value
}

func (ep *EdgePermutation) Apply(move cmn.Move) {
	nep := *ep

	crown := &cmn.CrownEdgeList[move.Side]
	for i := range crown {
		from, to := crown[i], crown[(i+move.Rotation.PositiveInt())%len(crown)]
		nep.Set(cmn.EdgeIndex(to), ep.Get(from))
	}
	*ep = nep
}

func (ep EdgePermutation) Equal(other EdgePermutation) bool {
	return ep.bits == other.bits
}
