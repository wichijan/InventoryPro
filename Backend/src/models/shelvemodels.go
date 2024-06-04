package models

import "github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"

type ShelveWithItems struct {
	ID             string `alias:"shelves.id"`
	RoomID         string `alias:"shelves.room_id"`

	Items []struct {
		model.Items
	}
}

type ShelveOTD struct {
	RoomID         string `json:"room_id" binding:"required"`
}