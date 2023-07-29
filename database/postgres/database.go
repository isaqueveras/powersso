// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Masterminds/squirrel"

	"github.com/isaqueveras/powersso/config"
)

type database interface {
	// transaction uses a transaction from a connection already opened in the database
	transaction(ctx context.Context, readOnly bool) (interface{}, error)
	// Open open connection with database
	open(cfg *config.Config) error
	// close connection with database
	close()
	// openConnectionsForTests opens connections to the mocked database
	openConnectionsForTests() (sqlmock.Sqlmock, error)
}

var connection database

// Transaction used to aggregate transactions
type Transaction struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
}

// OpenConnections open connections with database
func OpenConnections(c *config.Config) {
	if connection != nil {
		return
	}

	connection = &postgres{}
	if err := connection.open(c); err != nil {
		log.Fatal("Unable to open connections to database: ", err)
	}
}

// OpenConnectionsForTests opens connections to the mocked database
func OpenConnectionsForTests() (mock sqlmock.Sqlmock, err error) {
	connection = &postgres{}
	if mock, err = connection.openConnectionsForTests(); err != nil {
		return nil, err
	}

	return
}

// CloseConnections close all connections with database
func CloseConnections() {
	connection.close()
}

// NewTransaction uses a transaction from a connection already opened in the database
func NewTransaction(ctx context.Context, readOnly bool) (*Transaction, error) {
	tx := &Transaction{}

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
func (t *Transaction) Commit() (erro error) {
	return t.postgres.Commit()
}

// Rollback close all pending transaction for all open databases
func (t *Transaction) Rollback() {
	_ = t.postgres.Rollback()
}

// Query executes a query that returns rows, typically a SELECT.
func (t *Transaction) Query(query string, args ...any) (*sql.Rows, error) {
	return t.postgres.Query(query, args...)
}

// Execute executes a query that doesn't return rows, typically an INSERT/UPDATE/DELETE.
func (t *Transaction) Execute(query string, args ...any) (sql.Result, error) {
	return t.postgres.Exec(query, args...)
}
