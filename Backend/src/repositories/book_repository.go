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

type BookRepositoryI interface {
	GetBookById(itemId *uuid.UUID) (*model.Books, *models.INVError)
	CreateBook(tx *sql.Tx, book *model.Books) (*string, *models.INVError)
	UpdateBook(tx *sql.Tx, book *model.Books) *models.INVError
	DeleteBook(tx *sql.Tx, bookId *uuid.UUID) *models.INVError

	managers.DatabaseManagerI
}

type BookRepository struct {
	managers.DatabaseManagerI
}

func (br *BookRepository) GetBookById(itemId *uuid.UUID) (*model.Books, *models.INVError) {
	// Create the query
	stmt := mysql.SELECT(
		table.Books.AllColumns,
	).FROM(
		table.Books,
	).WHERE(
		table.Books.ItemID.EQ(mysql.String(itemId.String())),
	)

	// Execute the query
	var book model.Books
	err := stmt.Query(br.GetDatabaseConnection(), &book)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &book, nil
}

func (br *BookRepository) CreateBook(tx *sql.Tx, book *model.Books) (*string, *models.INVError) {
	// Create the query
	stmt := table.Books.INSERT(
		table.Books.ItemID,
		table.Books.Isbn,
		table.Books.Author,
		table.Books.Publisher,
		table.Books.Edition,
	).VALUES(
		book.ItemID,
		book.Isbn,
		book.Author,
		book.Publisher,
		book.Edition,
	)

	// Execute the query
	_, err := stmt.Exec(tx)
	if err != nil {
		return nil, inv_errors.INV_INTERNAL_ERROR
	}

	return &book.ItemID, nil
}

func (br *BookRepository) UpdateBook(tx *sql.Tx, book *model.Books) *models.INVError {
	// Create the query
	stmt := table.Books.UPDATE(
		table.Books.Isbn,
		table.Books.Author,
		table.Books.Publisher,
		table.Books.Edition,
	).SET(
		book.Isbn,
		book.Author,
		book.Publisher,
		book.Edition,
	).WHERE(
		table.Books.ItemID.EQ(mysql.String(book.ItemID)),
	)

	// Execute the query
	rows, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND
	}

	return nil
}

func (br *BookRepository) DeleteBook(tx *sql.Tx, bookId *uuid.UUID) *models.INVError {
	// TODO To be implemented
	return nil
}
