package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type KeywordControllerI interface {
	GetKeywords() (*[]model.Keywords, *models.INVError)
	CreateKeyword(keywordName *string) (*uuid.UUID, *models.INVError)
	UpdateKeyword(keyword *model.Keywords) *models.INVError
	DeleteKeyword(keywordId *uuid.UUID) *models.INVError
}

type KeywordController struct {
	keywordRepo repositories.KeywordRepositoryI
}

func (kc *KeywordController) GetKeywords() (*[]model.Keywords, *models.INVError) {
	keywords, inv_error := kc.keywordRepo.GetKeywords()
	if inv_error != nil {
		return nil, inv_error
	}

	return keywords, nil
}

func (kc *KeywordController) CreateKeyword(keywordName *string) (*uuid.UUID, *models.INVError) {
	tx, err := kc.keywordRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	keywordId, inv_error := kc.keywordRepo.CreateKeyword(tx, keywordName)
	if inv_error != nil {
		return nil, inv_error
	}

	return keywordId, nil
}

func (kc *KeywordController) UpdateKeyword(keyword *model.Keywords) *models.INVError {
	tx, err := kc.keywordRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	inv_error := kc.keywordRepo.UpdateKeyword(tx, keyword)
	if inv_error != nil {
		return inv_error
	}

	return nil
}

func (kc *KeywordController) DeleteKeyword(keywordId *uuid.UUID) *models.INVError {
	tx, err := kc.keywordRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}
	defer tx.Rollback()

	inv_error := kc.keywordRepo.DeleteKeyword(tx, keywordId)
	if inv_error != nil {
		return inv_error
	}

	return nil
}
