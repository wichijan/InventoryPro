package controllers

import (
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type KeywordControllerI interface {
	GetKeywords() (*[]model.Keywords, *models.INVError)
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
