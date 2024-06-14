package repositories

import (
	"database/sql"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

type RoomRepositoryI interface {
	GetRooms() (*[]model.Rooms, *models.INVError)
	GetRoomsById(id *uuid.UUID) (*model.Rooms, *models.INVError)
	GetRoomsWithShelves() (*[]models.RoomWithShelves, *models.INVError)
	GetRoomsByIdWithShelves(id *uuid.UUID) (*models.RoomWithShelves, *models.INVError)
	CreateRoom(tx *sql.Tx, room *models.RoomsODT) (*uuid.UUID, *models.INVError)
	UpdateRoom(tx *sql.Tx, room *model.Rooms) *models.INVError
	DeleteRoom(tx *sql.Tx, roomId *uuid.UUID) *models.INVError

	CheckIfWarehouseIdExists(warehouseId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type RoomRepository struct {
	managers.DatabaseManagerI
}

func (ror *RoomRepository) GetRooms() (*[]model.Rooms, *models.INVError) {
	var rooms []model.Rooms

	// Create the query
	stmt := mysql.SELECT(
		table.Rooms.AllColumns,
	).FROM(
		table.Rooms,
	)

	// Execute the query
	err := stmt.Query(ror.GetDatabaseConnection(), &rooms)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading rooms")
	}

	return &rooms, nil
}

func (wr *RoomRepository) GetRoomsById(id *uuid.UUID) (*model.Rooms, *models.INVError) {
	var rooms model.Rooms

	// Create the query
	stmt := mysql.SELECT(
		table.Rooms.AllColumns,
	).FROM(
		table.Rooms,
	).WHERE(
		table.Rooms.ID.EQ(utils.MySqlString(id.String())),
	)

	// Execute the query
	err := stmt.Query(wr.GetDatabaseConnection(), &rooms)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND.WithDetails("Room id not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading room")
	}

	return &rooms, nil
}

func (ror *RoomRepository) CreateRoom(tx *sql.Tx, room *models.RoomsODT) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Rooms.INSERT(
		table.Rooms.ID,
		table.Rooms.Name,
		table.Rooms.WarehouseID,
	).VALUES(
		uuid.String(),
		room.Name,
		room.WarehouseID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating room")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating room")
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Room already exists")
	}

	return &uuid, nil
}

func (ror *RoomRepository) UpdateRoom(tx *sql.Tx, room *model.Rooms) *models.INVError {
	// Create the update statement
	updateQuery := table.Rooms.UPDATE(
		table.Rooms.Name,
		table.Rooms.WarehouseID,
	).SET(
		room.Name,
		room.WarehouseID,
	).WHERE(table.Rooms.ID.EQ(mysql.String(room.ID)))

	// Execute the query
	_, err := updateQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating room")
	}

	return nil
}

func (ror *RoomRepository) DeleteRoom(tx *sql.Tx, roomId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.Rooms.DELETE().WHERE(
		table.Rooms.ID.EQ(mysql.String(roomId.String())),
	)

	// Execute the query
	rows, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting room")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting room")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Room Id not found")
	}

	return nil
}

func (wr *RoomRepository) GetRoomsWithShelves() (*[]models.RoomWithShelves, *models.INVError) {
	var rooms []models.RoomWithShelves

	// Create the query
	stmt := mysql.SELECT(
		table.Rooms.AllColumns,
		table.Shelves.AllColumns,
	).FROM(
		table.Rooms.
			LEFT_JOIN(table.Shelves, table.Shelves.RoomID.EQ(table.Rooms.ID)),
	)

	// Execute the query
	err := stmt.Query(wr.GetDatabaseConnection(), &rooms)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND.WithDetails("No rooms found with shelves")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading rooms")
	}

	return &rooms, nil
}

func (wr *RoomRepository) GetRoomsByIdWithShelves(id *uuid.UUID) (*models.RoomWithShelves, *models.INVError) {
	var rooms models.RoomWithShelves

	// Create the query
	stmt := mysql.SELECT(
		table.Rooms.AllColumns,
		table.Shelves.AllColumns,
	).FROM(
		table.Rooms.
			LEFT_JOIN(table.Shelves, table.Shelves.RoomID.EQ(table.Rooms.ID)),
	).WHERE(
		table.Rooms.ID.EQ(utils.MySqlString(id.String())),
	)

	// Execute the query
	err := stmt.Query(wr.GetDatabaseConnection(), &rooms)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, inv_errors.INV_NOT_FOUND.WithDetails("Room id not found")
		}
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading room")
	}

	return &rooms, nil
}

func (wr *RoomRepository) CheckIfWarehouseIdExists(warehouseId *uuid.UUID) *models.INVError {
	count, err := utils.CountStatement(table.Rooms, table.Rooms.WarehouseID.EQ(mysql.String(warehouseId.String())), wr.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if warehouseId exists in Rooms table")
	}
	if count <= 0 {
		return inv_errors.INV_CONFLICT.WithDetails("Rooms still has warehouseId in it")
	}
	return nil
}