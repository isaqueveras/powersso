package query

import (
	"errors"
	"reflect"
	"strings"
)

const (
	tagSQL    string = "sql"
	ignoreTag string = "ignore"
)

var (
	errInvalidInput error = errors.New("invalid input")
	errNoneTagFound error = errors.New("no tag found")
	errFieldNull    error = errors.New("field null")
)

// FormatValuesInUp format columns and values ​​for an insert or update query
func FormatValuesInUp(input interface{}, validateNullField ...bool) (cols []string, vals []interface{}, err error) {
	var (
		elements    = reflect.ValueOf(input).Elem()
		validateNil = false
		tagFound    = false
	)

	if len(validateNullField) == 0 {
		validateNil = true
	}

	if elements.Kind() != reflect.Struct {
		return nil, nil, errInvalidInput
	}

	for i := 0; i < elements.NumField(); i++ {
		elementField := elements.Type().Field(i)

		var tag string
		if tag = elementField.Tag.Get(ignoreTag); tag != "" {
			continue
		}

		if tag = elementField.Tag.Get(tagSQL); tag == "" {
			continue
		}

		if validateNil {
			if !elements.Field(i).IsNil() {
				cols = append(cols, strings.Split(tag, "::")[0])
				vals = append(vals, elements.Field(i).Interface())
				tagFound = true
			}
		} else {
			cols = append(cols, strings.Split(tag, "::")[0])
			vals = append(vals, elements.Field(i).Interface())
			tagFound = true
		}
	}

	if !tagFound {
		return nil, nil, errNoneTagFound
	}

	if len(cols) == 0 || len(vals) == 0 {
		return nil, nil, errFieldNull
	}

	return cols, vals, nil
}
