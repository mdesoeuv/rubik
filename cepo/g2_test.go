package cepo

import "testing"

func TestMakeG2CornerPermutationTable(t *testing.T) {
	g3Table := MakeG3CornerPermutationTable()
	g2Table := MakeG2CornerPermutationTable(g3Table)

	t.Logf("len(g2Table) = %v", len(g2Table))
}

func TestMakeG2EdgePermutationTable(t *testing.T) {
	g2Table := MakeG2EdgePermutationTable()

	t.Logf("len(g2Table) = %v", len(g2Table))
}
