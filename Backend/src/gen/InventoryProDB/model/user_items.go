//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type UserItems struct {
	UserID          string `sql:"primary_key"`
	ItemID          string `sql:"primary_key"`
	Quantity        *int32
	TransactionDate *time.Time
}
