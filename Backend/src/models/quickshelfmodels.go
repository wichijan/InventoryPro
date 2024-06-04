package models

import (
	"github.com/google/uuid"
)

type ItemQuickShelfInsert struct {
	QuickShelfID uuid.UUID `json:"quickShelfId"`
	UserID       uuid.UUID `json:"userId"`
	ItemID       uuid.UUID `json:"itemId"`
	Quantity     int32     `json:"quantity"`
}

type ItemQuickShelfInsertODT struct {
	QuickShelfID uuid.UUID `json:"quickShelfId"`
	ItemID       uuid.UUID `json:"itemId"`
	Quantity     int32     `json:"quantity"`
}

type GetQuantity struct {
	Quantity *int32 `json:"quantity"`
}

type ItemQuickShelfRemoveSingleItem struct {
	QuickShelfID uuid.UUID `json:"quickShelfId"`
	ItemID       uuid.UUID `json:"itemId"`
}
