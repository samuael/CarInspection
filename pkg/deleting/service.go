package deleting

import (
	"errors"
	"fmt"
	"os"
)

// Service provides Post adding operations.
type Service interface {
	DeleteInspection(uint) error
	DeleteInspector(uint) error
}

// Repository provides access to Post repository.
type Repository interface {
	DeleteInspection(uint) error
	DeleteInspector(uint) error
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

// DeleteInspector ...
func (s *service) DeleteInspector(id uint) error {
	if id <= 0 {
		return errors.New(os.Getenv("INVALID_USER_ID"))
	}
	return s.dR.DeleteInspector(id)
}
