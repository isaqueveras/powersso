// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

// Pointer returns a pointer reference
func Pointer[T any](value T) *T {
	return &value
}
