package models

type GetQuantityReserved struct {
	Quantity int32 `alias:"user_items.quantity"`
}