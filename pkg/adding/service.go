package adding

import (
	"fmt"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// ErrDuplicate defines the error message
var ErrDuplicate = fmt.Errorf("Post already exist")

// Service provides Post adding operations.
type Service interface {
	AddInspection(model.Inspection) error
}

// Repository provides access to Post repository.
type Repository interface {
	AddInspection(model.Inspection) error
}

type service struct {
	tR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
// AddPost adds the given Post to the database
func (s *service) AddInspection(u model.Inspection) error {
	return nil 
}
