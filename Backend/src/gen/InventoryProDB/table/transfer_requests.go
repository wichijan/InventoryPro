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

var TransferRequests = newTransferRequestsTable("InventoryProDB", "transfer_requests", "")

type transferRequestsTable struct {
	mysql.Table

	// Columns
	TransferRequestID mysql.ColumnString
	ItemID            mysql.ColumnString
	UserID            mysql.ColumnString
	TargetUserID      mysql.ColumnString
	RequestDate       mysql.ColumnTimestamp
	IsAccepted        mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type TransferRequestsTable struct {
	transferRequestsTable

	NEW transferRequestsTable
}

// AS creates new TransferRequestsTable with assigned alias
func (a TransferRequestsTable) AS(alias string) *TransferRequestsTable {
	return newTransferRequestsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TransferRequestsTable with assigned schema name
func (a TransferRequestsTable) FromSchema(schemaName string) *TransferRequestsTable {
	return newTransferRequestsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TransferRequestsTable with assigned table prefix
func (a TransferRequestsTable) WithPrefix(prefix string) *TransferRequestsTable {
	return newTransferRequestsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TransferRequestsTable with assigned table suffix
func (a TransferRequestsTable) WithSuffix(suffix string) *TransferRequestsTable {
	return newTransferRequestsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTransferRequestsTable(schemaName, tableName, alias string) *TransferRequestsTable {
	return &TransferRequestsTable{
		transferRequestsTable: newTransferRequestsTableImpl(schemaName, tableName, alias),
		NEW:                   newTransferRequestsTableImpl("", "new", ""),
	}
}

func newTransferRequestsTableImpl(schemaName, tableName, alias string) transferRequestsTable {
	var (
		TransferRequestIDColumn = mysql.StringColumn("transfer_request_id")
		ItemIDColumn            = mysql.StringColumn("item_id")
		UserIDColumn            = mysql.StringColumn("user_id")
		TargetUserIDColumn      = mysql.StringColumn("target_user_id")
		RequestDateColumn       = mysql.TimestampColumn("request_date")
		IsAcceptedColumn        = mysql.BoolColumn("is_accepted")
		allColumns              = mysql.ColumnList{TransferRequestIDColumn, ItemIDColumn, UserIDColumn, TargetUserIDColumn, RequestDateColumn, IsAcceptedColumn}
		mutableColumns          = mysql.ColumnList{ItemIDColumn, UserIDColumn, TargetUserIDColumn, RequestDateColumn, IsAcceptedColumn}
	)

	return transferRequestsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		TransferRequestID: TransferRequestIDColumn,
		ItemID:            ItemIDColumn,
		UserID:            UserIDColumn,
		TargetUserID:      TargetUserIDColumn,
		RequestDate:       RequestDateColumn,
		IsAccepted:        IsAcceptedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
