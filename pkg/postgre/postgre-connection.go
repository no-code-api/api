package postgre

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	internalLogger "github.com/leandro-d-santos/no-code-api/internal/logger"
)

var (
	connection *Connection
	logger     *internalLogger.Logger = internalLogger.NewLogger("Postgre")
)

type Connection struct {
	database *pgx.Conn
}

func newConnection(pgx *pgx.Conn) *Connection {
	return &Connection{
		database: pgx,
	}
}

func (conn *Connection) ExecuteNonQuery(query string) error {
	if _, err := conn.database.Exec(context.Background(), query); err != nil {
		logger.InfoF("error to exec query. Query: %s - Error: %s", query, err.Error())
		return errors.New("error to execute")
	}
	return nil
}

func (conn *Connection) ExecuteQuery(query string) (*Result, error) {
	rows, err := conn.database.Query(context.Background(), query)
	if err != nil {
		logger.InfoF("error to exec query. Query: %s - Error: %s", query, err.Error())
		return nil, errors.New("error to execute")
	}
	defer rows.Close()
	return newResult(rows)
}

func (conn *Connection) ExecuteSingleQuery(query string) (interface{}, error) {
	var result interface{}
	row := conn.database.QueryRow(context.Background(), query)
	err := row.Scan(&result)
	if err != nil {
		logger.InfoF("error to exec query. Query: %s - Error: %s", query, err.Error())
		return nil, errors.New("error to execute")
	}
	return result, nil
}
