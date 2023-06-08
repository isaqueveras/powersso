// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postgres

import (
	"context"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/powersso/internal/domain/auth"
	"github.com/isaqueveras/powersso/internal/utils"
	pg "github.com/isaqueveras/powersso/pkg/database/postgres"
	"github.com/isaqueveras/powersso/pkg/oops"
)

func TestAuthInfrastructure(t *testing.T) {
	suite.Run(t, new(authSuite))
}

type authSuite struct {
	pg   *PGAuth
	mock sqlmock.Sqlmock
	ctx  context.Context

	suite.Suite
}

func (a *authSuite) SetupTest() {
	a.pg = new(PGAuth)
	a.ctx = context.Background()

	var err error
	if a.mock, err = pg.OpenConnectionsForTests(); err != nil {
		a.Assert().FailNow(err.Error())
	}
}

func (a *authSuite) SetupSuite() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (a *authSuite) TearDownTest() {
	pg.CloseConnections()
}

func (a *authSuite) TestShouldCreateUser() {
	var (
		err    error
		userID *uuid.UUID
		input  = &auth.CreateAccount{
			FirstName: utils.Pointer("Ayrton"),
			LastName:  utils.Pointer("Senna"),
			Email:     utils.Pointer("ayrton.senna@powersso.io"),
			Password:  utils.Pointer("$2a$12$7scJnkljH5misH./.qM0YeZi7sFEU4nu4fHqOtMqHbi/p5MmzIxpG"),
		}
	)

	a.mock.ExpectBegin()
	a.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users (first_name,last_name,email,password) VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(input.FirstName, input.LastName, input.Email, input.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("9f4a65cf-099b-4ea6-b091-36a9c06ecc74"))

	a.pg.DB, err = pg.NewTransaction(a.ctx, false)
	a.Require().NotNil(a.pg.DB)
	a.Require().Nil(err, oops.Err(err))

	userID, err = a.pg.CreateAccount(input)
	a.Require().NotNil(userID)
	a.Require().Nil(err, oops.Err(err))
}
