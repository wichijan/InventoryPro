//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var Points = newPointsTable("InventoryProDB", "points", "")

type pointsTable struct {
	mysql.Table

	// Columns
	UserID mysql.ColumnString
	Points mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type PointsTable struct {
	pointsTable

	NEW pointsTable
}

// AS creates new PointsTable with assigned alias
func (a PointsTable) AS(alias string) *PointsTable {
	return newPointsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PointsTable with assigned schema name
func (a PointsTable) FromSchema(schemaName string) *PointsTable {
	return newPointsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PointsTable with assigned table prefix
func (a PointsTable) WithPrefix(prefix string) *PointsTable {
	return newPointsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PointsTable with assigned table suffix
func (a PointsTable) WithSuffix(suffix string) *PointsTable {
	return newPointsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPointsTable(schemaName, tableName, alias string) *PointsTable {
	return &PointsTable{
		pointsTable: newPointsTableImpl(schemaName, tableName, alias),
		NEW:         newPointsTableImpl("", "new", ""),
	}
}

func newPointsTableImpl(schemaName, tableName, alias string) pointsTable {
	var (
		UserIDColumn   = mysql.StringColumn("user_id")
		PointsColumn   = mysql.IntegerColumn("points")
		allColumns     = mysql.ColumnList{UserIDColumn, PointsColumn}
		mutableColumns = mysql.ColumnList{PointsColumn}
	)

	return pointsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID: UserIDColumn,
		Points: PointsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
