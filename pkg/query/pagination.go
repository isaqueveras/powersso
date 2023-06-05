package query

import (
	"database/sql"
	"reflect"

	"github.com/Masterminds/squirrel"
	"github.com/isaqueveras/powersso/pkg/params"
)

// MakePagination constructs pagination data for a given database query
func MakePagination[T any](query *squirrel.SelectBuilder, p *params.Params) (res []T, next *bool, err error) {
	var (
		model   T
		rows    *sql.Rows
		columns []any

		modelType = reflect.Indirect(reflect.ValueOf(&model)).Type()
		slice     = reflect.MakeSlice(reflect.SliceOf(modelType), 0, int(p.Limit+1))
		modelElem = reflect.ValueOf(&model).Elem()
	)

	cols, _, err := FormatValuesInUp(&model, false)
	if err != nil {
		return nil, nil, err
	}

	if rows, err = query.
		Columns(cols...).
		Limit(p.Limit + 1).
		Offset(p.Offset).
		Query(); err != nil {
		return res, next, err
	}

	for i := 0; i < modelElem.NumField(); i++ {
		if modelElem.Type().Field(i).Tag.Get(ignoreTag) != "" {
			continue
		}

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

	hasNext := slice.Len() > int(p.Limit)
	if hasNext {
		slice = slice.Slice(0, int(p.Limit))
	}

	next = &hasNext
	res = slice.Interface().([]T)

	return
}
