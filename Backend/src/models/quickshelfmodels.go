package models

import (
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
)

type ItemQuickShelfInsert struct {
	QuickShelfID uuid.UUID
	UserID       uuid.UUID
	ItemID       uuid.UUID
	Quantity     int32
}

type ItemQuickShelfInsertODT struct {
	QuickShelfID uuid.UUID `binding:"required"`
	ItemID       uuid.UUID `binding:"required"`
	Quantity     int32     `binding:"required"`
}

type GetQuantity struct {
	Quantity *int32 `alias:"item_quick_shelf.quantity"`
}

type ItemQuickShelfRemoveSingleItem struct {
	QuickShelfID uuid.UUID `json:"quickShelfId"`
	ItemID       uuid.UUID `json:"itemId"`
}

type QuickShelfCreate struct {
	RoomId uuid.UUID
}

type QuickShelfWithItems struct {
	QuickShelfID uuid.UUID `alias:"quick_shelves.quick_shelf_id" json:"quickShelfId"`
	RoomID       uuid.UUID `alias:"quick_shelves.room_id" json:"roomId"`

	Items []struct {
		model.Items
	}
}

type GetQuickShelf struct {
	QuickShelfID uuid.UUID `alias:"item_quick_shelf.quick_shelf_id"`
	Quantity     int32     `alias:"item_quick_shelf.quantity"`

	Items struct {
		model.Items
	}

	Users struct {
		model.Users
	}
}
