//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Items struct {
	ID                 string `sql:"primary_key"`
	Name               *string
	Description        *string
	ClassOne           *bool
	ClassTwo           *bool
	ClassThree         *bool
	ClassFour          *bool
	Damaged            *bool
	DamagedDescription *string
	Quantity           *int32
	Picture            *string
	StatusID           *string
}
