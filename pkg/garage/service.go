package garage

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type IGarageService interface {
	GetGarageByID(ctx context.Context) (*model.Garage, error)
}

type GarageService struct {
	Repo IGarageRepo
}

func NewGarageService(repo IGarageRepo) IGarageService {
	return &GarageService{
		Repo: repo,
	}
}

// GetGarageByID (ctx context.Context) (*model.Garage ,error )
func (gser *GarageService) GetGarageByID(ctx context.Context) (*model.Garage, error) {
	return gser.Repo.GetGarageByID(ctx)
}
