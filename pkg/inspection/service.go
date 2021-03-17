package inspection

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// IInspectionService ...
type IInspectionService interface {
	CreateInspection(context.Context) (*model.Inspection, error)
	IsInspectionOwner(ctx context.Context) (bool, error)
	GetInspectionByID(ctx context.Context) (*model.Inspection, error)
	UpdateInspection(ctx context.Context) (*model.Inspection, error)
	DeleteInspection(ctx context.Context) (bool, error)
	DoesThisVahicheWithLicensePlateExist(ctx context.Context) bool
	DoesThisVehicleWithVinNumberExists(ctx context.Context) bool
}

// InspectionService ...
type InspectionService struct {
	Repo IInspectionRepo
}

// NewInspectionService returning the IInspectionService Port to be implemented by the
// InspectionService
func NewInspectionService(repo IInspectionRepo) IInspectionService {
	return &InspectionService{
		Repo: repo,
	}
}

func (inser *InspectionService) CreateInspection(ctx context.Context) (*model.Inspection, error) {
	return inser.Repo.CreateInspection(ctx)
}

// IsInspectionOwner ( inspectionID uint  , uint inspectorID ) (bool  , error )
func (inser *InspectionService) IsInspectionOwner(ctx context.Context) (bool, error) {
	return inser.Repo.IsInspectionOwner(ctx)
}

// GetInspectionByID (ctx context.Context) (*model.Inspection, error)
// service method to call the Inspection Repositories GetInspectionBYID method
func (inser *InspectionService) GetInspectionByID(ctx context.Context) (*model.Inspection, error) {
	return inser.Repo.GetInspectionByID(ctx)
}

// UpdateInspection (ctx context.Context) (*model.Inspection, error)
func (inser *InspectionService) UpdateInspection(ctx context.Context) (*model.Inspection, error) {
	return inser.Repo.UpdateInspection(ctx)
}

// DeleteInspection (ctx context.Context) (bool, error)
func (inser *InspectionService) DeleteInspection(ctx context.Context) (bool, error) {
	return inser.Repo.DeleteInspection(ctx)
}

// DoesThisVahicheWithLicensePlateExist (ctx context.Context) bool
func (inser *InspectionService) DoesThisVahicheWithLicensePlateExist(ctx context.Context) bool {
	return inser.Repo.DoesThisVahicheWithLicensePlateExist(ctx)
}

// DoesThisVehicleWithVinNumberExists  ...
func (inser *InspectionService) DoesThisVehicleWithVinNumberExists(ctx context.Context) bool {
	return inser.Repo.DoesThisVehicleWithVinNumberExists(ctx)
}
