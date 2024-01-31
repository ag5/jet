package jet

// RangeExpression is interface for date range types
type RangeExpression interface {
	Expression

	EQ(rhs RangeExpression) BoolExpression
	NOT_EQ(rhs RangeExpression) BoolExpression

	LT(rhs RangeExpression) BoolExpression
	LT_EQ(rhs RangeExpression) BoolExpression
	GT(rhs RangeExpression) BoolExpression
	GT_EQ(rhs RangeExpression) BoolExpression

	CONTAINS(rhs Expression) BoolExpression
	OVERLAP(rhs RangeExpression) BoolExpression
	UNION(rhs RangeExpression) RangeExpression
	INTERSECTION(rhs RangeExpression) RangeExpression
	DIFFERENCE(rhs RangeExpression) RangeExpression
}

type rangeInterfaceImpl struct {
	parent RangeExpression
}

func (r *rangeInterfaceImpl) EQ(rhs RangeExpression) BoolExpression {
	return Eq(r.parent, rhs)
}

func (r *rangeInterfaceImpl) NOT_EQ(rhs RangeExpression) BoolExpression {
	return NotEq(r.parent, rhs)
}

func (r *rangeInterfaceImpl) LT(rhs RangeExpression) BoolExpression {
	return Lt(r.parent, rhs)
}

func (r *rangeInterfaceImpl) LT_EQ(rhs RangeExpression) BoolExpression {
	return LtEq(r.parent, rhs)

}

func (r *rangeInterfaceImpl) GT(rhs RangeExpression) BoolExpression {
	return Gt(r.parent, rhs)

}

func (r *rangeInterfaceImpl) GT_EQ(rhs RangeExpression) BoolExpression {
	return GtEq(r.parent, rhs)
}

func (r *rangeInterfaceImpl) CONTAINS(rhs Expression) BoolExpression {
	return Contains(r.parent, rhs)
}

func (r *rangeInterfaceImpl) OVERLAP(rhs RangeExpression) BoolExpression {
	return Overlap(r.parent, rhs)
}

func (r *rangeInterfaceImpl) UNION(rhs RangeExpression) RangeExpression {
	return RangeExp(Add(r.parent, rhs))
}

func (r *rangeInterfaceImpl) INTERSECTION(rhs RangeExpression) RangeExpression {
	return RangeExp(Mul(r.parent, rhs))
}

func (r *rangeInterfaceImpl) DIFFERENCE(rhs RangeExpression) RangeExpression {
	return RangeExp(Sub(r.parent, rhs))
}

//---------------------------------------------------//

type rangeExpressionWrapper struct {
	rangeInterfaceImpl
	Expression
}

func newRangeExpressionWrap(expression Expression) RangeExpression {
	rangeExpressionWrap := rangeExpressionWrapper{Expression: expression}
	rangeExpressionWrap.rangeInterfaceImpl.parent = &rangeExpressionWrap
	return &rangeExpressionWrap
}

// RangeExp is range expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as range expression.
// Does not add sql cast to generated sql builder output.
func RangeExp(expression Expression) RangeExpression {
	return newRangeExpressionWrap(expression)
}
