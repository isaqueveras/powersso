// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/authorization/permission"
	"github.com/isaqueveras/powersso/oops"
)

type Database struct {
	DB *postgres.Transaction
}

// New ...
func New(tx *postgres.Transaction) *Database {
	return &Database{DB: tx}
}

// Get ...
func (pg *Database) Get(ctx context.Context, userID, organizationID *uuid.UUID) (permissions []*string, err error) {
	var rows *sql.Rows
	if rows, err = pg.DB.Builder.
		Select("credential").
		From(`public."permission"`).
		Where(`id in (
			SELECT UNNEST(permissions_id)
			FROM public.box
			WHERE organization_id = ?
			AND ARRAY_TO_STRING(users , ',') = ?
		)`, organizationID, userID).
		Query(); err != nil {
		return nil, oops.Err(err)
	}

	for rows.Next() {
		var permission *string
		if err = rows.Scan(&permission); err != nil {
			if err == sql.ErrNoRows {
				return permissions, nil
			}
			return nil, oops.Err(err)
		}
		permissions = append(permissions, permission)
	}

	return
}

// Create ...
func (db *Database) Create(ctx context.Context, in *permission.Permission) error {
	if _, err := db.DB.Builder.
		Insert("permission").
		Columns("name", "credential", "created_by").
		Values(in.Name, in.Credential, in.CreatedByID).
		Exec(); err != nil {
		return oops.Err(err)
	}
	return nil
}
