// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

func TestAuth(t *testing.T) {
	suite.Run(t, new(authSuite))
}

type authSuite struct {
	pg   *pgAuth
	mock sqlmock.Sqlmock

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
		WillReturnRows()

	tx, err := postgres.NewTransaction(context.Background(), false)
	a.Require().Nil(err, oops.Err(err))
	a.Require().NotNil(tx)

	a.pg.DB = tx

	// FIXME: finalize the test implementation
	_, err = a.pg.register(&auth.Register{})
	a.Require().Nil(err, oops.Err(err))
}
