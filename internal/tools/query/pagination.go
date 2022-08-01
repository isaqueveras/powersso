package query

import (
	"database/sql"
	"reflect"

	"github.com/Masterminds/squirrel"
)

// Params models of data from params
type Params struct {
	Limit  uint64
	Offset uint64
}

// Pagination constructs pagination data for a given database query
func Pagination(
	values []any,
	model interface{},
	query *squirrel.SelectBuilder,
	params *Params,
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
		Limit(params.Limit + 1).
		Offset(params.Offset).
		Query(); err != nil {
		return nil, nil, err
	}

	slice = reflect.MakeSlice(reflect.SliceOf(modelType), 0, int(params.Limit+1))
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, nil, err
		}
		slice = reflect.Append(slice, reflect.Indirect(reflect.ValueOf(model)))
	}

	hasNext := slice.Len() > int(params.Limit)
	if hasNext {
		slice = slice.Slice(0, int(params.Limit))
	}

	next = &hasNext
	result = slice.Interface()

	return
}
