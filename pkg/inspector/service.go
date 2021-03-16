package inspector

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type IInspectorService interface {
	CreateInspector(context.Context) (*model.Inspector, error)
	DoesThisEmailExist(context.Context) bool
	InspectorByEmail(context.Context) (*model.Inspector, error)
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
