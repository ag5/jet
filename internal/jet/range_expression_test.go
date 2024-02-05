package jet

import "testing"

func TestRangeExpressionEQ(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.EQ(table2ColRange), "(table1.col_range = table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.EQ(NumRange(&low, &high, "[]")), "(table1.col_range = ([$1,$2]))", "1", "4")
}

func TestRangeExpressionNOT_EQ(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.NOT_EQ(table2ColRange), "(table1.col_range != table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.NOT_EQ(NumRange(&low, &high, "[]")), "(table1.col_range != ([$1,$2]))", "1", "4")
}

func TestRangeExpressionLT(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.LT(table2ColRange), "(table1.col_range < table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.LT(NumRange(&low, &high, "[]")), "(table1.col_range < ([$1,$2]))", "1", "4")
}

func TestRangeExpressionLT_EQ(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.LT_EQ(table2ColRange), "(table1.col_range <= table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.LT_EQ(NumRange(&low, &high, "[]")), "(table1.col_range <= ([$1,$2]))", "1", "4")
}

func TestRangeExpressionGT(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.GT(table2ColRange), "(table1.col_range > table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.GT(NumRange(&low, &high, "[]")), "(table1.col_range > ([$1,$2]))", "1", "4")
}

func TestRangeExpressionGT_EQ(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.GT_EQ(table2ColRange), "(table1.col_range >= table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.GT_EQ(NumRange(&low, &high, "[]")), "(table1.col_range >= ([$1,$2]))", "1", "4")
}

func TestRangeExpressionCONTAINS(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.CONTAINS(table2ColRange), "(table1.col_range @> table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.CONTAINS(NumRange(&low, &high, "[]")), "(table1.col_range @> ([$1,$2]))", "1", "4")
}

func TestRangeExpressionOVERLAP(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.OVERLAP(table2ColRange), "(table1.col_range && table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.OVERLAP(NumRange(&low, &high, "[]")), "(table1.col_range && ([$1,$2]))", "1", "4")
}

func TestRangeExpressionUNION(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.UNION(table2ColRange), "(table1.col_range + table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.UNION(NumRange(&low, &high, "[]")), "(table1.col_range + ([$1,$2]))", "1", "4")
}

func TestRangeExpressionINTERSECTION(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.INTERSECTION(table2ColRange), "(table1.col_range * table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.INTERSECTION(NumRange(&low, &high, "[]")), "(table1.col_range * ([$1,$2]))", "1", "4")
}

func TestRangeExpressionDIFFERENCE(t *testing.T) {
	assertClauseSerialize(t, table1ColRange.DIFFERENCE(table2ColRange), "(table1.col_range - table2.col_range)")
	low := 1
	high := 4
	assertClauseSerialize(t, table1ColRange.DIFFERENCE(NumRange(&low, &high, "[]")), "(table1.col_range - ([$1,$2]))", "1", "4")
}
