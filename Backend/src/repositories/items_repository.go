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

type ItemRepositoryI interface {
	GetItems() (*[]models.ItemWithEverything, *models.INVError)
	GetItemById(itemId *uuid.UUID) (*models.ItemWithEverything, *models.INVError)
	CreateItem(tx *sql.Tx, item *model.Items) (*uuid.UUID, *models.INVError)
	UpdateItem(tx *sql.Tx, item *model.Items) *models.INVError
	DeleteItem(tx *sql.Tx, itemId *uuid.UUID) *models.INVError

	StoreItemPicture(tx *sql.Tx, itemId *uuid.UUID) (*uuid.UUID, *models.INVError)
	GetPictureIdFromItem(itemId *uuid.UUID) (*uuid.UUID, *models.INVError)
	RemovePictureIdFromItem(tx *sql.Tx, itemId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type ItemRepository struct {
	managers.DatabaseManagerI
}

func (itr *ItemRepository) GetItems() (*[]models.ItemWithEverything, *models.INVError) {
	var items []models.ItemWithEverything

	// Create the query
	stmt := mysql.SELECT(
		table.Items.ID,
		table.Items.Name,
		table.Items.Description,
		table.Items.RegularShelfID,
		table.Items.ClassOne,
		table.Items.ClassTwo,
		table.Items.ClassThree,
		table.Items.ClassFour,
		table.Items.Damaged,
		table.Items.Picture,
		table.Items.HintText,
		table.Items.DamagedDescription,
		table.ItemsInShelf.Quantity,
		table.ItemSubjects.AllColumns,
		table.KeywordsForItems.AllColumns,
		table.Users.ID,
		table.Users.Username,
		table.ItemTypes.TypeName,
		table.Reservations.AllColumns,
	).FROM(
		table.Items.
			LEFT_JOIN(table.ItemsInShelf, table.ItemsInShelf.ItemID.EQ(table.Items.ID)).
			LEFT_JOIN(table.ItemTypes, table.ItemTypes.ID.EQ(table.Items.ItemTypeID)).
			LEFT_JOIN(table.UserItems, table.UserItems.ItemID.EQ(table.Items.ID)).
			LEFT_JOIN(table.ItemSubjects, table.ItemSubjects.ItemID.EQ(table.Items.ID)).
			LEFT_JOIN(table.KeywordsForItems, table.KeywordsForItems.ItemID.EQ(table.Items.ID)).
			LEFT_JOIN(table.Users, table.Users.ID.EQ(table.UserItems.UserID)).
			LEFT_JOIN(table.Reservations, table.Reservations.ItemID.EQ(table.Items.ID)),
	)

	// Execute the query
	err := stmt.Query(itr.GetDatabaseConnection(), &items)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading items")
	}

	return &items, nil
}

func (itr *ItemRepository) GetItemById(itemId *uuid.UUID) (*models.ItemWithEverything, *models.INVError) {
	var items models.ItemWithEverything

	// Create the query
	stmt := mysql.SELECT(
		table.Items.ID,
		table.Items.Name,
		table.Items.Description,
		table.Items.RegularShelfID,
		table.Items.ClassOne,
		table.Items.ClassTwo,
		table.Items.ClassThree,
		table.Items.ClassFour,
		table.Items.Damaged,
		table.Items.Picture,
		table.Items.HintText,
		table.Items.DamagedDescription,
		table.ItemsInShelf.Quantity,
		table.Users.ID,
		table.Users.Username,
		table.ItemSubjects.AllColumns,
		table.KeywordsForItems.AllColumns,
		table.Reservations.AllColumns,
	).FROM(
		table.Items.
			LEFT_JOIN(table.ItemsInShelf, table.ItemsInShelf.ItemID.EQ(table.Items.ID)).
			LEFT_JOIN(table.UserItems, table.UserItems.ItemID.EQ(table.Items.ID)).
			LEFT_JOIN(table.ItemSubjects, table.ItemSubjects.ItemID.EQ(table.Items.ID)).
			LEFT_JOIN(table.KeywordsForItems, table.KeywordsForItems.ItemID.EQ(table.Items.ID)).
			LEFT_JOIN(table.Users, table.Users.ID.EQ(table.UserItems.UserID)).
			LEFT_JOIN(table.Reservations, table.Reservations.ItemID.EQ(table.Items.ID)),
	).WHERE(
		table.Items.ID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	err := stmt.Query(itr.GetDatabaseConnection(), &items)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading items")
	}

	return &items, nil
}

func (itr *ItemRepository) CreateItem(tx *sql.Tx, item *model.Items) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Items.INSERT(
		table.Items.ID,
		table.Items.Name,
		table.Items.Description,
		table.Items.ClassOne,
		table.Items.ClassTwo,
		table.Items.ClassThree,
		table.Items.ClassFour,
		table.Items.Damaged,
		table.Items.DamagedDescription,
		table.Items.Picture,
		table.Items.HintText,
		table.Items.ItemTypeID,
		table.Items.RegularShelfID,
	).VALUES(
		uuid.String(),
		item.Name,
		item.Description,
		item.ClassOne,
		item.ClassTwo,
		item.ClassThree,
		item.ClassFour,
		item.Damaged,
		item.DamagedDescription,
		item.Picture,
		item.HintText,
		item.ItemTypeID,
		item.RegularShelfID,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating item")
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Item already exists")
	}

	return &uuid, nil
}

