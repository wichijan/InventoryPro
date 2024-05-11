package utils

import (
	"database/sql"
	"time"

	inv_errors "github.com/wichijan/InventoryPro/src/errors"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
	"github.com/wichijan/InventoryPro/src/models"
)

func MysqlUuid(uuid *uuid.UUID) mysql.StringExpression {
	binary_id, _ := uuid.MarshalBinary()
	return mysql.String(string(binary_id))
}

func MysqlUuidOrNil(uuid *uuid.UUID) mysql.Expression {
	if uuid == nil {
		return mysql.NULL
	}
	binary_id, _ := uuid.MarshalBinary()
	return mysql.String(string(binary_id))
}

func MysqlTime(time *time.Time) mysql.TimeExpression {
	return mysql.TimeT(*time)
}

func MySqlString(str string) mysql.StringExpression {
	return mysql.String(str)
}

func MySqlStringPtr(str *string) mysql.StringExpression {
	if str == nil || *str == "" {
		return nil
	}
	return mysql.String(*str)
}

func MysqlTimeNow() mysql.TimestampExpression {
	return mysql.NOW()
}

func ExcecuteInsertStatement(stmt mysql.InsertStatement, dbConnection *sql.DB) *models.INVError {
	result, err := stmt.Exec(dbConnection)

	if err != nil {
		return inv_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return inv_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected != 1 {
		return inv_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func CountStatement(table mysql.Table, where mysql.BoolExpression, conn *sql.DB) (int, *models.INVError) {
	var result CountQueryResult
	stmt := mysql.SELECT(
		mysql.COUNT(mysql.STAR).AS("CountQueryResult.Count"),
	).FROM(
		table,
	).WHERE(where)

	err := stmt.Query(conn, &result)
	if err != nil {
		return 0, inv_errors.KTS_INTERNAL_ERROR
	}
	return result.Count, nil
}

type CountQueryResult struct {
	Count int
}

func GetDateTime(x time.Time) mysql.TimestampExpression {
	return mysql.DateTimeT(x)
}
