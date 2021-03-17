package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/secretary"
)

// SecretaryRepository ...
type SecretaryRepository struct {
	DB *pgxpool.Pool
}

// NewSecretaryRepo returning the repository for secretary
func NewSecretaryRepo(db *pgxpool.Pool) secretary.ISecretaryRepo {
	return &SecretaryRepository{
		DB: db,
	}
}

// CreateSecretary method to create a secretary
func (secretr *SecretaryRepository) CreateSecretary(ctx context.Context) (secretary *model.Secretary, er error) {
	defer recover()
	// Getting the secretary value passed in with the context
	if ctx.Value("secretary") == nil {
		return nil, errors.New(" Invalid Input ")
	}
	secretary = ctx.Value("secretary").(*model.Secretary)
	conmdTag, er := secretr.DB.Exec(context.Background(), "insert into secretaries( email , firstname ,middlename ,lastname , password ,garageid , createdby ) values($1 , $2 , $3 , $4 , $5 , $6 , $7 )",
		secretary.Email, secretary.Firstname, secretary.Middlename, secretary.Lastname, secretary.Password, secretary.GarageID, secretary.Createdby)
	if er != nil || !(conmdTag.Insert()) || conmdTag.RowsAffected() == 0 {
		return nil, er
	}
	return secretary, nil
}

func (secretr *SecretaryRepository) DoesThisEmailExist(ctx context.Context) bool {
	// defer recover()
	// getting the secretory by email
	row := secretr.DB.QueryRow(ctx, "SELECT id from secretaries WHERE email=$1", ctx.Value("email").(string))
	var val = 0
	era := row.Scan(&val)
	if era != nil || val == 0 {
		return false
	}
	return true
}

// SecretaryByEmail parameter context will have email value that is to be used in the  database
func (secretr *SecretaryRepository) SecretaryByEmail(ctx context.Context) (*model.Secretary, error) {
	var secretary *model.Secretary
	secretary = &model.Secretary{}
	if ctx.Value("email") == nil {
		return nil, errors.New(" Invalid Input ")
	}
	row := secretr.DB.QueryRow(ctx, "SELECT * FROM secretaries WHERE email=$1 ", ctx.Value("email").(string))
	if row == nil {
		print("Error Finding admin ...")
		return nil, errors.New("Error Admin Not Found ")
	}
	if err := row.Scan(
		&(secretary.ID),
		&(secretary.Email),
		&(secretary.Firstname),
		&(secretary.Middlename),
		&(secretary.Lastname),
		&(secretary.Password),
		&(secretary.GarageID),
		&(secretary.Createdby),
	); err == nil {
		return secretary, nil
	} else {
		return nil, err
	}
}

// ChangePassword (ctx context.Context) (bool, error)
func (secretr *SecretaryRepository) ChangePassword(ctx context.Context) (bool, error) {
	id := ctx.Value("user_id").(uint)
	password := ctx.Value("password").(string)
	cmd, err := secretr.DB.Exec(ctx, "UPDATE secretaries SET password =$1 WHERE id=$2", password, id)
	if err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
