package pgx_storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

// AdminRepo ...
type AdminRepo struct {
	DB *pgx.Conn
}

// NewAdminRepo returning admin repo implementing al  the 
// interfaces specified by the cruds 
func NewAdminRepo( db  *pgx.Conn ) *AdminRepo {
	return &AdminRepo{
		DB : db ,
	}
}

func (adminr *AdminRepo) AdminLogin( ctx context.Context ) ( *model.Admin ,  error ) {
	var admin *model.Admin
	admin = &model.Admin{}
	println(ctx.Value("email").(string),ctx.Value("password" ).(string))
	row := adminr.DB.QueryRow(ctx , "SELECT * FROM admins WHERE email=$1 AND password=$2" , ctx.Value("email").(string) , ctx.Value("password" ).(string) )
	if row == nil {
		print("Error Finding admin ...")
		return nil  , errors.New("Error Admin Not Found ")
	}
	if err := row.Scan( admin.ID , admin.Email , admin.Firstname , admin.Middlename , admin.Lastname , admin.InspectorsCount ); err ==nil {
		return admin , nil
	}else {
		println(err.Error())
		return nil  , err 
	}
}

// func (adminr *AdminRepo) DoesAdminWithEmailExist( ctx context.Context , email string ) bool{
// 	row := adminr.DB.QueryRow(ctx , "SELECT id FROM admins WHERE EMAIL=")
// }