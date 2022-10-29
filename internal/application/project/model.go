// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"errors"
	"time"

	"github.com/gosimple/slug"
)

// CreateProjectReq models the data to create a project
type CreateProjectReq struct {
	Name         *string       `json:"name"`
	Description  *string       `json:"description,omitempty"`
	Color        *string       `json:"hex_color,omitempty"`
	Participants []Participant `json:"participants"`
	CreatedByID  *string       `json:"created_by_id,omitempty"`
	Slug         *string       `json:"slug,omitempty"`
}

// Participant models the data the a participant
type Participant struct {
	UserID        *string    `json:"user_id"`
	StartDate     *time.Time `json:"start_date"`
	DepartureDate *time.Time `json:"departure_date,omitempty"`
}

// Validate validation of data for registration
func (cpr *CreateProjectReq) Validate() (err error) {
	slug := slug.Make(*cpr.Name)
	cpr.Slug = &slug

	if cpr.Name == nil || *cpr.Name == "" {
		err = errors.New("cannot create a project without a name")
		return
	}

	if len(cpr.Participants) == 0 {
		err = errors.New("cannot create a project without participants")
		return
	}

	if cpr.Color == nil || *cpr.Color == "" {
		defaultColor := "#949494"
		cpr.Color = &defaultColor
	}

	return
}
