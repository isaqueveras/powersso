// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/internal/utils"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

func TestAuthInfrastructure(t *testing.T) {
	suite.Run(t, new(authSuite))
}

type authSuite struct {
	pg   *pgAuth
	mock sqlmock.Sqlmock
	ctx  context.Context

	suite.Suite
}

func (a *authSuite) SetupTest() {
	a.pg = new(pgAuth)
	a.ctx = context.Background()

	var err error
	if a.mock, err = postgres.OpenConnectionsForTests(); err != nil {
		a.Assert().FailNow(err.Error())
	}
}

func (a *authSuite) SetupSuite() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (a *authSuite) TearDownTest() {
	postgres.CloseConnections()
}

func (a *authSuite) TestShouldCreateUser() {
	var (
		err    error
		userID *string
		input  = &auth.Register{
			FirstName: utils.GetStringPointer("Ayrton"),
			LastName:  utils.GetStringPointer("Senna"),
			Email:     utils.GetStringPointer("ayrton.senna@powersso.io"),
			Password:  utils.GetStringPointer("$2a$12$7scJnkljH5misH./.qM0YeZi7sFEU4nu4fHqOtMqHbi/p5MmzIxpG"),
			Roles:     utils.GetStringPointer("[read:activation_token]"),
			City:      utils.GetStringPointer("São Paulo, SP"),
			Country:   utils.GetStringPointer("BR"),
		}
	)

	a.mock.ExpectBegin()
	a.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users (first_name,last_name,email,password,roles,city,country) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).
		WithArgs(input.FirstName, input.LastName, input.Email, input.Password, input.Roles, input.City, input.Country).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("9f4a65cf-099b-4ea6-b091-36a9c06ecc74"))

	a.pg.DB, err = postgres.NewTransaction(a.ctx, false)
	a.Require().NotNil(a.pg.DB)
	a.Require().Nil(err, oops.Err(err))

	userID, err = a.pg.register(input)
	a.Require().NotNil(userID)
	a.Require().Nil(err, oops.Err(err))
}
