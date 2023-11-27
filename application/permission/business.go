package permission

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/utils"
)

// GetPermissions ...
func GetPermissions(ctx context.Context, pid *uuid.UUID) (res *PermissionsRes, err error) {
	res = new(PermissionsRes)
	res.ProjectID = pid
	res.DateCache = utils.Pointer(time.Now())
	res.Permission = utils.Pointer(uint64(9003840938243279983))
	return res, nil
}
