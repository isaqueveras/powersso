// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package roles

import (
	"log"
	"time"

	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// pgRoles is the implementation
// of transaction for the roles repository
type pgRoles struct {
	DB *postgres.DBTransaction
}

// removeRoles remove roles from user in database
func (pg *pgRoles) removeRoles(userID *string, roles string) (err error) {
	log.Println(roles)

	if _, err = pg.DB.Execute("UPDATE users SET roles = array_remove(roles, $1), updated_at = $2 WHERE id = $3",
		roles, time.Now(), userID); err != nil {
		return oops.Err(err)
	}

	// if err = pg.DB.Builder.
	// 	Update("users").
	// 	Set("roles = array_remove(roles, ?)", roles).
	// 	Set("updated_at = ?", time.Now()).
	// 	Where("id = ?", userID).
	// 	Suffix("RETURNING id").
	// 	Scan(new(string)); err != nil {
	// 	return oops.Err(err)
	// }

	return
}
