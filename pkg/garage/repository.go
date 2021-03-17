package garage

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)


type IGarageRepo interface {
	GetGarageByID(ctx context.Context) (*model.Garage ,error )

}
