//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type ItemsItemTypes string

const (
	ItemsItemTypes_Book         ItemsItemTypes = "book"
	ItemsItemTypes_SingleObject ItemsItemTypes = "single_object"
	ItemsItemTypes_SetOfObjects ItemsItemTypes = "set_of_objects"
)

func (e *ItemsItemTypes) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "book":
		*e = ItemsItemTypes_Book
	case "single_object":
		*e = ItemsItemTypes_SingleObject
	case "set_of_objects":
		*e = ItemsItemTypes_SetOfObjects
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for ItemsItemTypes enum")
	}

	return nil
}

func (e ItemsItemTypes) String() string {
	return string(e)
}
