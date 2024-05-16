package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/utils"
)

type ItemSubjectRepositoryI interface {
	GetItemsForSubject(subjectId *string) (*[]model.ItemSubjects, *models.INVError)
	CreateSubjectForItem(keyword *models.ItemWithSubject) (*uuid.UUID, *models.INVError)
	DeleteSubjectForItem(keyword *models.ItemWithSubject) *models.INVError
	CheckIfSubjectAndItemExists(subjectAndItem models.ItemWithSubject) *models.INVError
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

func (isjr *ItemSubjectRepository) CreateSubjectForItem(keyword *models.ItemWithSubject) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.ItemSubjects.INSERT(
		table.ItemSubjects.ID,
		table.ItemSubjects.SubjectID,
		table.ItemSubjects.ItemID,
	).VALUES(
		uuid.String(),
		keyword.SubjectID,
		keyword.ItemID,
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

func (isjr *ItemSubjectRepository) DeleteSubjectForItem(keyword *models.ItemWithSubject) *models.INVError {
	deleteQuery := table.ItemSubjects.DELETE().WHERE(
		table.ItemSubjects.SubjectID.EQ(mysql.String(keyword.SubjectID)).
			AND(table.ItemSubjects.ItemID.EQ(mysql.String(keyword.ItemID))),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(isjr.DatabaseManager.GetDatabaseConnection())
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

func (isjr *ItemSubjectRepository) CheckIfSubjectAndItemExists(subjectAndItem models.ItemWithSubject) *models.INVError {
	count, err := utils.CountStatement(table.ItemSubjects, table.ItemSubjects.SubjectID.EQ(mysql.String(subjectAndItem.SubjectID)).AND(table.ItemSubjects.ItemID.EQ(mysql.String(subjectAndItem.ItemID))), isjr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	if count > 0 {
		return inv_errors.INV_SUBJECT_ITEM_COMBI_EXISTS
	}
	return nil
}
