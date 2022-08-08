// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/isaqueveras/power-sso/config"
)

type database interface {
	// transaction uses a transaction from a connection already opened in the database
	transaction(ctx context.Context, readOnly bool) (interface{}, error)
	// Open open connection with database
	open(cfg *config.Config) error
	// Close close connection with database
	close()
}

var connection database

// DBTransaction used to aggregate transactions
type DBTransaction struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
}

// OpenConnections open connections with database
func OpenConnections(c *config.Config) (err error) {
	if connection != nil {
		return nil
	}

	connection = &postgres{}
	if err = connection.open(c); err != nil {
		return err
	}
	return nil
}

// CloseConnections close all connections with database
func CloseConnections() {
	connection.close()
}

// NewTransaction uses a transaction from a connection already opened in the database
func NewTransaction(ctx context.Context, readOnly bool) (*DBTransaction, error) {
	tx := &DBTransaction{}

	pgsql, err := connection.transaction(ctx, readOnly)
	if err != nil {
		return nil, err
	}

	tx.postgres = pgsql.(*sql.Tx)
	tx.Builder = squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(tx.postgres)

	return tx, nil
}

// Commit confirm pending transactions for all open databases
func (t *DBTransaction) Commit() (erro error) {
	return t.postgres.Commit()
}

// Rollback close all pending transaction for all open databases
func (t *DBTransaction) Rollback() {
	_ = t.postgres.Rollback()
}
