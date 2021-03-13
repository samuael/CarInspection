package gorm

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// InspectionRepo ...
type InspectionRepo struct {
	DB *gorm.DB
}

// NewInspectionRepo gorm repository to handle persistent data to and from the
// database
func NewInspectionRepo(db *gorm.DB) *InspectionRepo {
	return &InspectionRepo{
		DB: db,
	}
}

// CreateInspection ... creating inspection instance  
func (inrep *InspectionRepo) CreateInspection(ctx context.Context, inspection *model.Inspection) (*model.Inspection, error) {

	return nil, nil
}
