package admin

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type IAdminRepo interface {
	AdminByEmail(ctx context.Context) (*model.Admin, error)
	ChangePassword(ctx context.Context) (bool, error)
}