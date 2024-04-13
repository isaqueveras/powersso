// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package box

import (
	"context"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/utils"
)

// PGBox ...
type repository struct {
	DB  *postgres.Transaction
	ctx context.Context
}

// Create
func (r *repository) Create() (id *uuid.UUID, err error) {
	// r.DB.Builder.Insert()
	return utils.Pointer(uuid.New()), nil
}
