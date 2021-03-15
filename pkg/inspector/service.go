package inspector

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type IInspectorService interface {
	CreateInspector(context.Context) (*model.Inspector, error)
	DoesThisEmailExist(context.Context) bool
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
