package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type ItemSubjectRepositoryI interface {
	GetItemsForSubject(subjectId *string) (*[]model.ItemSubjects, *models.INVError)
	GetSubjectsForItem(itemId *string) (*[]model.ItemSubjects, *models.INVError)
	CreateItemForSubject(itemForSubject *model.ItemSubjects) (*uuid.UUID, *models.INVError)
	UpdateItemForSubject(itemForSubject *model.ItemSubjects) *models.INVError
	DeleteItemForSubject(itemIdForSubject *uuid.UUID) *models.INVError
}

type ItemSubjectRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (isjr *ItemSubjectRepository) GetItemsForSubject(subjectId *string) (*[]model.ItemSubjects, *models.INVError) {
	var itemsForSubject []model.ItemSubjects

	// Create the query
	stmt := mysql.SELECT(
		table.ItemSubjects.AllColumns,
	).FROM(
		table.ItemSubjects,
	).WHERE(table.ItemSubjects.SubjectID.EQ(mysql.String(*subjectId)))

	// Execute the query
	err := stmt.Query(isjr.DatabaseManager.GetDatabaseConnection(), &itemsForSubject)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(itemsForSubject) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &itemsForSubject, nil
}

func (isjr *ItemSubjectRepository) GetSubjectsForItem(itemId *string) (*[]model.ItemSubjects, *models.INVError) {
	var itemWithSubject []model.ItemSubjects

	// Create the query
	stmt := mysql.SELECT(
		table.ItemSubjects.AllColumns,
	).FROM(
		table.ItemSubjects,
	).WHERE(table.ItemSubjects.ItemID.EQ(mysql.String(*itemId)))

	// Execute the query
	err := stmt.Query(isjr.DatabaseManager.GetDatabaseConnection(), &itemWithSubject)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(itemWithSubject) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &itemWithSubject, nil
}

func (isjr *ItemSubjectRepository) CreateItemForSubject(itemForSubject *model.ItemSubjects) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.ItemSubjects.INSERT(
		table.ItemSubjects.ID,
		table.ItemSubjects.ItemID,
		table.ItemSubjects.SubjectID,
	).VALUES(
		uuid.String(),
		itemForSubject.ItemID,
		itemForSubject.SubjectID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(isjr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &uuid, nil
}

func (isjr *ItemSubjectRepository) UpdateItemForSubject(itemForSubject *model.ItemSubjects) *models.INVError {
	// Create the update statement
	updateQuery := table.ItemSubjects.UPDATE(
		table.ItemSubjects.ItemID,
		table.ItemSubjects.SubjectID,
	).SET(
		itemForSubject.ItemID,
		itemForSubject.SubjectID,
	).WHERE(table.ItemSubjects.ID.EQ(mysql.String(itemForSubject.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(isjr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND
	}

	return nil
}

func (isjr *ItemSubjectRepository) DeleteItemForSubject(itemIdForSubject *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return nil
}
