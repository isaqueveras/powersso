package postgres

import (
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/oops"
)

type PGFlag struct{ DB *postgres.Transaction }

func (pg *PGFlag) Set(userID *uuid.UUID, flag auth.Flag) error {
	if _, err := pg.DB.Builder.
		Update("users").
		Set("flag", flag).
		Where("id = ?::UUID", userID).
		Exec(); err != nil {
		return oops.Err(err)
	}
	return nil
}

func (pg *PGFlag) Get(userID *uuid.UUID) (flag *int64, err error) {
	if err = pg.DB.Builder.
		Select("flag").
		From("users").
		Where("id = ?::UUID", userID).
		Scan(&flag); err != nil {
		return nil, oops.Err(err)
	}
	return
}
