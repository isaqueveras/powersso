// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import "time"

// CreateProject models the data to create a project
type CreateProject struct {
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
