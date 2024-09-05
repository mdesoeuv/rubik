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
	Bits EdgeOrientationBits
}

func NewEdgeOrientationsSolved() EdgeOrientations {
	return EdgeOrientations{Bits: 0}
}

func (eo EdgeOrientations) IsSolved() bool {
	return eo.Bits == 0
}

func (eo EdgeOrientations) Get(index cmn.EdgeIndex) cmn.EdgeOrientation {
	var i = index - 1
	return (eo.Bits>>i)&1 == 1
}

func (eo *EdgeOrientations) Set(index cmn.EdgeIndex, orientation cmn.EdgeOrientation) {
	var i = index - 1
	eo.Bits &= ^(1 << i) // Clear bitb
	n := EdgeOrientationBits(0)
	if orientation {
		n = EdgeOrientationBits(1)
	}
	eo.Bits |= (n << i) // Set value
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
		eo.Bits ^= EdgeOrientationMaskUp
	} else if move.Side == cmn.Down {
		eo.Bits ^= EdgeOrientationMaskDown
	}
}

func (eo EdgeOrientations) Distance() int {
	const UpDownMask = EdgeOrientationMaskUp | EdgeOrientationMaskDown
	upDownDistance := bits.OnesCount16(eo.Bits & UpDownMask)
	middleDistance := 2 * bits.OnesCount16(eo.Bits & ^UpDownMask)
	return (upDownDistance + middleDistance + 3) / 4
}

func (eo EdgeOrientations) Equal(other EdgeOrientations) bool {
	return eo.Bits == other.Bits
}

type EdgePermutationBits = uint64
type EdgePermutation struct {
	Bits EdgePermutationBits
}

func NewEdgePermutationSolved() EdgePermutation {
	ep := EdgePermutation{Bits: 0}
	for index := cmn.FirstEdgeIndex; index <= cmn.LastEdgeIndex; index++ {
		ep.Set(index, index)
	}
	return ep
}

func (ep EdgePermutation) IsSolved() bool {
	return ep.Equal(NewEdgePermutationSolved())
}

// TODO: rename slices
func (ep EdgePermutation) URBLInCorrectSlice() bool {
	found := 1 << ep.Get(cmn.UL)
	found |= 1 << ep.Get(cmn.DL)
	found |= 1 << ep.Get(cmn.UR)
	found |= 1 << ep.Get(cmn.DR)
	// Extra 0 as edge start at index 1
	return found == 0b0000_0000_1111_0
}

func (ep EdgePermutation) FRBLInCorrectSlice() bool {
	found := 1 << ep.Get(cmn.LB)
	found |= 1 << ep.Get(cmn.LF)
	found |= 1 << ep.Get(cmn.RF)
	found |= 1 << ep.Get(cmn.RB)
	// Extra 0 as edge start at index 1
	return found == 0b0000_1111_0000_0
}

func (ep EdgePermutation) FUBDInCorrectSlice() bool {
	found := 1 << ep.Get(cmn.UF)
	found |= 1 << ep.Get(cmn.DF)
	found |= 1 << ep.Get(cmn.DB)
	found |= 1 << ep.Get(cmn.UB)
	// Extra 0 as edge start at index 1
	return found == 0b1111_0000_0000_0
}

func (ep EdgePermutation) URBLCorrectSliceDistance() int {
	// TODO: check/explain computation
	distance := 0
	if ep.Get(cmn.UL) > cmn.DR {
		distance += 1
	}
	if ep.Get(cmn.DL) > cmn.DR {
		distance += 1
	}
	if ep.Get(cmn.UR) > cmn.DR {
		distance += 1
	}
	if ep.Get(cmn.DR) > cmn.DR {
		distance += 1
	}
	return (distance + 1) >> 1
}

func (ep EdgePermutation) FRBLCorrectSliceDistance() int {
	// TODO: check/explain computation
	distance := 0
	if x := ep.Get(cmn.LB); x < cmn.LB || x > cmn.RB {
		distance += 1
	}
	if x := ep.Get(cmn.LF); x < cmn.LB || x > cmn.RB {
		distance += 1
	}
	if x := ep.Get(cmn.RF); x < cmn.LB || x > cmn.RB {
		distance += 1
	}
	if x := ep.Get(cmn.RB); x < cmn.LB || x > cmn.RB {
		distance += 1
	}
	return (distance + 1) >> 1
}

func (ep EdgePermutation) FUBDCorrectSliceDistance() int {
	// TODO: check/explain computation
	distance := 0
	if ep.Get(cmn.UF) < cmn.UF {
		distance += 1
	}
	if ep.Get(cmn.DF) < cmn.UF {
		distance += 1
	}
	if ep.Get(cmn.DB) < cmn.UF {
		distance += 1
	}
	if ep.Get(cmn.UB) < cmn.UF {
		distance += 1
	}
	return (distance + 1) >> 1
}

func (ep EdgePermutation) AllInCorrectSlice() bool {
	return (ep.URBLInCorrectSlice() &&
		ep.FRBLInCorrectSlice() &&
		ep.FUBDInCorrectSlice())
}

func (ep EdgePermutation) AllInCorrectSliceDistance() int {
	distance := 0
	if ep.Get(cmn.UL) > cmn.DR {
		distance += 1
	}
	if ep.Get(cmn.DL) > cmn.DR {
		distance += 1
	}
	if ep.Get(cmn.UR) > cmn.DR {
		distance += 1
	}
	if ep.Get(cmn.DR) > cmn.DR {
		distance += 1
	}
	if x := ep.Get(cmn.LB); x < cmn.LB || x > cmn.RB {
		distance += 1
	}
	if x := ep.Get(cmn.LF); x < cmn.LB || x > cmn.RB {
		distance += 1
	}
	if x := ep.Get(cmn.RF); x < cmn.LB || x > cmn.RB {
		distance += 1
	}
	if x := ep.Get(cmn.RB); x < cmn.LB || x > cmn.RB {
		distance += 1
	}
	if ep.Get(cmn.UF) < cmn.UF {
		distance += 1
	}
	if ep.Get(cmn.DF) < cmn.UF {
		distance += 1
	}
	if ep.Get(cmn.DB) < cmn.UF {
		distance += 1
	}
	if ep.Get(cmn.UB) < cmn.UF {
		distance += 1
	}
	return (distance + 3) >> 2
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
	return cmn.EdgeIndex(((ep.Bits >> (i * 4)) & 0xf) + 1)
}

func (ep *EdgePermutation) Set(index, value cmn.EdgeIndex) {
	var i = index - 1
	ep.Bits &= ^(0xf << (i * 4))                       // Clear edge
	ep.Bits |= EdgePermutationBits(value-1) << (i * 4) // Set value
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
	return ep.Bits == other.Bits
}

func (ep EdgePermutation) Distance() int {
	count := 0
	for edge := cmn.FirstEdgeIndex; edge <= cmn.LastEdgeIndex; edge++ {
		if ep.Get(edge) != edge {
			count += 1
		}
	}
	return (count + 3) / 4
}
