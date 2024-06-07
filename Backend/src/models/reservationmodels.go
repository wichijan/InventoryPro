package models

import "time"

type ReservationCreate struct {
	ItemID   string    `alias:"reservations.item_id"`
	UserID   string    `alias:"reservations.user_id"`
	Quantity int       `alias:"reservations.quantity"`
	TimeFrom time.Time `alias:"reservations.time_from"`
	TimeTo   time.Time `alias:"reservations.time_to"`
}

type ReservationCreateODT struct {
	ItemID   string `alias:"reservations.item_id" binding:"required"`
	Quantity int    `alias:"reservations.quantity" binding:"required"`
	TimeFrom string `alias:"reservations.time_from" binding:"required"`
	TimeTo   string `alias:"reservations.time_to" binding:"required"`
}
