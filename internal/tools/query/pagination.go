package query

import (
	"database/sql"
	"reflect"

	"github.com/Masterminds/squirrel"
)

// MakePagination constructs pagination data for a given database query
func MakePagination(
	values []any,
	model interface{},
	query *squirrel.SelectBuilder,
	params struct {
		limit  uint64
		offset uint64
	},
) (
	result any,
	next *bool,
	err error,
) {
	var (
		modelType = reflect.Indirect(reflect.ValueOf(model)).Type()
		slice     = reflect.MakeSlice(reflect.SliceOf(modelType), 0, 0)
		rows      *sql.Rows
	)

	if rows, err = query.
		Limit(params.limit + 1).
		Offset(params.offset).
		Query(); err != nil {
		return nil, nil, err
	}

	slice = reflect.MakeSlice(reflect.SliceOf(modelType), 0, int(params.limit+1))
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, nil, err
		}
		slice = reflect.Append(slice, reflect.Indirect(reflect.ValueOf(model)))
	}

	hasNext := slice.Len() > int(params.limit)
	if hasNext {
		slice = slice.Slice(0, int(params.limit))
	}

	next = &hasNext
	result = slice.Interface()

	return
}
