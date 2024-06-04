package repositories

import (
	"github.com/go-jet/jet/v2/mysql"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/managers"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/table"
)

type ItemTypeRepositoryI interface {
	GetItemTypes() (*[]models.ItemTypes, *models.INVError)
	GetItemTypesByName(typeName *string) (*models.ItemTypes, *models.INVError)
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
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if len(itemTypes) == 0 {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &itemTypes, nil
}

func (itr *ItemTypeRepository) GetItemTypesByName(typeName *string) (*models.ItemTypes, *models.INVError) {
	var itemType models.ItemTypes

	// Create the query
	stmt := mysql.SELECT(
		table.ItemTypes.TypeName,
	).FROM(
		table.ItemTypes,
	).WHERE(table.ItemTypes.TypeName.EQ(mysql.String(*typeName)))

	// Execute the query
	err := stmt.Query(itr.GetDatabaseConnection(), &itemType)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	if itemType.TypeName == "" {
		return nil, inv_errors.INV_NOT_FOUND
	}

	return &itemType, nil
}