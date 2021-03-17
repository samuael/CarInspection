package inspector

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type IInspectorRepo interface {
	CreateInspector(context.Context) (*model.Inspector, error)
	DoesThisEmailExist(context.Context) bool
	InspectorByEmail(context.Context) (*model.Inspector, error)
	ChangePassword(ctx context.Context) (bool, error)
	GetInspectionsByInspectorID(ctx context.Context) ([]*model.Inspection, error)
	GetInspectorID(ctx context.Context) (*model.Inspector, error)
}
