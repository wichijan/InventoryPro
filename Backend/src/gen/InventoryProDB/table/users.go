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

var Users = newUsersTable("InventoryProDB", "users", "")

type usersTable struct {
	mysql.Table

	// Columns
	ID                   mysql.ColumnString
	FirstName            mysql.ColumnString
	LastName             mysql.ColumnString
	Username             mysql.ColumnString
	Email                mysql.ColumnString
	Password             mysql.ColumnString
	JobTitle             mysql.ColumnString
	PhoneNumber          mysql.ColumnString
	UserTypeID           mysql.ColumnString
	ProfilePicture       mysql.ColumnString
	RegistrationTime     mysql.ColumnTimestamp
	RegistrationAccepted mysql.ColumnBool
	IsActive             mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type UsersTable struct {
	usersTable

	NEW usersTable
}

// AS creates new UsersTable with assigned alias
func (a UsersTable) AS(alias string) *UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UsersTable with assigned schema name
func (a UsersTable) FromSchema(schemaName string) *UsersTable {
	return newUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UsersTable with assigned table prefix
func (a UsersTable) WithPrefix(prefix string) *UsersTable {
	return newUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UsersTable with assigned table suffix
func (a UsersTable) WithSuffix(suffix string) *UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUsersTable(schemaName, tableName, alias string) *UsersTable {
	return &UsersTable{
		usersTable: newUsersTableImpl(schemaName, tableName, alias),
		NEW:        newUsersTableImpl("", "new", ""),
	}
}

func newUsersTableImpl(schemaName, tableName, alias string) usersTable {
	var (
		IDColumn                   = mysql.StringColumn("id")
		FirstNameColumn            = mysql.StringColumn("first_name")
		LastNameColumn             = mysql.StringColumn("last_name")
		UsernameColumn             = mysql.StringColumn("username")
		EmailColumn                = mysql.StringColumn("email")
		PasswordColumn             = mysql.StringColumn("password")
		JobTitleColumn             = mysql.StringColumn("job_title")
		PhoneNumberColumn          = mysql.StringColumn("phone_number")
		UserTypeIDColumn           = mysql.StringColumn("user_type_id")
		ProfilePictureColumn       = mysql.StringColumn("profile_picture")
		RegistrationTimeColumn     = mysql.TimestampColumn("registration_time")
		RegistrationAcceptedColumn = mysql.BoolColumn("registration_accepted")
		IsActiveColumn             = mysql.BoolColumn("is_active")
		allColumns                 = mysql.ColumnList{IDColumn, FirstNameColumn, LastNameColumn, UsernameColumn, EmailColumn, PasswordColumn, JobTitleColumn, PhoneNumberColumn, UserTypeIDColumn, ProfilePictureColumn, RegistrationTimeColumn, RegistrationAcceptedColumn, IsActiveColumn}
		mutableColumns             = mysql.ColumnList{FirstNameColumn, LastNameColumn, UsernameColumn, EmailColumn, PasswordColumn, JobTitleColumn, PhoneNumberColumn, UserTypeIDColumn, ProfilePictureColumn, RegistrationTimeColumn, RegistrationAcceptedColumn, IsActiveColumn}
	)

	return usersTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                   IDColumn,
		FirstName:            FirstNameColumn,
		LastName:             LastNameColumn,
		Username:             UsernameColumn,
		Email:                EmailColumn,
		Password:             PasswordColumn,
		JobTitle:             JobTitleColumn,
		PhoneNumber:          PhoneNumberColumn,
		UserTypeID:           UserTypeIDColumn,
		ProfilePicture:       ProfilePictureColumn,
		RegistrationTime:     RegistrationTimeColumn,
		RegistrationAccepted: RegistrationAcceptedColumn,
		IsActive:             IsActiveColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
