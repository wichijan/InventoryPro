package controllers

import (
	"github.com/google/uuid"
	inv_errors "github.com/wichijan/InventoryPro/src/errors"
	"github.com/wichijan/InventoryPro/src/gen/InventoryProDB/model"
	"github.com/wichijan/InventoryPro/src/models"
	"github.com/wichijan/InventoryPro/src/repositories"
)

type SubjectControllerI interface {
	GetSubjects() (*[]model.Subjects, *models.INVError)
	GetSubjectByName(subjectName *string) (*model.Subjects, *models.INVError)
	CreateSubject(subject *models.SubjectODT) (*uuid.UUID, *models.INVError)
	UpdateSubject(subject *model.Subjects) *models.INVError
	DeleteSubject(subjectId *uuid.UUID) *models.INVError
}

type SubjectController struct {
	SubjectRepo repositories.SubjectRepositoryI
}

func (sc *SubjectController) GetSubjects() (*[]model.Subjects, *models.INVError) {
	subjects, inv_error := sc.SubjectRepo.GetSubjects()
	if inv_error != nil {
		return nil, inv_error
	}
	return subjects, nil
}

func (sc *SubjectController) GetSubjectByName(subjectName *string) (*model.Subjects, *models.INVError) {
	subject, inv_error := sc.SubjectRepo.GetSubjectByName(subjectName)
	if inv_error != nil {
		return nil, inv_error
	}
	return subject, nil
}

func (sc *SubjectController) CreateSubject(subject *models.SubjectODT) (*uuid.UUID, *models.INVError) {
	tx, err := sc.SubjectRepo.NewTransaction()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	subjectId, inv_error := sc.SubjectRepo.CreateSubject(tx, subject)
	if inv_error != nil {
		return nil, inv_error
	}

	if err = tx.Commit(); err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return subjectId, nil
}

func (sc *SubjectController) UpdateSubject(subject *model.Subjects) *models.INVError {
	tx, err := sc.SubjectRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	inv_error := sc.SubjectRepo.UpdateSubject(tx, subject)
	if inv_error != nil {
		return inv_error
	}

	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}

func (sc *SubjectController) DeleteSubject(subjectId *uuid.UUID) *models.INVError {
	tx, err := sc.SubjectRepo.NewTransaction()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating transaction")
	}
	defer tx.Rollback()

	inv_error := sc.SubjectRepo.DeleteSubject(tx, subjectId)
	if inv_error != nil {
		return inv_error
	}
	
	if err = tx.Commit(); err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error committing transaction")
	}

	return nil
}
