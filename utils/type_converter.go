// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package utils

import "encoding/json"

// TypeConverter converts all data to a destination data output.
func TypeConverter[T any](data any) (*T, error) {
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	var result T
	if err = json.Unmarshal(b, &result); err != nil {
		return nil, err
	}

	return &result, err
}
