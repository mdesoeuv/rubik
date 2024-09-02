package common

type EdgeFace struct {
	Index   EdgeIndex
	FaceNmb uint8
}

// TODO: Find formula
var EdgeFaceMap = map[CubeCoord]EdgeFace{
	{Side: Up, FaceCoord: FaceCoord10}:    {UpLeft, 0},
	{Side: Left, FaceCoord: FaceCoord01}:  {UpLeft, 1},
	{Side: Down, FaceCoord: FaceCoord10}:  {DownLeft, 0},
	{Side: Left, FaceCoord: FaceCoord21}:  {DownLeft, 1},
	{Side: Up, FaceCoord: FaceCoord12}:    {UpRight, 0},
	{Side: Right, FaceCoord: FaceCoord01}: {UpRight, 1},
	{Side: Down, FaceCoord: FaceCoord12}:  {DownRight, 0},
	{Side: Right, FaceCoord: FaceCoord21}: {DownRight, 1},
	{Side: Left, FaceCoord: FaceCoord10}:  {LeftBack, 0},
	{Side: Back, FaceCoord: FaceCoord12}:  {LeftBack, 1},
	{Side: Left, FaceCoord: FaceCoord12}:  {LeftFront, 0},
	{Side: Front, FaceCoord: FaceCoord10}: {LeftFront, 1},
	{Side: Right, FaceCoord: FaceCoord10}: {RightFront, 0},
	{Side: Front, FaceCoord: FaceCoord12}: {RightFront, 1},
	{Side: Right, FaceCoord: FaceCoord12}: {RightBack, 0},
	{Side: Back, FaceCoord: FaceCoord10}:  {RightBack, 1},
	{Side: Up, FaceCoord: FaceCoord21}:    {UpFront, 0},
	{Side: Front, FaceCoord: FaceCoord01}: {UpFront, 1},
	{Side: Down, FaceCoord: FaceCoord01}:  {DownFront, 0},
	{Side: Front, FaceCoord: FaceCoord21}: {DownFront, 1},
	{Side: Down, FaceCoord: FaceCoord21}:  {DownBack, 0},
	{Side: Back, FaceCoord: FaceCoord21}:  {DownBack, 1},
	{Side: Up, FaceCoord: FaceCoord01}:    {UpBack, 0},
	{Side: Back, FaceCoord: FaceCoord01}:  {UpBack, 1},
}

// TODO: Find formula
var CornerCoordMap = map[CubeCoord]CornerIndex{
	{Side: Up, FaceCoord: FaceCoord00}:    UpLeftBack,
	{Side: Left, FaceCoord: FaceCoord00}:  UpLeftBack,
	{Side: Back, FaceCoord: FaceCoord02}:  UpLeftBack,
	{Side: Down, FaceCoord: FaceCoord20}:  DownLeftFront,
	{Side: Left, FaceCoord: FaceCoord20}:  DownLeftFront,
	{Side: Back, FaceCoord: FaceCoord22}:  DownLeftFront,
	{Side: Down, FaceCoord: FaceCoord22}:  DownRightBack,
	{Side: Right, FaceCoord: FaceCoord22}: DownRightBack,
	{Side: Back, FaceCoord: FaceCoord20}:  DownRightBack,
	{Side: Up, FaceCoord: FaceCoord22}:    UpRightFront,
	{Side: Right, FaceCoord: FaceCoord00}: UpRightFront,
	{Side: Front, FaceCoord: FaceCoord02}: UpRightFront,
	{Side: Up, FaceCoord: FaceCoord20}:    UpLeftFront,
	{Side: Left, FaceCoord: FaceCoord02}:  UpLeftFront,
	{Side: Front, FaceCoord: FaceCoord00}: UpLeftFront,
	{Side: Down, FaceCoord: FaceCoord20}:  DownLeftBack,
	{Side: Left, FaceCoord: FaceCoord20}:  DownLeftBack,
	{Side: Back, FaceCoord: FaceCoord22}:  DownLeftBack,
	{Side: Down, FaceCoord: FaceCoord02}:  DownRightFront,
	{Side: Right, FaceCoord: FaceCoord20}: DownRightFront,
	{Side: Front, FaceCoord: FaceCoord22}: DownRightFront,
	{Side: Up, FaceCoord: FaceCoord02}:    UpRightBack,
	{Side: Right, FaceCoord: FaceCoord02}: UpRightBack,
	{Side: Back, FaceCoord: FaceCoord00}:  UpRightBack,
}

