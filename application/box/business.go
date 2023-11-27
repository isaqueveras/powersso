package box

import (
	"context"

	"github.com/google/uuid"
)

// GetMyBox ...
func GetMyBox(ctx context.Context, uid *uuid.UUID) (res *MyBoxesRes, err error) {
	res = new(MyBoxesRes)
	return
}
