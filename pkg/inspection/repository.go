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
}
