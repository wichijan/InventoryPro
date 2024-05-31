package models

type ReservationCreate struct {
	ItemID string `alias:"reservations.item_id"`
	UserID string `alias:"reservations.user_id"`
	Quantity int `alias:"reservations.quantity"`
	TimeFrom string `alias:"reservations.time_from"`
	TimeTo string `alias:"reservations.time_to"`
}