package postgre

import (
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Result struct {
	rows         []map[string]interface{}
	fields       []string
	currentIndex int
	rowsCount    int
}

func newResult(rows pgx.Rows) (*Result, error) {
	fieldNames := getFieldNames(rows)
	data, err := readRows(rows, fieldNames)
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &Result{
		rows:         data,
		currentIndex: -1,
		fields:       fieldNames,
		rowsCount:    len(data),
	}, nil
}

func (r *Result) Next() bool {
	r.currentIndex++
	return r.currentIndex < r.rowsCount
}

func (r *Result) GetFieldValue(field string) interface{} {
	if r.currentIndex < 0 || r.currentIndex >= len(r.rows) {
		logger.Fatal("current row is out of range")
	}
	row := r.rows[r.currentIndex]
	value, exists := row[field]
	if !exists {
		message := fmt.Sprintf("column %s not found", field)
		logger.Fatal(message)
	}
	return value
}

func (r *Result) ReadString(field string) string {
	fieldValue := r.GetFieldValue(field)
	if fieldValue == nil {
		return ""
	}
	result, ok := fieldValue.(string)
	if !ok {
		return ""
	}
	return result
}

func (r *Result) ReadInt(field string) int {
	fieldValue := r.GetFieldValue(field)
	if fieldValue == nil {
		return 0
	}
	result, ok := fieldValue.(int32)
	if !ok {
		return 0
	}
	return int(result)
}

func getFieldNames(rows pgx.Rows) []string {
	fieldDescriptions := rows.FieldDescriptions()
	fieldNames := make([]string, len(fieldDescriptions))
	for i, field := range fieldDescriptions {
		fieldNames[i] = string(field.Name)
	}
	return fieldNames
}

func readRows(rows pgx.Rows, fieldNames []string) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		row := make(map[string]interface{})
		for i, fieldName := range fieldNames {
			row[fieldName] = values[i]
		}
		data = append(data, row)
	}
	return data, nil
}
