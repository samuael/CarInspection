package adding

import (
	"fmt"
)

// ErrDuplicate defines the error message
var ErrDuplicate = fmt.Errorf("Post already exist")

// Service provides Post adding operations.
type Service interface {
	AddInspection(Inspection) error
}

// Repository provides access to Post repository.
type Repository interface {
	AddInspection(Inspection) error
}

type service struct {
	tR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddPost adds the given Post to the database
func (s *service) AddInspection(u Inspection) error {
	if u.AuthorID == 0 || u.Content == "" {
		return fmt.Errorf("invalid input")
	}
	return s.tR.AddInspection(u)
}
