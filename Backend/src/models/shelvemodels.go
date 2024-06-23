package models

import "github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"

type ShelveWithItems struct {
	ID string `alias:"shelves.id" sql:"primary_key"`

	Items []struct {
		model.Items
		Quantity int32 `alias:"items_in_shelf.quantity"`
	}

	Room struct {
		model.Rooms
	}
}

type ShelveOTD struct {
	RoomID string `json:"roomId" binding:"required"`
}
