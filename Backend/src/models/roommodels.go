package models

import "github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"

type RoomWithShelves struct {
	model.Rooms

	Shelves []struct {
		model.Shelves
	}
}

type RoomsODT struct {
	Name        *string
	WarehouseID *string
}