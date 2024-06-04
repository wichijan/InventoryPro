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

var Reservations = newReservationsTable("InventoryProDB", "reservations", "")

type reservationsTable struct {
	mysql.Table

	// Columns
	ReservationID   mysql.ColumnString
	ItemID          mysql.ColumnString
	UserID          mysql.ColumnString
	Quantity        mysql.ColumnInteger
	ReservationDate mysql.ColumnTimestamp
	TimeFrom        mysql.ColumnTimestamp
	TimeTo          mysql.ColumnTimestamp
	IsCancelled     mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ReservationsTable struct {
	reservationsTable

	NEW reservationsTable
}

// AS creates new ReservationsTable with assigned alias
func (a ReservationsTable) AS(alias string) *ReservationsTable {
	return newReservationsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ReservationsTable with assigned schema name
func (a ReservationsTable) FromSchema(schemaName string) *ReservationsTable {
	return newReservationsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ReservationsTable with assigned table prefix
func (a ReservationsTable) WithPrefix(prefix string) *ReservationsTable {
	return newReservationsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ReservationsTable with assigned table suffix
func (a ReservationsTable) WithSuffix(suffix string) *ReservationsTable {
	return newReservationsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newReservationsTable(schemaName, tableName, alias string) *ReservationsTable {
	return &ReservationsTable{
		reservationsTable: newReservationsTableImpl(schemaName, tableName, alias),
		NEW:               newReservationsTableImpl("", "new", ""),
	}
}

func newReservationsTableImpl(schemaName, tableName, alias string) reservationsTable {
	var (
		ReservationIDColumn   = mysql.StringColumn("reservation_id")
		ItemIDColumn          = mysql.StringColumn("item_id")
		UserIDColumn          = mysql.StringColumn("user_id")
		QuantityColumn        = mysql.IntegerColumn("quantity")
		ReservationDateColumn = mysql.TimestampColumn("reservation_date")
		TimeFromColumn        = mysql.TimestampColumn("time_from")
		TimeToColumn          = mysql.TimestampColumn("time_to")
		IsCancelledColumn     = mysql.BoolColumn("is_cancelled")
		allColumns            = mysql.ColumnList{ReservationIDColumn, ItemIDColumn, UserIDColumn, QuantityColumn, ReservationDateColumn, TimeFromColumn, TimeToColumn, IsCancelledColumn}
		mutableColumns        = mysql.ColumnList{ItemIDColumn, UserIDColumn, QuantityColumn, ReservationDateColumn, TimeFromColumn, TimeToColumn, IsCancelledColumn}
	)

	return reservationsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ReservationID:   ReservationIDColumn,
		ItemID:          ItemIDColumn,
		UserID:          UserIDColumn,
		Quantity:        QuantityColumn,
		ReservationDate: ReservationDateColumn,
		TimeFrom:        TimeFromColumn,
		TimeTo:          TimeToColumn,
		IsCancelled:     IsCancelledColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
