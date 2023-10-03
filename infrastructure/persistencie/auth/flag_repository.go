// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/auth/postgres"
)

var _ domain.IFlag = (*repoFlag)(nil)

type repoFlag struct{ pg *infra.PGFlag }

func NewFlagRepo(tx *postgres.Transaction) domain.IFlag {
	return &repoFlag{pg: &infra.PGFlag{DB: tx}}
}

func (r *repoFlag) Set(userID *uuid.UUID, flag domain.Flag) error {
	return r.pg.Set(userID, flag)
}

func (r *repoFlag) Get(userID *uuid.UUID) (*int64, error) {
	return r.pg.Get(userID)
}
