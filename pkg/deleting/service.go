package deleting

import (
	"fmt"
)

// Service provides Post adding operations.
type Service interface {
	DeleteInspection(uint) error
}

// Repository provides access to Post repository.
type Repository interface {
	DeleteInspection(uint) error
}

type service struct {
	dR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddPost adds the given Post to the database
func (s *service) DeleteInspection(u uint) error {
	if u == 0 {
		return fmt.Errorf("invalid post id")
	}
	return s.dR.DeleteInspection(u)
}
