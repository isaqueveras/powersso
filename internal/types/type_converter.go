package types

import "encoding/json"

// TypeConverter converts all data to a destination data output.
func TypeConverter[T any](data any) (*T, error) {
	var ( 
		result 	T
		b 			[]byte
		err error
	)

	if b, err = json.Marshal(&data); err != nil {
		return nil, err
	}
	
	if err = json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	
	return &result, err
}