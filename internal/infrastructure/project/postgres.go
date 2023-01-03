// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"github.com/isaqueveras/power-sso/internal/domain/project"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// pg is the implementation of transaction for the session repository
type pg struct {
	DB *postgres.DBTransaction
}

// Create contains the flow for create project in database
func (pg *pg) create(input *project.CreateProject) (err error) {
	var projectID *string
	if err = pg.DB.Builder.
		Insert("projects").
		Columns("created_by", "name", "description", "color", "slug").
		Values(input.CreatedByID, input.Name, input.Description, input.Color, input.Slug).
		Suffix(`RETURNING "id"`).
		Scan(&projectID); err != nil {
		return oops.Err(err)
	}

	for _, value := range input.Participants {
		if err = pg.DB.Builder.
			Insert("project_participants").
			Columns("user_id", "start_date", "departure_date", "project_id").
			Values(value.UserID, value.StartDate, value.DepartureDate, projectID).
			Suffix(`RETURNING "user_id"`).
			Scan(new(string)); err != nil {
			return oops.Err(err)
		}
	}

	return
}
