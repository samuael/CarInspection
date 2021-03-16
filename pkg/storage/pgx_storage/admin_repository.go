package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/admin"
)

// AdminRepo ...
type AdminRepo struct {
	DB *pgx.Conn
}

// NewAdminRepo returning admin repo implementing al  the
// interfaces specified by the cruds
func NewAdminRepo(db *pgx.Conn) admin.IAdminRepo {
	return &AdminRepo{
		DB: db,
	}
}
func (adminr *AdminRepo) AdminByEmail(ctx context.Context) (*model.Admin, error) {
	var admin *model.Admin
	admin = &model.Admin{}
	if ctx.Value("email") ==nil {
			return nil  , errors.New(" Invalid Input ")
	}
	row := adminr.DB.QueryRow(ctx, "SELECT * FROM admins WHERE email=$1 ", ctx.Value("email").(string))
	if row == nil {
		print("Error Finding admin ...")
		return nil, errors.New("Error Admin Not Found ")
	}
	if err := row.Scan(&(admin.ID), &(admin.Email), &(admin.Firstname), &(admin.Middlename), &(admin.Lastname), &(admin.Password), &(admin.GarageID),&(admin.InspectorsCount)); err == nil {
		return admin, nil
	} else {
		return nil, err
	}
}
// func (adminr *AdminRepo) DoesAdminWithEmailExist( ctx context.Context , email string ) bool{
// 	row := adminr.DB.QueryRow(ctx , "SELECT id FROM admins WHERE EMAIL=")
// }