func (itr *ItemRepository) UpdateItem(tx *sql.Tx, item *model.Items) *models.INVError {
	// Create the update statement
	updateQuery := table.Items.UPDATE(
		table.Items.Name,
		table.Items.Description,
		table.Items.ClassOne,
		table.Items.ClassTwo,
		table.Items.ClassThree,
		table.Items.ClassFour,
		table.Items.Damaged,
		table.Items.DamagedDescription,
		table.Items.Picture,
		table.Items.HintText,
		table.Items.ItemTypeID,
		table.Items.RegularShelfID,
	).SET(
		item.Name,
		item.Description,
		item.ClassOne,
		item.ClassTwo,
		item.ClassThree,
		item.ClassFour,
		item.Damaged,
		item.DamagedDescription,
		item.Picture,
		item.HintText,
		item.ItemTypeID,
		item.RegularShelfID,
	).WHERE(table.Items.ID.EQ(mysql.String(item.ID)))

	// Execute the query
	_, err := updateQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating item")
	}

	return nil
}

func (itr *ItemRepository) DeleteItem(tx *sql.Tx, itemId *uuid.UUID) *models.INVError {
	// TODO - Implement DeleteWarehouse
	return inv_errors.INV_INTERNAL_ERROR.WithDetails("DeleteItem not implemented")
}

func (itr *ItemRepository) StoreItemPicture(tx *sql.Tx, itemId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	updatePictureQuery := table.Items.UPDATE(
		table.Items.Picture,
	).SET(
		uuid.String(),
	).WHERE(table.Items.ID.EQ(mysql.String(itemId.String())))

	// Execute the query
	rows, err := updatePictureQuery.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error storing picture for item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error storing picture for item")
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Item not found")
	}

	return &uuid, nil
}

func (itr *ItemRepository) GetPictureIdFromItem(itemId *uuid.UUID) (*uuid.UUID, *models.INVError) {
	var picture models.ItemPicture

	// Create the query
	stmt := mysql.SELECT(
		table.Items.Picture,
	).FROM(
		table.Items,
	).WHERE(
		table.Items.ID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	err := stmt.Query(itr.GetDatabaseConnection(), &picture)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading picture for item")
	}

	if picture.PictureId == "" {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Picture not found")
	}

	pictureId, err := uuid.Parse(picture.PictureId)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error parsing picture id")
	}

	return &pictureId, nil
}

func (itr *ItemRepository) RemovePictureIdFromItem(tx *sql.Tx, itemId *uuid.UUID) *models.INVError {
	// Create the update statement
	updatePictureQuery := table.Items.UPDATE(
		table.Items.Picture,
	).SET(
		"",
	).WHERE(table.Items.ID.EQ(mysql.String(itemId.String())))

	// Execute the query
	rows, err := updatePictureQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error removing picture for item")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error removing picture for item")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Item not found")
	}

	return nil
}
