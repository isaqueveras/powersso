package permission

import (
	"time"

	"github.com/google/uuid"
)

// PermissionsRes models ...
type PermissionsRes struct {
	ProjectID  *uuid.UUID `json:"pid"`
	Permission *uint64    `json:"permission"`
	DateCache  *time.Time `json:"date_cache,omitempty"`
}
