package permission

import (
	"time"

	"github.com/google/uuid"
)

// PermissionsRes models ...
type PermissionsRes struct {
	OrganizationID *uuid.UUID `json:"organization_id"`
	Permission     *[]string  `json:"permission"`
	DateCache      *time.Time `json:"date_cache,omitempty"`
}

// Permission ...
type Permission struct {
	Name        *string    `json:"name"`
	Credential  *string    `json:"credential"`
	CreatedByID *uuid.UUID `json:"-"`
}
