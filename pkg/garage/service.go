package garage 

type IGarageService interface {

}


type GarageService struct {
	Repo IGarageRepo
}

func NewGarageService(repo IGarageRepo) IGarageService{
	return &GarageService{
		Repo: repo,
	}
}