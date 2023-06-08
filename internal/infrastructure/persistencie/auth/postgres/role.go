package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/internal/domain/auth"
	"github.com/isaqueveras/powersso/pkg/database/postgres"
	"github.com/isaqueveras/powersso/pkg/oops"
)

// PGSession is the implementation of transaction for the role repository
type PGRole struct {
	DB *postgres.Transaction
}

// Set put the flag value in the database
func (pg *PGRole) Set(userID *uuid.UUID, flag *auth.Flag) error {
	if _, err := pg.DB.Builder.
		Update("users").
		Set("flag", flag).
		Where("id = ?::UUID", userID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}
	return nil
}
