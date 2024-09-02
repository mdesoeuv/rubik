package main

type CornerIndex = uint8
type EdgeIndex = uint8
type EdgeOrientation = bool

// TODO: Make start at 0
const (
	CornerIndexUpLeftBack     CornerIndex = 1
	CornerIndexDownLeftFront  CornerIndex = 2
	CornerIndexDownRightBack  CornerIndex = 3
	CornerIndexUpRightFront   CornerIndex = 4
	CornerIndexUpLeftFront    CornerIndex = 5
	CornerIndexDownLeftBack   CornerIndex = 6
	CornerIndexDownRightFront CornerIndex = 7
	CornerIndexUpRightBack    CornerIndex = 8
	// For iteration purposes
	CornerIndexFirst CornerIndex = 1
	CornerIndexLast  CornerIndex = 8
)

// TODO: Make start at 0
const (
	EdgeIndexUpLeft     EdgeIndex = 1
	EdgeIndexDownLeft   EdgeIndex = 2
	EdgeIndexUpRight    EdgeIndex = 3
	EdgeIndexDownRight  EdgeIndex = 4
	EdgeIndexLeftBack   EdgeIndex = 5
	EdgeIndexLeftFront  EdgeIndex = 6
	EdgeIndexRightFront EdgeIndex = 7
	EdgeIndexRightBack  EdgeIndex = 8
	EdgeIndexUpFront    EdgeIndex = 9
	EdgeIndexDownFront  EdgeIndex = 10
	EdgeIndexDownBack   EdgeIndex = 11
	EdgeIndexUpBack     EdgeIndex = 12
	// For iteration purposes
	FirstEdgeIndex CornerIndex = 1
	LastEdgeIndex  CornerIndex = 12
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

var (
	CrownEdgeUp    = [4]EdgeIndex{12, 3, 9, 1}
	CrownEdgeDown  = [4]EdgeIndex{10, 4, 11, 2}
	CrownEdgeLeft  = [4]EdgeIndex{1, 6, 2, 5}
	CrownEdgeRight = [4]EdgeIndex{3, 8, 4, 7}
	CrownEdgeFront = [4]EdgeIndex{9, 7, 10, 6}
	CrownEdgeBack  = [4]EdgeIndex{12, 5, 11, 8}
)

var CrownEdgeList = [SideCount][4]EdgeIndex{
	CrownEdgeUp,
	CrownEdgeDown,
	CrownEdgeLeft,
	CrownEdgeRight,
	CrownEdgeFront,
	CrownEdgeBack,
}

type EdgeOrientations struct {
	bits EdgeOrientationBits
}

func NewEdgeOrientationsSolved() EdgeOrientations {
	return EdgeOrientations{bits: 0}
}

func (eo EdgeOrientations) isSolved() bool {
	return eo.bits == 0
}

func (eo EdgeOrientations) get(index EdgeIndex) EdgeOrientation {
	var i = index - 1
	return (eo.bits>>i)&1 == 1
}

func (eo *EdgeOrientations) set(index EdgeIndex, orientation EdgeOrientation) {
	var i = index - 1
	eo.bits &= ^(1 << i) // Clear bitb
	n := EdgeOrientationBits(0)
	if orientation {
		n = EdgeOrientationBits(1)
	}
	eo.bits |= (n << i) // Set value
}

func (eo *EdgeOrientations) apply(move Move) {
	neo := *eo

	crown := &CrownEdgeList[move.Side]
	for i := range crown {
		from, to := crown[i], crown[(i+4+move.NumRotations)%len(crown)]
		neo.set(EdgeIndex(to), eo.get(from))
	}
	*eo = neo

	if move.NumRotations%2 == 0 {
		return
	}
	if move.Side == Up {
		eo.bits ^= EdgeOrientationMaskUp
	} else if move.Side == Down {
		eo.bits ^= EdgeOrientationMaskDown
	}
}

type EdgePermutationBits = uint64
type EdgePermutation struct {
	bits EdgePermutationBits
}

func NewEdgePermutationSolved() EdgePermutation {
	ep := EdgePermutation{bits: 0}
	for index := FirstEdgeIndex; index <= LastEdgeIndex; index++ {
		ep.set(index, index)
	}
	return ep
}

func (ep EdgePermutation) isSolved() bool {
	for index := FirstEdgeIndex; index <= LastEdgeIndex; index++ {
		if ep.get(index) != index {
			return false
		}
	}
	return true
}

const (
	EdgePermutationMaskUp    EdgePermutationBits = 0b111_000_000_111_000_000_000_000_000_111_000_111
	EdgePermutationMaskDown  EdgePermutationBits = 0b000_111_111_000_000_000_000_000_111_000_111_000
	EdgePermutationMaskLeft  EdgePermutationBits = 0b000_000_000_000_000_000_111_111_000_000_111_111
	EdgePermutationMaskRight EdgePermutationBits = 0b000_000_000_000_111_111_000_000_111_111_000_000
	EdgePermutationMaskFront EdgePermutationBits = 0b000_000_111_111_000_111_001_000_000_000_000_000
	EdgePermutationMaskBack  EdgePermutationBits = 0b111_111_000_000_111_000_000_111_000_000_000_000
)

func (ep EdgePermutation) get(index EdgeIndex) EdgeIndex {
	var i = index - 1
	return EdgeIndex((ep.bits>>(i*4))&0xf + 1)
}

func (ep *EdgePermutation) set(index, value EdgeIndex) {
	var i = index - 1
	ep.bits &= ^(0xf << (i * 4))                       // Clear edge
	ep.bits |= EdgePermutationBits(value-1) << (i * 4) // Set value
}

func (ep *EdgePermutation) apply(move Move) {
	nep := *ep

	crown := &CrownEdgeList[move.Side]
	for i := range crown {
		from, to := crown[i], crown[(i+4+move.NumRotations)%len(crown)]
		nep.set(EdgeIndex(to), ep.get(from))
	}
	*ep = nep
}

type CornerOrientations struct {
	bits uint32
}

func NewCornerOrientationsSolved() CornerOrientations {
	return CornerOrientations{bits: 0}
}

func (co CornerOrientations) isSolved() bool {
	return co.bits == 0
}

var (
	CrownCornerUp    = [4]CornerIndex{1, 8, 4, 5}
	CrownCornerDown  = [4]CornerIndex{2, 7, 3, 6}
	CrownCornerLeft  = [4]CornerIndex{1, 5, 2, 6}
	CrownCornerRight = [4]CornerIndex{4, 8, 3, 7}
	CrownCornerFront = [4]CornerIndex{5, 4, 7, 2}
	CrownCornerBack  = [4]CornerIndex{8, 1, 6, 3}
)

var CrownCornerList = [SideCount][4]CornerIndex{
	CrownCornerUp,
	CrownCornerDown,
	CrownCornerLeft,
	CrownCornerRight,
	CrownCornerFront,
	CrownCornerBack,
}

type CornerPermutationBits = uint32
type CornerPermutation struct {
	bits CornerPermutationBits
}

func NewCornerPermutationSolved() CornerPermutation {
	cp := CornerPermutation{bits: 0}
	for index := CornerIndexFirst; index <= CornerIndexLast; index++ {
		cp.set(index, index)
	}
	return cp
}

func (cp CornerPermutation) isSolved() bool {
	for index := CornerIndexFirst; index <= CornerIndexLast; index++ {
		if cp.get(index) != index {
			return false
		}
	}
	return true
}

type CornerOrientation = uint8

const (
	CornerOrientationLeftRight CornerOrientation = 0
	CornerOrientationUpDown    CornerOrientation = 1
	CornerOrientationFrontBack CornerOrientation = 2
	CornerOrientationCount     CornerOrientation = 3
)

func (co CornerOrientations) get(index CornerIndex) CornerOrientation {
	var i = index - 1
	return CornerOrientation((co.bits >> (i * 2)) & 3)
}

func (co *CornerOrientations) set(index CornerIndex, orientation CornerOrientation) {
	var i = index - 1
	co.bits &= ^(3 << (i * 2))                  // Clear bits
	co.bits |= (uint32(orientation) << (i * 2)) // Set value
}

func (co *CornerOrientations) apply(move Move) {
	nco := *co
	crown := &CrownCornerList[move.Side]
	rotation := ((move.Side >> 1) + 2) << 1
	for i := range crown {
		from, to := crown[i], crown[(i+4+move.NumRotations)%len(crown)]
		orientation := co.get(from)
		// TODO: Remove condition
		if move.NumRotations != 2 {
			orientation = ((orientation << 1) + rotation<<1) % CornerOrientationCount
		}
		nco.set(CornerIndex(to), orientation)
	}
	*co = nco
}

func (cp CornerPermutation) get(index CornerIndex) CornerIndex {
	var i = index - 1
	return CornerIndex((cp.bits>>(i*3))&0x7 + 1)
}

func (cp *CornerPermutation) set(index, value CornerIndex) {
	var i = index - 1
	cp.bits &= ^CornerPermutationBits(0x7 << (i * 3))    // Clear edge
	cp.bits |= CornerPermutationBits(value-1) << (i * 3) // Set value
}

func (cp *CornerPermutation) apply(move Move) {
	ncp := *cp
	crown := &CrownCornerList[move.Side]
	for i := range crown {
		from, to := crown[i], crown[(i+4+move.NumRotations)%len(crown)]
		ncp.set(EdgeIndex(to), cp.get(from))
	}
	*cp = ncp
}

type CubeV2 struct {
	// TODO: Conbine edge description in a single uint64
	edgeOrientations   EdgeOrientations
	edgePermutation    EdgePermutation
	cornerOrientations CornerOrientations
	cornerPermutation  CornerPermutation
}

func NewCubeV2Solved() CubeV2 {
	return CubeV2{
		edgeOrientations:   NewEdgeOrientationsSolved(),
		edgePermutation:    NewEdgePermutationSolved(),
		cornerOrientations: NewCornerOrientationsSolved(),
		cornerPermutation:  NewCornerPermutationSolved(),
	}
}

func (c *CubeV2) apply(m Move) {
	c.edgeOrientations.apply(m)
	c.edgePermutation.apply(m)
	c.cornerOrientations.apply(m)
	c.cornerPermutation.apply(m)
}

type EdgeFace struct {
	index   EdgeIndex
	faceNmb uint8
}

// TODO: Find formula
var EdgeFaceMap = map[CubeCoord]EdgeFace{
	{Up, FaceCoord{1, 0}}:    {EdgeIndexUpLeft, 0},
	{Left, FaceCoord{0, 1}}:  {EdgeIndexUpLeft, 1},
	{Down, FaceCoord{1, 0}}:  {EdgeIndexDownLeft, 0},
	{Left, FaceCoord{2, 1}}:  {EdgeIndexDownLeft, 1},
	{Up, FaceCoord{1, 2}}:    {EdgeIndexUpRight, 0},
	{Right, FaceCoord{0, 1}}: {EdgeIndexUpRight, 1},
	{Down, FaceCoord{1, 2}}:  {EdgeIndexDownRight, 0},
	{Right, FaceCoord{2, 1}}: {EdgeIndexDownRight, 1},
	{Left, FaceCoord{1, 0}}:  {EdgeIndexLeftBack, 0},
	{Back, FaceCoord{1, 2}}:  {EdgeIndexLeftBack, 1},
	{Left, FaceCoord{1, 2}}:  {EdgeIndexLeftFront, 0},
	{Front, FaceCoord{1, 0}}: {EdgeIndexLeftFront, 1},
	{Right, FaceCoord{1, 0}}: {EdgeIndexRightFront, 0},
	{Front, FaceCoord{1, 2}}: {EdgeIndexRightFront, 1},
	{Right, FaceCoord{1, 2}}: {EdgeIndexRightBack, 0},
	{Back, FaceCoord{1, 0}}:  {EdgeIndexRightBack, 1},
	{Up, FaceCoord{2, 1}}:    {EdgeIndexUpFront, 0},
	{Front, FaceCoord{0, 1}}: {EdgeIndexUpFront, 1},
	{Down, FaceCoord{0, 1}}:  {EdgeIndexDownFront, 0},
	{Front, FaceCoord{2, 1}}: {EdgeIndexDownFront, 1},
	{Down, FaceCoord{2, 1}}:  {EdgeIndexDownBack, 0},
	{Back, FaceCoord{2, 1}}:  {EdgeIndexDownBack, 1},
	{Up, FaceCoord{0, 1}}:    {EdgeIndexUpBack, 0},
	{Back, FaceCoord{0, 1}}:  {EdgeIndexUpBack, 1},
}

// TODO: Find formula
var CornerCoordMap = map[CubeCoord]CornerIndex{
	{Up, FaceCoord{0, 0}}:    CornerIndexUpLeftBack,
	{Left, FaceCoord{0, 0}}:  CornerIndexUpLeftBack,
	{Back, FaceCoord{0, 2}}:  CornerIndexUpLeftBack,
	{Down, FaceCoord{2, 0}}:  CornerIndexDownLeftFront,
	{Left, FaceCoord{2, 0}}:  CornerIndexDownLeftFront,
	{Back, FaceCoord{2, 2}}:  CornerIndexDownLeftFront,
	{Down, FaceCoord{2, 2}}:  CornerIndexDownRightBack,
	{Right, FaceCoord{2, 2}}: CornerIndexDownRightBack,
	{Back, FaceCoord{2, 0}}:  CornerIndexDownRightBack,
	{Up, FaceCoord{2, 2}}:    CornerIndexUpRightFront,
	{Right, FaceCoord{0, 0}}: CornerIndexUpRightFront,
	{Front, FaceCoord{0, 2}}: CornerIndexUpRightFront,
	{Up, FaceCoord{2, 0}}:    CornerIndexUpLeftFront,
	{Left, FaceCoord{0, 2}}:  CornerIndexUpLeftFront,
	{Front, FaceCoord{0, 0}}: CornerIndexUpLeftFront,
	{Down, FaceCoord{2, 0}}:  CornerIndexDownLeftBack,
	{Left, FaceCoord{2, 0}}:  CornerIndexDownLeftBack,
	{Back, FaceCoord{2, 2}}:  CornerIndexDownLeftBack,
	{Down, FaceCoord{0, 2}}:  CornerIndexDownRightFront,
	{Right, FaceCoord{2, 0}}: CornerIndexDownRightFront,
	{Front, FaceCoord{2, 2}}: CornerIndexDownRightFront,
	{Up, FaceCoord{0, 2}}:    CornerIndexUpRightBack,
	{Right, FaceCoord{0, 2}}: CornerIndexUpRightBack,
	{Back, FaceCoord{0, 0}}:  CornerIndexUpRightBack,
}

var EdgeIndexMap = map[EdgeIndex]EdgeCoords{
	EdgeIndexUpLeft: {
		a: CubeCoord{Up, FaceCoord{1, 0}},
		b: CubeCoord{Left, FaceCoord{0, 1}},
	},
	EdgeIndexDownLeft: {
		a: CubeCoord{Down, FaceCoord{1, 0}},
		b: CubeCoord{Left, FaceCoord{2, 1}},
	},
	EdgeIndexUpRight: {
		a: CubeCoord{Up, FaceCoord{1, 2}},
		b: CubeCoord{Right, FaceCoord{0, 1}},
	},
	EdgeIndexDownRight: {
		a: CubeCoord{Down, FaceCoord{1, 2}},
		b: CubeCoord{Right, FaceCoord{2, 1}},
	},
	EdgeIndexLeftBack: {
		a: CubeCoord{Left, FaceCoord{1, 0}},
		b: CubeCoord{Back, FaceCoord{1, 2}},
	},
	EdgeIndexLeftFront: {
		a: CubeCoord{Left, FaceCoord{1, 2}},
		b: CubeCoord{Front, FaceCoord{1, 0}},
	},
	EdgeIndexRightFront: {
		a: CubeCoord{Right, FaceCoord{1, 0}},
		b: CubeCoord{Front, FaceCoord{1, 2}},
	},
	EdgeIndexRightBack: {
		a: CubeCoord{Right, FaceCoord{1, 2}},
		b: CubeCoord{Back, FaceCoord{1, 0}},
	},
	EdgeIndexUpFront: {
		a: CubeCoord{Up, FaceCoord{2, 1}},
		b: CubeCoord{Front, FaceCoord{0, 1}},
	},
	EdgeIndexDownFront: {
		a: CubeCoord{Down, FaceCoord{0, 1}},
		b: CubeCoord{Front, FaceCoord{2, 1}},
	},
	EdgeIndexDownBack: {
		a: CubeCoord{Down, FaceCoord{2, 1}},
		b: CubeCoord{Back, FaceCoord{2, 1}},
	},
	EdgeIndexUpBack: {
		a: CubeCoord{Up, FaceCoord{0, 1}},
		b: CubeCoord{Back, FaceCoord{0, 1}},
	},
}

var CornerIndexMap = map[CornerIndex]CornerCoords{
	CornerIndexUpLeftBack: {
		a: CubeCoord{Up, FaceCoord{0, 0}},
		b: CubeCoord{Left, FaceCoord{0, 0}},
		c: CubeCoord{Back, FaceCoord{0, 2}},
	},
	CornerIndexDownLeftFront: {
		a: CubeCoord{Down, FaceCoord{2, 0}},
		b: CubeCoord{Left, FaceCoord{2, 0}},
		c: CubeCoord{Back, FaceCoord{2, 2}},
	},
	CornerIndexDownRightBack: {
		a: CubeCoord{Down, FaceCoord{2, 2}},
		b: CubeCoord{Right, FaceCoord{2, 2}},
		c: CubeCoord{Back, FaceCoord{2, 0}},
	},
	CornerIndexUpRightFront: {
		a: CubeCoord{Up, FaceCoord{2, 2}},
		b: CubeCoord{Right, FaceCoord{0, 0}},
		c: CubeCoord{Front, FaceCoord{0, 2}},
	},
	CornerIndexUpLeftFront: {
		a: CubeCoord{Up, FaceCoord{2, 0}},
		b: CubeCoord{Left, FaceCoord{0, 2}},
		c: CubeCoord{Front, FaceCoord{0, 0}},
	},
	CornerIndexDownLeftBack: {
		a: CubeCoord{Down, FaceCoord{2, 0}},
		b: CubeCoord{Left, FaceCoord{2, 0}},
		c: CubeCoord{Back, FaceCoord{2, 2}},
	},
	CornerIndexDownRightFront: {
		a: CubeCoord{Down, FaceCoord{0, 2}},
		b: CubeCoord{Right, FaceCoord{2, 0}},
		c: CubeCoord{Front, FaceCoord{2, 2}},
	},
	CornerIndexUpRightBack: {
		a: CubeCoord{Up, FaceCoord{0, 2}},
		b: CubeCoord{Right, FaceCoord{0, 2}},
		c: CubeCoord{Back, FaceCoord{0, 0}},
	},
}

func (c CubeV2) get(coord CubeCoord) Side {
	if coord.faceCoord.isEdge() {
		edgeFace := EdgeFaceMap[coord]
		originalIndex := c.edgePermutation.get(edgeFace.index)
		originalCoords := EdgeIndexMap[originalIndex]
		turned := c.edgeOrientations.get(originalIndex)
		if (edgeFace.faceNmb == 0) != turned {
			return originalCoords.a.side
		} else {
			return originalCoords.b.side
		}
	} else {
		// cornerIndex := CornerCoordMap[coord]
		// originalIndex := c.cornerPermutation.get(cornerIndex)
		// originalCoords := CornerIndexMap[originalIndex]
		// TODO: Implement
		panic("TODO")
	}
}

func (c CubeV2) isSolved() bool {
	return (c.edgeOrientations.isSolved() &&
		c.edgePermutation.isSolved() &&
		c.cornerOrientations.isSolved() &&
		c.cornerPermutation.isSolved())
}
