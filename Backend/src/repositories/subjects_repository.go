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

type SubjectRepositoryI interface {
	GetSubjects() (*[]model.Subjects, *models.INVError)
	GetSubjectByName(subjectName *string) (*model.Subjects, *models.INVError)
	CreateSubject(tx *sql.Tx, subject *models.SubjectODT) (*uuid.UUID, *models.INVError)
	UpdateSubject(tx *sql.Tx, subject *model.Subjects) *models.INVError
	DeleteSubject(tx *sql.Tx, subjectId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type SubjectRepository struct {
	managers.DatabaseManagerI
}

func (sr *SubjectRepository) GetSubjects() (*[]model.Subjects, *models.INVError) {
	var subjects []model.Subjects

	// Create the query
	stmt := mysql.SELECT(
		table.Subjects.AllColumns,
	).FROM(
		table.Subjects,
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &subjects)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading subjects")
	}

	return &subjects, nil
}

func (sr *SubjectRepository) GetSubjectByName(subjectName *string) (*model.Subjects, *models.INVError) {
	var subject model.Subjects

	// Create the query
	stmt := mysql.SELECT(
		table.Subjects.AllColumns,
	).FROM(
		table.Subjects,
	).WHERE(
		table.Subjects.Name.EQ(mysql.String(*subjectName)),
	)

	// Execute the query
	err := stmt.Query(sr.GetDatabaseConnection(), &subject)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading subject")
	}

	if subject.ID == "" {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Subject not found")
	}

	return &subject, nil
}

func (sr *SubjectRepository) CreateSubject(tx *sql.Tx, subject *models.SubjectODT) (*uuid.UUID, *models.INVError) {
	uuid := uuid.New()

	// Create the insert statement
	insertQuery := table.Subjects.INSERT(
		table.Subjects.ID,
		table.Subjects.Name,
		table.Subjects.Description,
	).VALUES(
		uuid.String(),
		subject.Name,
		subject.Description,
	)

	// Execute the query
	rows, err := insertQuery.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating subject")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating subject")
	}

	if rowsAff == 0 {
		return nil, inv_errors.INV_NOT_FOUND.WithDetails("Subject already exists")
	}

	return &uuid, nil
}

func (sr *SubjectRepository) UpdateSubject(tx *sql.Tx, subject *model.Subjects) *models.INVError {
	// Create the update statement
	updateQuery := table.Subjects.UPDATE(
		table.Subjects.Name,
		table.Subjects.Description,
	).SET(
		subject.Name,
		subject.Description,
	).WHERE(table.Subjects.ID.EQ(mysql.String(subject.ID)))

	// Execute the query
	_, err := updateQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating subject")
	}

	return nil
}

func (sr *SubjectRepository) DeleteSubject(tx *sql.Tx, subjectId *uuid.UUID) *models.INVError {
	// Create the delete statement
	deleteQuery := table.Subjects.DELETE().WHERE(table.Subjects.ID.EQ(mysql.String(subjectId.String())))

	// Execute the query
	rows, err := deleteQuery.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting subject")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting subject")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Subject id not found")
	}

	return nil
}
