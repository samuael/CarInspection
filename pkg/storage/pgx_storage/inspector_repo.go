package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/inspector"
)

/*
	CreateInspector(context.Context )  (*model.Inspector , error)
	DoesThisEmailExist(context.Context) bool
*/

// InspectorRepo ... returning an Inspector Repository
type InspectorRepo struct {
	DB *pgx.Conn
}

// NewInspectorRepo returning a repository at the top of pgx postgres repository.
func NewInspectorRepo(conn *pgx.Conn) inspector.IInspectorRepo {
	return &InspectorRepo{
		DB: conn,
	}
}

// CreateSecretary method to create a secretary
func (insorrepo *InspectorRepo) CreateInspector(ctx context.Context) (inspector *model.Inspector, er error) {
	defer recover()
	// Getting the secretary value passed in with the context
	if ctx.Value("inspector") == nil {
		return nil, errors.New(" Invalid Input ")
	}
	inspector = ctx.Value("inspector").(*model.Inspector)
	conmdTag, er := insorrepo.DB.Exec(context.Background(), "insert into inspectors( email , firstname ,middlename ,lastname , password ,garageid , imageurl , createdby ) values($1 , $2 , $3 , $4 , $5 , $6 , $7 , $8 )",
		inspector.Email, inspector.Firstname, inspector.Middlename, inspector.Lastname, inspector.Password, inspector.GarageID, inspector.Imageurl, inspector.Createdby)
	if er != nil || !(conmdTag.Insert()) || conmdTag.RowsAffected() == 0 {
		return nil, er
	}
	return inspector, nil
}

func (insorrepo *InspectorRepo) DoesThisEmailExist(ctx context.Context) bool {
	defer recover()
	// getting the secretory by email
	row := insorrepo.DB.QueryRow(ctx, "SELECT id from inspectors WHERE email=$1", ctx.Value("email").(string))
	var val = 0
	era := row.Scan(&val)
	if era != nil || val == 0 {
		return false
	}
	return true
}

// InspectorByEmail the parameter context will have the email string and
// we wil be using that email string as a way of selecting the instance data
func (insorrepo *InspectorRepo) InspectorByEmail(ctx context.Context) (*model.Inspector, error) {
	defer recover()
	var inspector *model.Inspector
	inspector = &model.Inspector{}
	if ctx.Value("email") == nil {
		return nil, errors.New(" Invalid Input ")
	}
	row := insorrepo.DB.QueryRow(ctx, "SELECT * FROM inspectors WHERE email=$1 ", ctx.Value("email").(string))
	if row == nil {
		print("Error Finding admin ...")
		return nil, errors.New("Error Admin Not Found ")
	}
	if err := row.Scan(
		&(inspector.ID),
		&(inspector.Email),
		&(inspector.Firstname),
		&(inspector.Middlename),
		&(inspector.Lastname),
		&(inspector.Password),
		&(inspector.Imageurl),
		&(inspector.GarageID),
		&(inspector.Createdby),
	); err == nil {
		return inspector, nil
	} else {
		return nil, err
	}
}
