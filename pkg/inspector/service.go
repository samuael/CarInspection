package inspector

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type IInspectorService interface {
	CreateInspector(context.Context) (*model.Inspector, error)
	DoesThisEmailExist(context.Context) bool
	InspectorByEmail(context.Context) (*model.Inspector, error)
	ChangePassword(ctx context.Context) (bool, error)
	GetInspectionsByInspectorID(ctx context.Context) ([]*model.Inspection, error)
	GetInspectorByID(ctx context.Context) (*model.Inspector, error)
	UpdateProfileImage(ctx context.Context) error
	DeleteInspectorByID(ctx context.Context) error
}

// InspectorService  ...
type InspectorService struct {
	Repo IInspectorRepo
}

func NewInspectorService(repo IInspectorRepo) IInspectorService {
	return &InspectorService{
		Repo: repo,
	}
}

// CreateInspector (context.Context) (*model.Inspector, error)
func (inorser InspectorService) CreateInspector(ctx context.Context) (*model.Inspector, error) {
	return inorser.Repo.CreateInspector(ctx)
}

// DoesThisEmailExist for verifying the presence of the inspector
func (inorser InspectorService) DoesThisEmailExist(ctx context.Context) bool {
	return inorser.Repo.DoesThisEmailExist(ctx)
}

// InspectorByEmail ( context.Context ) (*model.Inspector, error )
func (insorser InspectorService) InspectorByEmail(ctx context.Context) (*model.Inspector, error) {
	return insorser.Repo.InspectorByEmail(ctx)
}

// ChangePassword ... method
func (insorser *InspectorService) ChangePassword(ctx context.Context) (bool, error) {
	return insorser.Repo.ChangePassword(ctx)
}

// GetInspectionsByInspectorID ...
func (insorser InspectorService) GetInspectionsByInspectorID(ctx context.Context) ([]*model.Inspection, error) {
	return insorser.Repo.GetInspectionsByInspectorID(ctx)
}

// GetInspectoryID (ctx context.Context )  (*model.Inspection , error )
func (insorser InspectorService) GetInspectorByID(ctx context.Context) (*model.Inspector, error) {
	return insorser.Repo.GetInspectorByID(ctx)
}

// UpdateProfileImage (ctx context.Context) error
func (insorser *InspectorService) UpdateProfileImage(ctx context.Context) error {
	return insorser.Repo.UpdateProfileImage(ctx)
}

// DeleteInspectorByID (ctx context.Context) error
func (insorser *InspectorService) DeleteInspectorByID(ctx context.Context) error {
	return insorser.Repo.DeleteInspectorByID(ctx)
}
