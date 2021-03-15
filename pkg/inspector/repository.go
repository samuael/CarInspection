package inspector

import (
	"context"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)


type IInspectorRepo interface {
	CreateInspector(context.Context )  (*model.Inspector , error)
	DoesThisEmailExist(context.Context) bool 
}