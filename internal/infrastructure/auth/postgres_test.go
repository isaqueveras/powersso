// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO users").
		WithArgs("Ayrton, Senna, ayrton.senna@powersso.io, f8c6f60e48bc3458bc65df99325415bd").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	var tx *sql.Tx
	if tx, err = db.Begin(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	defer tx.Rollback()

	if _, err = tx.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)",
		"Ayrton, Senna, ayrton.senna@powersso.io, f8c6f60e48bc3458bc65df99325415bd"); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if err = tx.Commit(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
