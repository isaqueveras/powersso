package permission

import (
	"context"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/authorization/permission"
	"github.com/isaqueveras/powersso/infrastructure/persistencie/authorization/permission"
	"github.com/isaqueveras/powersso/oops"
)

// Create ...
func Create(ctx context.Context, in *Permission) error {
	tx, err := postgres.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repo := permission.NewRepository(ctx, tx)
	if err := repo.Create(&domain.Permission{
		Name:        in.Name,
		Credential:  in.Credential,
		CreatedByID: in.CreatedByID,
	}); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}

// Get ...
func Get(ctx context.Context, userID, organizationID *uuid.UUID) (permissions []*string, err error) {
	tx, err := postgres.NewTransaction(ctx, true)
	if err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	repo := permission.NewRepository(ctx, tx)
	if permissions, err = repo.Get(userID, organizationID); err != nil {
		return nil, oops.Err(err)
	}

	return permissions, nil
}
