package jet

import "testing"

func TestRangeExpressionEQ(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.EQ(table2ColRange), "(table1.col_range = table2.col_range)")
	assertClauseSerialize(t, table1ColRange.EQ(RangeRaw("[1,4]")), "(table1.col_range = $1)", "[1,4]")
}

func TestRangeExpressionNOT_EQ(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.NOT_EQ(table2ColRange), "(table1.col_range != table2.col_range)")
	assertClauseSerialize(t, table1ColRange.NOT_EQ(RangeRaw("[1,4]")), "(table1.col_range != $1)", "[1,4]")
}

func TestRangeExpressionLT(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.LT(table2ColRange), "(table1.col_range < table2.col_range)")
	assertClauseSerialize(t, table1ColRange.LT(RangeRaw("[1,4]")), "(table1.col_range < $1)", "[1,4]")
}

func TestRangeExpressionLT_EQ(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.LT_EQ(table2ColRange), "(table1.col_range <= table2.col_range)")
	assertClauseSerialize(t, table1ColRange.LT_EQ(RangeRaw("[1,4]")), "(table1.col_range <= $1)", "[1,4]")
}

func TestRangeExpressionGT(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.GT(table2ColRange), "(table1.col_range > table2.col_range)")
	assertClauseSerialize(t, table1ColRange.GT(RangeRaw("[1,4]")), "(table1.col_range > $1)", "[1,4]")
}

func TestRangeExpressionGT_EQ(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.GT_EQ(table2ColRange), "(table1.col_range >= table2.col_range)")
	assertClauseSerialize(t, table1ColRange.GT_EQ(RangeRaw("[1,4]")), "(table1.col_range >= $1)", "[1,4]")
}

func TestRangeExpressionCONTAINS(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.CONTAINS(table2ColRange), "(table1.col_range @> table2.col_range)")
	assertClauseSerialize(t, table1ColRange.CONTAINS(RangeRaw("[1,4]")), "(table1.col_range @> $1)", "[1,4]")
}

func TestRangeExpressionOVERLAP(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.OVERLAP(table2ColRange), "(table1.col_range && table2.col_range)")
	assertClauseSerialize(t, table1ColRange.OVERLAP(RangeRaw("[1,4]")), "(table1.col_range && $1)", "[1,4]")
}

func TestRangeExpressionUNION(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.UNION(table2ColRange), "(table1.col_range + table2.col_range)")
	assertClauseSerialize(t, table1ColRange.UNION(RangeRaw("[1,4]")), "(table1.col_range + $1)", "[1,4]")
}

func TestRangeExpressionINTERSECTION(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.INTERSECTION(table2ColRange), "(table1.col_range * table2.col_range)")
	assertClauseSerialize(t, table1ColRange.INTERSECTION(RangeRaw("[1,4]")), "(table1.col_range * $1)", "[1,4]")
}

func TestRangeExpressionDIFFERENCE(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.DIFFERENCE(table2ColRange), "(table1.col_range - table2.col_range)")
	assertClauseSerialize(t, table1ColRange.DIFFERENCE(RangeRaw("[1,4]")), "(table1.col_range - $1)", "[1,4]")
}
