// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth_test

import (
	"database/sql"
	"testing"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
	"github.com/isaqueveras/power-sso/internal/domain/auth"
)
 
func TestAuth(t *testing.T) {
	suite.Run(t, new(authSuite))
}

type authSuite struct {
	pg *pgAuth
	mock sqlmock.sqlmock

	suite.Suite
}

func (f *authSuite) SetupTest() {
	f.pg = new(pgAuth)

	var err error
	if f.mock, err = postgres.OpenConnectionsForTests(); err != nil {
		f.Assert().FailNow(err.Error())
	}
}

func (f *authSuite) SetupSuite() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (f *authSuite) TearDownTest() {
	postgres.CloseConnections()
}

func (a *authSuite) TestShouldCreateUser() {
	a.mock.ExpectBegin()
	a.mock.ExpectQuery("INSERT INTO users").
		WithArgs("Ayrton, Senna, ayrton.senna@powersso.io, f8c6f60e48bc3458bc65df99325415bd").
		WillReturnResult(sqlmock.NewResult(1, 1))

	tx, err := postgres.NewTransaction(ctx, false)
	a.Require().Nil(err, oops.Err(err))
	a.Require().NotNil(tx)

	a.pg.DB = tx

	// FIXME: finalize the test implementation
	_, err = a.pg.register(&auth.Register{})
	a.Require().Nil(err, oops.Err(err))
}
