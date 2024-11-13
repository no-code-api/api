package postgre

import (
	"context"
	"errors"
	"fmt"

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
	fmt.Println("Result: ", result)
	return result, nil
}

// func (c *Connection) Save(data interface{}, exists bool) (ok bool) {
// 	var result *gorm.DB
// 	if !exists {
// 		result = c.db.Create(data)
// 	} else {
// 		result = c.db.Updates(data)
// 	}
// 	if result.Error != nil {
// 		logger.ErrorF("error to save data: %v", result.Error.Error())
// 	}
// 	return result.Error == nil
// }

// func (c *Connection) FindQuery(dest interface{}, query string, conds ...interface{}) (ok bool) {
// 	var result *gorm.DB
// 	if query != "" {
// 		result = c.db.Where(query, conds...).Find(dest)
// 	} else {
// 		result = c.db.Where(conds).Find(dest)
// 	}
// 	fmt.Println("Query: ", result.Statement.SQL)
// 	if result.Error != nil {
// 		logger.ErrorF("error to search data: filters: %v error: %v", conds, result.Error.Error())
// 		return false
// 	}
// 	return true
// }

// func (c *Connection) Find(dest interface{}, conds interface{}) (ok bool) {
// 	var result *gorm.DB
// 	if conds == nil {
// 		result = c.db.Find(dest)
// 	} else {
// 		result = c.db.Where(conds).Find(dest)
// 	}
// 	if result.Error != nil {
// 		logger.ErrorF("error to search data: filters: %v error: %v", conds, result.Error.Error())
// 		return false
// 	}
// 	return true
// }

// func (c *Connection) Delete(data interface{}, conds ...interface{}) (ok bool) {
// 	result := c.db.Delete(data, conds...)
// 	if result.Error != nil {
// 		logger.ErrorF("error to delete data: filters: %v error: %v", conds, result.Error.Error())
// 	}
// 	return result.Error == nil
// }
