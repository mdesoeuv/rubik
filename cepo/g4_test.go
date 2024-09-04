package cepo

import (
	"math"
	"testing"
)

func TestG4CombinedHeuristic(t *testing.T) {
	cornerHeuristics := MakeG3CornerPermutationTable()
	edgeHeuristics := MakeG3EdgePermutationTable()
	heuristics := MakeG3BetterHeuristicTable()

	meanDifference := 0
	maxDifference := 0
	squareMean := 0
	for cube, distance := range heuristics {
		ch := cornerHeuristics[cube.corner]
		eh := edgeHeuristics[cube.edge]
		var h uint8 = eh
		if ch > eh {
			h = ch
		}

		difference := int(distance) - int(h)
		squareMean += difference * difference
		meanDifference += difference
		if difference > maxDifference {
			maxDifference = difference
		}
	}
	meanDifference /= len(heuristics)
	squareMean /= len(heuristics)
	stddev := float64(squareMean - (meanDifference * meanDifference))
	stddev = math.Sqrt(stddev)

	t.Logf("Difference (mean %v, max %v, stddev %v)", meanDifference, maxDifference, stddev)
}

func TestG4BasicHeuristic(t *testing.T) {
	heuristics := MakeG3BetterHeuristicTable()

	squareMean := 0
	meanDifference := 0
	maxDifference := 0
	for cube, distance := range heuristics {
		epDistance := cube.edge.Distance()
		cpDistance := cube.corner.Distance()

		var h = cpDistance
		if epDistance > cpDistance {
			h = epDistance
		}
		difference := int(distance) - int(h)
		if difference > maxDifference {
			maxDifference = difference
		}
		meanDifference += difference
		squareMean += difference * difference
	}
	meanDifference /= len(heuristics)
	squareMean /= len(heuristics)
	stddev := float64(squareMean - (meanDifference * meanDifference))
	stddev = math.Sqrt(stddev)

	t.Logf("Difference (mean %v, max %v, stddev %v)", meanDifference, maxDifference, stddev)
}
