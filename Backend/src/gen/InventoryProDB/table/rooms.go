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

var Rooms = newRoomsTable("InventoryProDB", "rooms", "")

type roomsTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnString
	Name        mysql.ColumnString
	WarehouseID mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type RoomsTable struct {
	roomsTable

	NEW roomsTable
}

// AS creates new RoomsTable with assigned alias
func (a RoomsTable) AS(alias string) *RoomsTable {
	return newRoomsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RoomsTable with assigned schema name
func (a RoomsTable) FromSchema(schemaName string) *RoomsTable {
	return newRoomsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RoomsTable with assigned table prefix
func (a RoomsTable) WithPrefix(prefix string) *RoomsTable {
	return newRoomsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RoomsTable with assigned table suffix
func (a RoomsTable) WithSuffix(suffix string) *RoomsTable {
	return newRoomsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRoomsTable(schemaName, tableName, alias string) *RoomsTable {
	return &RoomsTable{
		roomsTable: newRoomsTableImpl(schemaName, tableName, alias),
		NEW:        newRoomsTableImpl("", "new", ""),
	}
}

func newRoomsTableImpl(schemaName, tableName, alias string) roomsTable {
	var (
		IDColumn          = mysql.StringColumn("id")
		NameColumn        = mysql.StringColumn("name")
		WarehouseIDColumn = mysql.StringColumn("warehouse_id")
		allColumns        = mysql.ColumnList{IDColumn, NameColumn, WarehouseIDColumn}
		mutableColumns    = mysql.ColumnList{NameColumn, WarehouseIDColumn}
	)

	return roomsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		WarehouseID: WarehouseIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
