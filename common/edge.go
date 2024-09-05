package common

type EdgeIndex uint8
type EdgeOrientation bool

const (
	UpLeft     EdgeIndex = 1
	DownLeft   EdgeIndex = 2
	UpRight    EdgeIndex = 3
	DownRight  EdgeIndex = 4
	LeftBack   EdgeIndex = 5
	LeftFront  EdgeIndex = 6
	RightFront EdgeIndex = 7
	RightBack  EdgeIndex = 8
	UpFront    EdgeIndex = 9
	DownFront  EdgeIndex = 10
	DownBack   EdgeIndex = 11
	UpBack     EdgeIndex = 12

	// Short form
	UL = UpLeft
	DL = DownLeft
	UR = UpRight
	DR = DownRight
	LB = LeftBack
	LF = LeftFront
	RF = RightFront
	RB = RightBack
	UF = UpFront
	DF = DownFront
	DB = DownBack
	UB = UpBack

	// For iteration purposes
	FirstEdgeIndex EdgeIndex = 1
	LastEdgeIndex  EdgeIndex = 12
	EdgeIndexCount int       = 12
)

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
