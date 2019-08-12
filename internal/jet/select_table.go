package jet

import "errors"

// SelectTable is interface for SELECT sub-queries
type SelectTable interface {
	Alias() string
	AllColumns() ProjectionList
}

type SelectTableImpl2 struct {
	selectStmt StatementWithProjections
	alias      string

	projections []Projection
}

func NewSelectTable(selectStmt StatementWithProjections, alias string) SelectTableImpl2 {
	selectTable := SelectTableImpl2{selectStmt: selectStmt, alias: alias}

	for _, projection := range selectStmt.projections() {
		newProjection := projection.fromImpl(&selectTable)

		selectTable.projections = append(selectTable.projections, newProjection)
	}

	return selectTable
}

func (s *SelectTableImpl2) Alias() string {
	return s.alias
}

func (s *SelectTableImpl2) AllColumns() ProjectionList {
	return s.projections
}

func (s *SelectTableImpl2) serialize(statement StatementType, out *SqlBuilder, options ...SerializeOption) error {
	if s == nil {
		return errors.New("jet: Expression table is nil. ")
	}

	err := s.selectStmt.serialize(statement, out)

	if err != nil {
		return err
	}

	out.WriteString("AS")
	out.WriteIdentifier(s.alias)

	return nil
}
