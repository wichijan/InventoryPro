package models

import "github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"

type WarehouseWithRooms struct {
	model.Warehouses

	Rooms []struct {
		model.Rooms
	}
}

type WarehousesODT struct {
	Name        *string
	Description *string
}