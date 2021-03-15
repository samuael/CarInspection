package secretary

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type ISecretaryService interface {
	CreateSecretary(ctx context.Context) (secretary *model.Secretary , err error )
	DoesThisEmailExist(ctx context.Context) bool
}

// SecretaryService ...
type SecretaryService struct {
	Repo ISecretaryRepo
}

func NewSecretaryService(repo ISecretaryRepo) ISecretaryService {
	return &SecretaryService{
		Repo: repo,
	}
}


func (secretser *SecretaryService)   CreateSecretary(ctx context.Context) (secretary *model.Secretary , err error ) {
	return  secretser.Repo.CreateSecretary(ctx)
	
}

func (secretser *SecretaryService)  DoesThisEmailExist(ctx context.Context) bool {
	return secretser.Repo.DoesThisEmailExist(ctx)
}

