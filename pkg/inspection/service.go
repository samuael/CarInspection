package inspection

// IInspectionService ...
type IInspectionService interface {
}

// InspectionService ...
type InspectionService struct {
	Repo IInspectionRepo
}

// NewInspectionService returning the IInspectionService Port to be implemented by the 
// InspectionService 
func NewInspectionService(repo IInspectionRepo)  IInspectionService {
	return &InspectionService{}
}

