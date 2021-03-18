package inspection

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type IInspectionRepo interface {
	CreateInspection(context.Context) (*model.Inspection, error)
	IsInspectionOwner(ctx context.Context) (bool, error)
	GetInspectionByID(ctx context.Context) (*model.Inspection, error)
	UpdateInspection(ctx context.Context) (*model.Inspection, error)
	DeleteInspection(ctx context.Context) (bool, error)
	DoesThisVahicheWithLicensePlateExist(ctx context.Context) bool 
	DoesThisVehicleWithVinNumberExists(ctx context.Context) bool
	SearchInspection(ctx context.Context) ([]*model.Inspection , error )
}
