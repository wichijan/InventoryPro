package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type QuickShelfControllerI interface {
	GetQuickShelves() (*[]models.QuickShelfWithItems, *models.INVError)
	GetQuickShelfById(quickShelfId *uuid.UUID) (*models.QuickShelfWithItems, *models.INVError)
	CreateQuickShelf(quickShelf *models.QuickShelfCreate) (*uuid.UUID, *models.INVError)
	UpdateQuickShelf(quickShelf *model.QuickShelves) *models.INVError
	DeleteQuickShelf(quickShelfId *uuid.UUID) *models.INVError // Return to regular shelf & clear quick shelf
}

type QuickShelfController struct {
	QuickShelfRepo     repositories.QuickShelfRepositoryI
	ItemQuickShelfRepo repositories.ItemQuickShelfRepositoryI
}

func (qsc *QuickShelfController) GetQuickShelves() (*[]models.QuickShelfWithItems, *models.INVError) {
	return qsc.QuickShelfRepo.GetQuickShelves()
}

func (qsc *QuickShelfController) GetQuickShelfById(quickShelfId *uuid.UUID) (*models.QuickShelfWithItems, *models.INVError) {
	return qsc.QuickShelfRepo.GetQuickShelfById(quickShelfId)
}

func (qsc *QuickShelfController) CreateQuickShelf(quickShelf *models.QuickShelfCreate) (*uuid.UUID, *models.INVError) {
	tx, err := qsc.QuickShelfRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	id, inv_error := qsc.QuickShelfRepo.CreateQuickShelf(tx, quickShelf)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return id, nil
}

func (qsc *QuickShelfController) UpdateQuickShelf(quickShelf *model.QuickShelves) *models.INVError {
	tx, err := qsc.QuickShelfRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	inv_error := qsc.QuickShelfRepo.UpdateQuickShelf(tx, quickShelf)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return nil
}

func (qsc *QuickShelfController) DeleteQuickShelf(quickShelfId *uuid.UUID) *models.INVError {
	tx, err := qsc.QuickShelfRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	items, inv_error := qsc.ItemQuickShelfRepo.GetItemsInQuickShelf(quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	if len(*items) > 0 {
		return inv_errors.INV_CONFLICT.WithDetails("Quick shelf is not empty")
	}

	inv_error = qsc.QuickShelfRepo.DeleteQuickShelf(tx, quickShelfId)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}
	return nil
}
