package models

import "github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"

type ShelveWithItems struct {
	ID             string `alias:"shelves.id"`
	ShelveTypeName string `alias:"shelve_types.type_name"`
	RoomID         string `alias:"shelves.room_id"`

	Items []struct {
		model.Items
	}
}

type OwnShelve struct {
	ID             string `alias:"shelves.id"`
	ShelveTypeName string `alias:"shelve_types.type_name"`
	RoomID         string `alias:"shelves.room_id"`
}

type ShelveOTD struct {
	ShelveTypeName string `json:"shelve_type_name" binding:"required"`
	RoomID         string `json:"room_id" binding:"required"`
}