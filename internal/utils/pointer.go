// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"time"

	"github.com/google/uuid"
)

// GetStringPointer returns a pointer reference
func GetStringPointer(value string) *string {
	return &value
}

// GetIntPointer returns a pointer reference
func GetIntPointer(value int) *int {
	return &value
}

// GetTimePointer returns a pointer reference
func GetTimePointer(value time.Time) *time.Time {
	return &value
}

// GetUUIDPointer returns a pointer reference
func GetUUIDPointer(value uuid.UUID) *uuid.UUID {
	return &value
}
