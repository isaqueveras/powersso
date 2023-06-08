// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"errors"
	"time"

	"github.com/gosimple/slug"

	"github.com/isaqueveras/powersso/internal/utils"
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
	if cpr.Name == nil || *cpr.Name == "" {
		return errors.New("cannot create a project without a name")
	}

	cpr.Slug = utils.Pointer(slug.Make(*cpr.Name))
	if len(cpr.Participants) == 0 {
		return errors.New("cannot create a project without participants")
	}

	if cpr.Color == nil || *cpr.Color == "" {
		cpr.Color = utils.Pointer("#949494")
	}

	return
}
