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
)

type PointRepositoryI interface {
	GetPointsByUserId(userId *uuid.UUID) (*model.Points, *models.INVError)
	CreateBook(tx *sql.Tx, points *model.Points) *models.INVError
	AddPointsToUser(tx *sql.Tx, userId *uuid.UUID, points int) *models.INVError
	SubtractPointsToUser(tx *sql.Tx, userId *uuid.UUID, points int) *models.INVError

	managers.DatabaseManagerI
}

type PointRepository struct {
	managers.DatabaseManagerI
}

func (pr *PointRepository) GetPointsByUserId(userId *uuid.UUID) (*model.Points, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.Points.AllColumns,
	).FROM(
		table.Points,
	).WHERE(
		table.Points.UserID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	var points model.Points
	err := stmt.Query(pr.GetDatabaseConnection(), &points)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &points, nil
}

func (pr *PointRepository) CreateBook(tx *sql.Tx, points *model.Points) *models.INVError {
	// Create the query
	stmt := table.Points.INSERT(
		table.Points.UserID,
		table.Points.Points,
	).VALUES(
		points.UserID,
		points.Points,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (pr *PointRepository) AddPointsToUser(tx *sql.Tx, userId *uuid.UUID, points int) *models.INVError {
	// Create the query
	stmt := table.Points.UPDATE(
		table.Points.Points,
	).SET(
		points,
	).WHERE(
		table.Points.UserID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}

func (pr *PointRepository) SubtractPointsToUser(tx *sql.Tx, userId *uuid.UUID, points int) *models.INVError {
	// TODO Check if minus points is possible
	
	// Create the query
	stmt := table.Points.UPDATE(
		table.Points.Points,
	).SET(
		points,
	).WHERE(
		table.Points.UserID.EQ(mysql.String(userId.String())),
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	return nil
}