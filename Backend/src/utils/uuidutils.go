package utils

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/models"
)

func NewUUID() *uuid.UUID {
	uuid := uuid.New()
	return &uuid
}

func ConvertStringToUUID(uuidString string) (*uuid.UUID, *models.INVError) {
	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error parsing UUID: " + uuidString)
	}

	return &uuid, nil
}
