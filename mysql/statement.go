package mysql

import "github.com/ag5/jet/v2/internal/jet"

// RawStatement creates new sql statements from raw query and optional map of named arguments
func RawStatement(rawQuery string, namedArguments ...RawArgs) Statement {
	return jet.RawStatement(Dialect, rawQuery, namedArguments...)
}
