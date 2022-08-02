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
func Pagination[T any](query *squirrel.SelectBuilder, params *Params) (res []T, next *bool, err error) {
	var (
		model   T
		rows    *sql.Rows
		columns []any

		modelType = reflect.Indirect(reflect.ValueOf(&model)).Type()
		slice     = reflect.MakeSlice(reflect.SliceOf(modelType), 0, int(params.Limit+1))
		modelElem = reflect.ValueOf(&model).Elem()
	)

	if rows, err = query.
		Limit(params.Limit + 1).
		Offset(params.Offset).
		Query(); err != nil {
		return res, next, err
	}

	for i := 0; i < modelElem.NumField(); i++ {
		pt := reflect.New(reflect.PtrTo(modelElem.Field(i).Type()))
		pt.Elem().Set(modelElem.Field(i).Addr())
		columns = append(columns, pt.Elem().Interface())
	}

	for rows.Next() {
		if err = rows.Scan(columns...); err != nil {
			return res, next, err
		}
		slice = reflect.Append(slice, reflect.Indirect(reflect.ValueOf(&model)))
	}

	hasNext := slice.Len() > int(params.Limit)
	if hasNext {
		slice = slice.Slice(0, int(params.Limit))
	}

	next = &hasNext
	res = slice.Interface().([]T)

	return
}
