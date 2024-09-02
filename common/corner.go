package common

type CornerIndex uint8

// TODO: Make start at 0
const (
	UpLeftBack     CornerIndex = 1
	DownLeftFront  CornerIndex = 2
	DownRightBack  CornerIndex = 3
	UpRightFront   CornerIndex = 4
	UpLeftFront    CornerIndex = 5
	DownLeftBack   CornerIndex = 6
	DownRightFront CornerIndex = 7
	UpRightBack    CornerIndex = 8

	// Short form
	ULB = UpLeftBack
	DLF = DownLeftFront
	DRB = DownRightBack
	URF = UpRightFront
	ULF = UpLeftFront
	DLB = DownLeftBack
	DRF = DownRightFront
	URB = UpRightBack

	// For iteration purposes
	FirstCornerIndex CornerIndex = 1
	LastCornerIndex  CornerIndex = 8
)

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
