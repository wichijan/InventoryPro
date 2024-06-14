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

type ItemSubjectRepositoryI interface {
	GetItemsForSubject(subjectId *string) (*[]model.ItemSubjects, *models.INVError)
	CreateSubjectForItem(tx *sql.Tx, keyword *models.ItemWithSubject) *models.INVError
	DeleteSubjectForItem(tx *sql.Tx, keyword *models.ItemWithSubject) *models.INVError
	DeleteSubjectsForItem(tx *sql.Tx, itemId *uuid.UUID) *models.INVError
	CheckIfSubjectAndItemExists(subjectAndItem models.ItemWithSubject) *models.INVError

	CheckIfItemIdExists(itemId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type ItemSubjectRepository struct {
	managers.DatabaseManagerI
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
	err := stmt.Query(isjr.GetDatabaseConnection(), &itemsForSubject)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading items for subject")
	}

	return &itemsForSubject, nil
}

func (isjr *ItemSubjectRepository) CreateSubjectForItem(tx *sql.Tx, keyword *models.ItemWithSubject) *models.INVError {

	// Create the insert statement
	insertQuery := table.ItemSubjects.INSERT(
		table.ItemSubjects.SubjectID,
		table.ItemSubjects.ItemID,
	).VALUES(
		keyword.SubjectID,
		keyword.ItemID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating keyword for item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating keyword for item")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Combination already exists")
	}

	return nil
}

func (isjr *ItemSubjectRepository) DeleteSubjectForItem(tx *sql.Tx, keyword *models.ItemWithSubject) *models.INVError {
	deleteQuery := table.ItemSubjects.DELETE().WHERE(
		table.ItemSubjects.SubjectID.EQ(mysql.String(keyword.SubjectID)).
			AND(table.ItemSubjects.ItemID.EQ(mysql.String(keyword.ItemID))),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting keyword for item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting keyword for item")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Combination not found")
	}

	return nil
}


func (isjr *ItemSubjectRepository) DeleteSubjectsForItem(tx *sql.Tx, itemId *uuid.UUID) *models.INVError {
	deleteQuery := table.ItemSubjects.DELETE().WHERE(
		table.ItemSubjects.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	rows, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting keyword for item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting keyword for item")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Combination not found")
	}

	return nil
}

func (isjr *ItemSubjectRepository) CheckIfSubjectAndItemExists(subjectAndItem models.ItemWithSubject) *models.INVError {
	count, err := utils.CountStatement(table.ItemSubjects, table.ItemSubjects.SubjectID.EQ(mysql.String(subjectAndItem.SubjectID)).AND(table.ItemSubjects.ItemID.EQ(mysql.String(subjectAndItem.ItemID))), isjr.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if subject and item exists")
	}
	if count > 0 {
		return inv_errors.INV_SUBJECT_ITEM_COMBI_EXISTS.WithDetails("Subject and item combination already exists")
	}
	return nil
}

func (isjr *ItemSubjectRepository) CheckIfItemIdExists(itemId *uuid.UUID) *models.INVError {
	count, err := utils.CountStatement(table.ItemSubjects, table.ItemSubjects.ItemID.EQ(mysql.String(itemId.String())), isjr.GetDatabaseConnection())
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if itemId exists in ItemSubjects table")
	}
	if count <= 0 {
		return inv_errors.INV_CONFLICT.WithDetails("ItemSubjects still has items in it")
	}
	return nil
}