var EdgeIndexMap = map[EdgeIndex]EdgeCoords{
	UpLeft: {
		A: CubeCoord{Side: Up, FaceCoord: FaceCoord10},
		B: CubeCoord{Side: Left, FaceCoord: FaceCoord01},
	},
	DownLeft: {
		A: CubeCoord{Side: Down, FaceCoord: FaceCoord10},
		B: CubeCoord{Side: Left, FaceCoord: FaceCoord21},
	},
	UpRight: {
		A: CubeCoord{Side: Up, FaceCoord: FaceCoord12},
		B: CubeCoord{Side: Right, FaceCoord: FaceCoord01},
	},
	DownRight: {
		A: CubeCoord{Side: Down, FaceCoord: FaceCoord12},
		B: CubeCoord{Side: Right, FaceCoord: FaceCoord21},
	},
	LeftBack: {
		A: CubeCoord{Side: Left, FaceCoord: FaceCoord10},
		B: CubeCoord{Side: Back, FaceCoord: FaceCoord12},
	},
	LeftFront: {
		A: CubeCoord{Side: Left, FaceCoord: FaceCoord12},
		B: CubeCoord{Side: Front, FaceCoord: FaceCoord10},
	},
	RightFront: {
		A: CubeCoord{Side: Right, FaceCoord: FaceCoord10},
		B: CubeCoord{Side: Front, FaceCoord: FaceCoord12},
	},
	RightBack: {
		A: CubeCoord{Side: Right, FaceCoord: FaceCoord12},
		B: CubeCoord{Side: Back, FaceCoord: FaceCoord10},
	},
	UpFront: {
		A: CubeCoord{Side: Up, FaceCoord: FaceCoord21},
		B: CubeCoord{Side: Front, FaceCoord: FaceCoord01},
	},
	DownFront: {
		A: CubeCoord{Side: Down, FaceCoord: FaceCoord01},
		B: CubeCoord{Side: Front, FaceCoord: FaceCoord21},
	},
	DownBack: {
		A: CubeCoord{Side: Down, FaceCoord: FaceCoord21},
		B: CubeCoord{Side: Back, FaceCoord: FaceCoord21},
	},
	UpBack: {
		A: CubeCoord{Side: Up, FaceCoord: FaceCoord01},
		B: CubeCoord{Side: Back, FaceCoord: FaceCoord01},
	},
}

var CornerIndexMap = map[CornerIndex]CornerCoords{
	UpLeftBack: {
		A: CubeCoord{Side: Up, FaceCoord: FaceCoord00},
		B: CubeCoord{Side: Left, FaceCoord: FaceCoord00},
		C: CubeCoord{Side: Back, FaceCoord: FaceCoord02},
	},
	DownLeftFront: {
		A: CubeCoord{Side: Down, FaceCoord: FaceCoord20},
		B: CubeCoord{Side: Left, FaceCoord: FaceCoord20},
		C: CubeCoord{Side: Back, FaceCoord: FaceCoord22},
	},
	DownRightBack: {
		A: CubeCoord{Side: Down, FaceCoord: FaceCoord22},
		B: CubeCoord{Side: Right, FaceCoord: FaceCoord22},
		C: CubeCoord{Side: Back, FaceCoord: FaceCoord20},
	},
	UpRightFront: {
		A: CubeCoord{Side: Up, FaceCoord: FaceCoord22},
		B: CubeCoord{Side: Right, FaceCoord: FaceCoord00},
		C: CubeCoord{Side: Front, FaceCoord: FaceCoord02},
	},
	UpLeftFront: {
		A: CubeCoord{Side: Up, FaceCoord: FaceCoord20},
		B: CubeCoord{Side: Left, FaceCoord: FaceCoord02},
		C: CubeCoord{Side: Front, FaceCoord: FaceCoord00},
	},
	DownLeftBack: {
		A: CubeCoord{Side: Down, FaceCoord: FaceCoord20},
		B: CubeCoord{Side: Left, FaceCoord: FaceCoord20},
		C: CubeCoord{Side: Back, FaceCoord: FaceCoord22},
	},
	DownRightFront: {
		A: CubeCoord{Side: Down, FaceCoord: FaceCoord02},
		B: CubeCoord{Side: Right, FaceCoord: FaceCoord20},
		C: CubeCoord{Side: Front, FaceCoord: FaceCoord22},
	},
	UpRightBack: {
		A: CubeCoord{Side: Up, FaceCoord: FaceCoord02},
		B: CubeCoord{Side: Right, FaceCoord: FaceCoord02},
		C: CubeCoord{Side: Back, FaceCoord: FaceCoord00},
	},
}
