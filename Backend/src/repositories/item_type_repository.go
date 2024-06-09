package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
)

type ItemTypeRepositoryI interface {
	GetItemTypes() (*[]models.ItemTypes, *models.INVError)
	GetItemTypesByName(typeName *string) (*model.ItemTypes, *models.INVError)
}

type ItemTypeRepository struct {
	managers.DatabaseManagerI
}

func (itr *ItemTypeRepository) GetItemTypes() (*[]models.ItemTypes, *models.INVError) {
	var itemTypes []models.ItemTypes

	// Create the query
	stmt := mysql.SELECT(
		table.ItemTypes.TypeName,
	).FROM(
		table.Items,
	)

	// Execute the query
	err := stmt.Query(itr.GetDatabaseConnection(), &itemTypes)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading item types")
	}

	return &itemTypes, nil
}

func (itr *ItemTypeRepository) GetItemTypesByName(typeName *string) (*model.ItemTypes, *models.INVError) {
	var itemType model.ItemTypes

	// Create the query
	stmt := mysql.SELECT(
		table.ItemTypes.AllColumns,
	).FROM(
		table.ItemTypes,
	).WHERE(table.ItemTypes.TypeName.EQ(mysql.String(*typeName)))

	// Execute the query
	err := stmt.Query(itr.GetDatabaseConnection(), &itemType)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading item type")
	}

	if itemType.TypeName == nil {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Item type not found")
	}

	return &itemType, nil
}