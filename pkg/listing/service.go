package listing

// Service provides Post listing operations.
type Service interface {
	GetMyInspections(userid uint ) ([]Inspection, error)
}

// Repository provides access to Inspections repository.
type Repository interface {
	GetMyInspections( userid uint ) ([]Inspection, error)
}

type service struct {
	tR Repository
}

// NewService creates an list service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetAllPosts returns all Posts from the storage
func (s *service) GetMyInspections(userid uint ) ([]Inspection, error) {
	return s.tR.GetMyInspections(  userid )
}

//func (s *service) GetUserPosts(id uint) ([]Post, error) {
//	return s.tR.GetUserPosts(id)
//}
