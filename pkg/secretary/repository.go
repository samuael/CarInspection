// package secretary
package secretary

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// This interface lists the methods that the secretary repo should implement.
type ISecretaryRepo interface {
	CreateSecretary(ctx context.Context) (secretary *model.Secretary, er error)
	DoesThisEmailExist(ctx context.Context) bool
	SecretaryByEmail(context.Context) (*model.Secretary, error)
}
