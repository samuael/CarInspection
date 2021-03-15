package admin

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// Interfaces to be implemented by the admin service instances
type IAdminService interface {
	AdminByEmail(ctx context.Context) (*model.Admin , error )
}

// AdminService struct representing a admin service
type AdminService struct {
	Repo IAdminRepo
}

// NewAdminService function returninng an admin service  instance
func NewAdminService(repo IAdminRepo) IAdminService {
	return &AdminService{
		Repo: repo,
	}
}

func (adminser *AdminService) AdminByEmail(ctx context.Context) (*model.Admin  , error ) {
	return adminser.Repo.AdminByEmail(ctx)
}
