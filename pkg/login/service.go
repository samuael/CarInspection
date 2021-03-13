package login

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	// "github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// Service ...
type Service interface {
	AdminByEmail( ctx context.Context ) ( *model.Admin ,  error ) 
}

// Repository ...
type Repository interface {
	AdminByEmail( ctx context.Context ) ( *model.Admin ,  error ) 
}

type service struct {
	lR Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{r}
}

// Login ...
func (s *service) AdminByEmail( ctx context.Context ) ( *model.Admin ,  error ) {
	// input validation
	return s.lR.AdminByEmail(ctx)
}
