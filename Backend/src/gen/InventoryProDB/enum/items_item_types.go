//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/mysql"

var ItemsItemTypes = &struct {
	Book         mysql.StringExpression
	SingleObject mysql.StringExpression
	SetOfObjects mysql.StringExpression
}{
	Book:         mysql.NewEnumValue("book"),
	SingleObject: mysql.NewEnumValue("single_object"),
	SetOfObjects: mysql.NewEnumValue("set_of_objects"),
}
