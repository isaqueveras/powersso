// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"

	"github.com/isaqueveras/powersso/config"
)

type postgres struct {
	db      *sql.DB
	timeout int
}

// open open a transaction with the database
func (p *postgres) open(c *config.Config) (err error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Name,
		c.Database.Password,
	)

	driverConfig := stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			RuntimeParams: map[string]string{
				"application_name": "power-sso",
				"DateStyle":        "ISO",
				"IntervalStyle":    "iso_8601",
				"search_path":      "public",
			},
		},
	}

	stdlib.RegisterDriverConfig(&driverConfig)
	db, err := sql.Open("pgx", driverConfig.ConnectionString(dataSourceName))
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(c.Database.MaxOpenConns)
	db.SetConnMaxLifetime(c.Database.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(c.Database.MaxIdleConns)
	db.SetConnMaxIdleTime(c.Database.ConnMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return err
	}

	p.db = db
	p.timeout = int(c.Database.Timeout)

	return nil
}

// close close the connections with database
func (p *postgres) close() {
	if p.db != nil {
		_ = p.db.Close()
	}
}

// openConnectionsForTests opens connections to the mocked database
func (p *postgres) openConnectionsForTests() (mock sqlmock.Sqlmock, err error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, fmt.Errorf("an error '%v' was not expected when opening a stub database connection", err.Error())
	}

	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(5 * time.Second)
	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(5 * time.Second)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	p.db = db
	p.timeout = 5

	return
}

// transaction opens a transaction on some already open connection
func (p *postgres) transaction(ctx context.Context, readOnly bool) (interface{}, error) {
	var (
		tx  *sql.Tx
		err error
	)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-time.After(time.Duration(p.timeout+1) * time.Second)
		if tx == nil {
			cancel()
		}
	}()

	if tx, err = p.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  readOnly,
	}); err != nil {
		return nil, err
	}

	return tx, nil
}
