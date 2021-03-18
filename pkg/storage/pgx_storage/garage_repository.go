package pgx_storage

import (
	"context"

	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/garage"
)

// GarageRepository ...
type GarageRepository struct {
	DB *pgxpool.Pool
}

// NewGarageRepo returning a repository
func NewGarageRepo(conn *pgxpool.Pool) garage.IGarageService {
	return &GarageRepository{
		DB: conn,
	}
}

// GetGarageByID (ctx context.Context) (*model.Garage, error)
func (gararepo *GarageRepository) GetGarageByID(ctx context.Context) (*model.Garage, error) {
	garageID := ctx.Value("garage_id").(uint)
	garage := &model.Garage{}
	garage.Address = &model.Address{}
	println(garageID)
	val := `SELECT addresses.id,addresses.country,addresses.region,
	addresses.zone,addresses.woreda,addresses.city,addresses.kebele,
	garage.id , garage.name 
	FROM garage
	INNER JOIN addresses ON garage.id=addresses.id  WHERE garage.id=$1 ;`;
	era := gararepo.DB.QueryRow(ctx, val, garageID).Scan( &(garage.ID), &(garage.Address.Country), &(garage.Address.Region), &(garage.Address.Zone), &(garage.Address.Woreda),
		&(garage.Address.City), &(garage.Address.Kebele), &(garage.ID), &(garage.Name))
	if era != nil {
		println(era.Error())
		return nil, era
	}
	garage.Address = garage.Address
	return garage, nil
}
