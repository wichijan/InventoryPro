//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Books struct {
	ItemID    string `sql:"primary_key"`
	Isbn      string
	Author    *string
	Publisher *string
	Edition   *string
}
