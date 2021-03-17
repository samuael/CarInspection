package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/Project/CarInspection/pkg/admin"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// AdminRepo ...
type AdminRepo struct {
	DB *pgxpool.Pool
}

// NewAdminRepo returning admin repo implementing al  the
// interfaces specified by the cruds
func NewAdminRepo(db *pgxpool.Pool) admin.IAdminRepo {
	return &AdminRepo{
		DB: db,
	}
}
func (adminr *AdminRepo) AdminByEmail(ctx context.Context) (*model.Admin, error) {
	var admin *model.Admin
	admin = &model.Admin{}
	if ctx.Value("email") == nil {
		return nil, errors.New(" Invalid Input ")
	}
	row := adminr.DB.QueryRow(ctx, "SELECT * FROM admins WHERE email=$1 ", ctx.Value("email").(string))
	if row == nil {
		print("Error Finding admin ...")
		return nil, errors.New("Error Admin Not Found ")
	}
	if err := row.Scan(&(admin.ID), &(admin.Email), &(admin.Firstname), &(admin.Middlename), &(admin.Lastname), &(admin.Password), &(admin.GarageID), &(admin.InspectorsCount)); err == nil {
		return admin, nil
	} else {
		return nil, err
	}
}

// func (adminr *AdminRepo) DoesAdminWithEmailExist( ctx context.Context , email string ) bool{
// 	row := adminr.DB.QueryRow(ctx , "SELECT id FROM admins WHERE EMAIL=")
// }

// ChangePassword (ctx context.Context) (bool, error)
func (adminr *AdminRepo) ChangePassword(ctx context.Context) (bool, error) {
	id := ctx.Value("user_id").(uint)
	password := ctx.Value("password").(string)
	cmd, err := adminr.DB.Exec(ctx, "UPDATE admins SET password =$1 WHERE id=$2", password, id)
	if err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
