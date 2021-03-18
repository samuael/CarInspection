package secretary

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

type ISecretaryService interface {
	CreateSecretary(ctx context.Context) (secretary *model.Secretary, err error)
	DoesThisEmailExist(ctx context.Context) bool
	SecretaryByEmail(context.Context) (*model.Secretary, error)
	ChangePassword(ctx context.Context) (bool, error)
	GetSecretaryByID(ctx context.Context) (*model.Secretary, error)
	DeleteSecretaryByID(ctx context.Context)  error
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

func (secretser *SecretaryService) CreateSecretary(ctx context.Context) (secretary *model.Secretary, err error) {
	return secretser.Repo.CreateSecretary(ctx)

}

func (secretser *SecretaryService) DoesThisEmailExist(ctx context.Context) bool {
	return secretser.Repo.DoesThisEmailExist(ctx)
}

// SecretaryByEmail (context.Context) (*model.Secretary, error)
func (secretser *SecretaryService) SecretaryByEmail(ctx context.Context) (*model.Secretary, error) {
	return secretser.Repo.SecretaryByEmail(ctx)
}

// ChangePassword (ctx context.Context) (bool, error)
func (secretser *SecretaryService) ChangePassword(ctx context.Context) (bool, error) {
	return secretser.Repo.ChangePassword(ctx)
}

// GetSecretaryByID ...
func (secretser *SecretaryService) GetSecretaryByID(ctx context.Context) (*model.Secretary, error) {
	return secretser.Repo.GetSecretaryByID(ctx)
}

// DeleteSecretaryByID ... 
func (secretser *SecretaryService) DeleteSecretaryByID(ctx context.Context) error {
	return secretser.Repo.DeleteSecretaryByID(ctx)
}
