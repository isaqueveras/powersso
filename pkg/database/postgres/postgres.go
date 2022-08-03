// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"

	"github.com/isaqueveras/power-sso/config"
)

type postgres struct {
	db      *sql.DB
	timeout int
}

// open open a transaction with the database
func (p *postgres) open(c *config.Config) (err error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.User,
		c.Postgres.Dbname,
		c.Postgres.Password,
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
	db, err := sql.Open(c.Postgres.Driver, driverConfig.ConnectionString(dataSourceName))
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(c.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(c.Postgres.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(c.Postgres.MaxIdleConns)
	db.SetConnMaxIdleTime(c.Postgres.ConnMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return err
	}

	p.db = db
	p.timeout = int(c.Postgres.Timeout)

	return nil
}

// close close the connections with database
func (p *postgres) close() {
	if p.db != nil {
		_ = p.db.Close()
	}
}

// transaction opens a transaction on some already open connection
func (p *postgres) transaction(ctx context.Context) (interface{}, error) {
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

	if tx, err = p.db.BeginTx(ctx, nil); err != nil {
		return nil, err
	}

	return tx, nil
}
