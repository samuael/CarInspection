package pgx_storage

import (
	"context"
	"errors"

	// "github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/inspector"
)

/*
	CreateInspector(context.Context )  (*model.Inspector , error)
	DoesThisEmailExist(context.Context) bool
*/

// InspectorRepo ... returning an Inspector Repository
type InspectorRepo struct {
	DB *pgxpool.Pool
}

// NewInspectorRepo returning a repository at the top of pgx postgres repository.
func NewInspectorRepo(conn *pgxpool.Pool) inspector.IInspectorRepo {
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

// ChangePassword (ctx context.Context) (bool, error)
func (insorrepo *InspectorRepo) ChangePassword(ctx context.Context) (bool, error) {
	id := ctx.Value("user_id").(uint)
	password := ctx.Value("password").(string)
	cmd, err := insorrepo.DB.Exec(ctx, "UPDATE inspectors SET password =$1 WHERE id=$2", password, id)
	if err != nil || cmd.RowsAffected() == 0 {
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

// GetInspectionsByInspectorID   method returning the list of inspections registered by the user whose ID is specified in the
// context value
func (insorrepo *InspectorRepo) GetInspectionsByInspectorID(ctx context.Context) ([]*model.Inspection, error) {
	inspectorID := ctx.Value("inspector_id").(uint)

	inspections := []*model.Inspection{}

	// seelcting multiple inspections using the QueryRows Method of golang
	colmns, er := insorrepo.DB.Query(ctx, "SELECT * FROM  inspections WHERE inspector_id=$1", inspectorID)
	if er != nil {
		return inspections, er
	}
	for colmns.Next() {
		handbrake := 0
		steeringSystem := 0
		brakeSystem := 0
		seatBelt := 0
		doorAndWindow := 0
		dashBoardLight := 0
		windshield := 0
		baggageDoorWindow := 0
		gearBox := 0
		shockAbsorber := 0
		frontHighAndLowBeamLight := 0
		rearLightAndBrakeLight := 0
		wiperOperation := 0
		carHorn := 0
		sideMirror := 0
		generalBodyCondition := 0
		// fetch the Inspection From the database
		inspection := &model.Inspection{}

		if errs := colmns.Scan(
			&(inspection.ID),
			&(inspection.GarageID),
			&(inspection.InspectorID),
			&(inspection.Drivername),
			&(inspection.VehicleModel),
			&(inspection.VehicleYear),
			&(inspection.VehicleMake),
			&(inspection.VehicleColor),
			&(inspection.LicensePlate),
			&(inspection.FrontImage),
			&(inspection.LeftSideImage),
			&(inspection.RightSideImage),
			&(inspection.BackImage),
			&(inspection.SignatureImage),
			&(inspection.VinNumber),
			&handbrake,
			&steeringSystem,
			&brakeSystem,
			&seatBelt,
			&doorAndWindow,
			&dashBoardLight,
			&windshield,
			&baggageDoorWindow,
			&gearBox,
			&shockAbsorber,
			&frontHighAndLowBeamLight,
			&rearLightAndBrakeLight,
			&wiperOperation,
			&carHorn,
			&sideMirror,
			&generalBodyCondition,
			&(inspection.DriverPerformance),
			&(inspection.Balancing),
			&(inspection.Hazard),
			&(inspection.SignalLightUsage),
			&(inspection.Passed)); errs == nil {
			inspection.HandBrake = insorrepo.GetFunctionalityResultByID(ctx, handbrake)
			inspection.SteeringSystem = insorrepo.GetFunctionalityResultByID(ctx, steeringSystem)
			inspection.BrakeSystem = insorrepo.GetFunctionalityResultByID(ctx, brakeSystem)
			inspection.SeatBelt = insorrepo.GetFunctionalityResultByID(ctx, seatBelt)
			inspection.DoorAndWindow = insorrepo.GetFunctionalityResultByID(ctx, doorAndWindow)
			inspection.DashBoardLight = insorrepo.GetFunctionalityResultByID(ctx, dashBoardLight)
			inspection.WindShield = insorrepo.GetFunctionalityResultByID(ctx, windshield)
			inspection.BaggageDoorWindow = insorrepo.GetFunctionalityResultByID(ctx, baggageDoorWindow)
			inspection.GearBox = insorrepo.GetFunctionalityResultByID(ctx, gearBox)
			inspection.ShockAbsorber = insorrepo.GetFunctionalityResultByID(ctx, shockAbsorber)
			inspection.FrontHighAndLowBeamLight = insorrepo.GetFunctionalityResultByID(ctx, frontHighAndLowBeamLight)
			inspection.RearLightAndBrakeLight = insorrepo.GetFunctionalityResultByID(ctx, rearLightAndBrakeLight)
			inspection.WiperOperation = insorrepo.GetFunctionalityResultByID(ctx, wiperOperation)
			inspection.CarHorn = insorrepo.GetFunctionalityResultByID(ctx, carHorn)
			inspection.SideMirrors = insorrepo.GetFunctionalityResultByID(ctx, sideMirror)
			inspection.GeneralBodyCondition = insorrepo.GetFunctionalityResultByID(ctx, generalBodyCondition)
			inspections = append(inspections, inspection)
		}
	}
	return inspections, nil
}

// GetFunctionalityResult instance by ID
// returning functionality result value searching from the possible Functionality Results list in the
// state package and if it doesn't found any it will search from the database
// the parameter is the ID of the Functionality results table instance
// if it doesn't find any from either of them it will return a nil instance
func (insorrepo *InspectorRepo) GetFunctionalityResultByID(ctx context.Context, ID int) *model.FunctionalityResult {
	//  looping over the functionality results if found return the instance
	for _, fr := range state.FunctionalityResultInstances {
		if fr.ID == uint(ID) {
			return fr
		}
	}
	// search from the database using the ID
	funcres := &model.FunctionalityResult{}
	err := insorrepo.DB.QueryRow(ctx, "SELECT * FROM  functionality_results WHERE id=$1", ID).Scan(&(funcres.ID), &(funcres.Result), &(funcres.Reason))
	if err != nil {
		println(err.Error())
		return nil
	}
	return funcres

}

// GetInspectoryID (ctx context.Context) (*model.Inspection, error)
func (insorrepo *InspectorRepo) GetInspectorByID(ctx context.Context) (*model.Inspector, error) {
	inspectorID, val := ctx.Value("inspector_id").(uint)
	if !val {
		return nil, errors.New("Internal Error ")
	}
	inspector := &model.Inspector{}
	era := insorrepo.DB.QueryRow(ctx, "SELECT * FROM inspectors WHERE id=$1", inspectorID).Scan(
		&(inspector.ID),
		&(inspector.Email),
		&(inspector.Firstname),
		&(inspector.Middlename),
		&(inspector.Lastname),
		&(inspector.Password),
		&(inspector.Imageurl),
		&(inspector.GarageID),
		&(inspector.Createdby))

	if era != nil {
		println(era.Error())
		return nil, era
	}
	return inspector, nil
}

// UpdateProfileImage (ctx context.Context) error
func (insorrepo *InspectorRepo) UpdateProfileImage(ctx context.Context) error {
	defer recover()
	inspector := ctx.Value("inspector").(*model.Inspector)
	cmt, era := insorrepo.DB.Exec(ctx, "UPDATE inspectors SET imageurl=$2 WHERE id=$1", inspector.ID , inspector.Imageurl)
	if era != nil || cmt.RowsAffected() == 0 {
		if era != nil {
			println(era.Error())
		}
		return errors.New(" ERRROR : DB Error ... ")
	}
	return nil
}

// DeleteInspectorByID (ctx context.Context) error 
func (insorrepo *InspectorRepo) DeleteInspectorByID(ctx context.Context) error  {
	inspectorID := ctx.Value("inspector_id").(uint)
	cmd, era := insorrepo.DB.Exec(ctx, " DELETE FROM inspectors WHERE id=$1 ", inspectorID)
	if era != nil || cmd.RowsAffected() == 0 {
		return errors.New("No Rows Deleted ")
	}
	return nil
}