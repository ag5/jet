//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/ag5/jet/v2/postgres"
)

var Category = newCategoryTable("dvds", "category", "")

type categoryTable struct {
	postgres.Table

	//Columns
	CategoryID postgres.ColumnInteger
	Name       postgres.ColumnString
	LastUpdate postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CategoryTable struct {
	categoryTable

	EXCLUDED categoryTable
}

// AS creates new CategoryTable with assigned alias
func (a CategoryTable) AS(alias string) *CategoryTable {
	return newCategoryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CategoryTable with assigned schema name
func (a CategoryTable) FromSchema(schemaName string) *CategoryTable {
	return newCategoryTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CategoryTable with assigned table prefix
func (a CategoryTable) WithPrefix(prefix string) *CategoryTable {
	return newCategoryTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CategoryTable with assigned table suffix
func (a CategoryTable) WithSuffix(suffix string) *CategoryTable {
	return newCategoryTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCategoryTable(schemaName, tableName, alias string) *CategoryTable {
	return &CategoryTable{
		categoryTable: newCategoryTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newCategoryTableImpl("", "excluded", ""),
	}
}

func newCategoryTableImpl(schemaName, tableName, alias string) categoryTable {
	var (
		CategoryIDColumn = postgres.IntegerColumn("category_id")
		NameColumn       = postgres.StringColumn("name")
		LastUpdateColumn = postgres.TimestampColumn("last_update")
		allColumns       = postgres.ColumnList{CategoryIDColumn, NameColumn, LastUpdateColumn}
		mutableColumns   = postgres.ColumnList{NameColumn, LastUpdateColumn}
	)

	return categoryTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CategoryID: CategoryIDColumn,
		Name:       NameColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
