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
	"github.com/wichijan/InventoryPro/src/utils"
)

type BookRepositoryI interface {
	GetBookById(itemId *uuid.UUID) (*model.Books, *models.INVError)
	CreateBook(tx *sql.Tx, book *model.Books) *models.INVError
	UpdateBook(tx *sql.Tx, book *model.Books) *models.INVError
	DeleteBook(tx *sql.Tx, bookId *uuid.UUID) *models.INVError
	CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError)

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
		return nil, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error reading book")
	}

	return &book, nil
}

func (br *BookRepository) CreateBook(tx *sql.Tx, book *model.Books) *models.INVError {
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
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error creating book")
	}

	return nil
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
	_, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error updating book")
	}

	return nil
}

func (br *BookRepository) DeleteBook(tx *sql.Tx, bookId *uuid.UUID) *models.INVError {
	// Create the query
	stmt := table.Books.DELETE().WHERE(
		table.Books.ItemID.EQ(mysql.String(bookId.String())),
	)

	// Execute the query
	rows, err := stmt.Exec(tx)
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error deleting book")
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return inv_errors.INV_INTERNAL_ERROR.WithDetails("Error: No changes on entry")
	}

	if rowsAff == 0 {
		return inv_errors.INV_NOT_FOUND.WithDetails("Book Id not found")
	}

	return nil
}

func (br *BookRepository) CheckIfItemIdExists(itemId *uuid.UUID) (bool, *models.INVError) {
	count, err := utils.CountStatement(table.Books, table.Books.ItemID.EQ(mysql.String(itemId.String())), br.GetDatabaseConnection())
	if err != nil {
		return false, inv_errors.INV_INTERNAL_ERROR.WithDetails("Error checking if itemId exists in Book table")
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
