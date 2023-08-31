package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/dzemildupljak/web_app_with_unitest/generic/common"
	"github.com/lib/pq"
)

func GetRow[T any, PT common.Ptr[T]](ctx context.Context, q string, args ...any) (T, error) {
	row := database.QueryRowContext(ctx, q, args...)
	var t T
	ptr := PT(&t)
	if err := row.Scan(ptr.PtrFields()...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return t, ErrNotFound
		}
		return t, fmt.Errorf("GetRow row.Scan error\n%s: %w", q, err)
	}
	return t, nil
}

func FindRows[T any, PT common.Ptr[T]](ctx context.Context, q string, args ...any) ([]T, error) {
	rows, err := database.QueryContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("FindRows QueryContext failed\n%s: %w", q, err)
	}
	defer func() { _ = rows.Close() }()

	var result []T
	for rows.Next() {
		var t T
		ptr := PT(&t)
		if err := rows.Scan(ptr.PtrFields()...); err != nil {
			return nil, fmt.Errorf("FindRows row.Scan error\n%s: %w", q, err)
		}
		result = append(result, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindRows rows.Err()\n%s: %w", q, err)
	}
	return result, nil
}

type TOrderDir string

const (
	ASC  TOrderDir = "ASC"
	DESC TOrderDir = "DESC"
)

type QueryOptions struct {
	Filter   map[string]interface{}
	OrderBy  string
	OrderDir TOrderDir
	Limit    int
	Page     int
}

func (qo *QueryOptions) CreateQueryFromOptions() string {
	var query string
	if qo.OrderBy != "" {
		query += " ORDER BY " + qo.OrderBy
	} else {
		query += " ORDER BY id"
	}

	if qo.OrderDir != "" {
		query += " " + string(qo.OrderDir)
	} else {
		query += " DESC"
	}

	if qo.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", qo.Limit)
	} else {
		query += fmt.Sprintf(" LIMIT %d", 10)
	}

	if qo.Page > 0 {
		offset := (qo.Page - 1) * qo.Limit
		query += fmt.Sprintf(" OFFSET %d", offset)
	} else if qo.Page > 0 && qo.Limit <= 0 {
		offset := (qo.Page - 1) * 10
		query += fmt.Sprintf(" OFFSET %d", offset)
	} else {
		query += fmt.Sprintf(" OFFSET %d", 0)
	}
	return query
}

func GetValidColumns[T any, PT common.Ptr[T]]() (map[string]bool, []string) {
	validColumns := make(map[string]bool)
	columnsName := []string{}
	t := reflect.TypeOf((PT)(nil)).Elem()

	for i := 0; i < t.NumField(); i++ {
		columnName := t.Field(i).Tag.Get("column_name")
		if columnName != "" {
			validColumns[columnName] = true
			columnsName = append(columnsName, columnName)
		}
	}

	return validColumns, columnsName
}

func BuildWhereClauses(filters map[string]interface{}, startIdx int) ([]string, []interface{}) {
	var whereClauses []string
	var whereValues []interface{}

	for column, value := range filters {
		switch val := value.(type) {
		case []string:
			clause := fmt.Sprintf("%s = ANY($%d)", column, startIdx+len(whereValues)+1)
			whereClauses = append(whereClauses, clause)
			whereValues = append(whereValues, pq.Array(val))
		default:
			clause := fmt.Sprintf("%s = $%d", column, startIdx+len(whereValues)+1)
			whereClauses = append(whereClauses, clause)
			whereValues = append(whereValues, value)
		}
	}

	return whereClauses, whereValues
}
