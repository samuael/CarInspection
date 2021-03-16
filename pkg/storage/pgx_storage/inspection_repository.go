package pgx_storage

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/inspection"
)

// InspectionRepo struct representing inspection repository and it's methods
type InspectionRepo struct {
	DB *pgx.Conn
}

func NewInspectionRepo(conn *pgx.Conn) inspection.IInspectionRepo {
	return &InspectionRepo{
		DB: conn,
	}
}

func (insrepo *InspectionRepo) CreateInspection(ctx context.Context) (*model.Inspection, error) {
	// recover from exception incase it happens ...
	defer recover()
	// getting the inspection from the context
	inspection := ctx.Value("inspection").(*model.Inspection)
	// getting the FunctionalityResultID from the List of possible Functionality results listed in the state
	// package of the models .
	// if i didn't find the reason in the pre defined list of functionality results i will create an instance and pass it
	// to the newly created functionality result ID to the Inspection as a foreign reference key
	handbrake := state.GetMatchingFunctionalityResultID(inspection.HandBrake)
	if handbrake == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &handbrake, inspection.HandBrake)
		if err != nil {
			println("Handbrake ... ")
			return nil, err
		}
	}
	steeringSystem := state.GetMatchingFunctionalityResultID(inspection.SteeringSystem)
	if steeringSystem == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &steeringSystem, inspection.SteeringSystem)
		if err != nil {
			println("Steering System ... ")
			return nil, err
		}
	}
	brakeSystem := state.GetMatchingFunctionalityResultID(inspection.BrakeSystem)
	if brakeSystem == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &brakeSystem, inspection.BrakeSystem)
		if err != nil {
			println("Brake System ... ")
			return nil, err
		}
	}
	seatBelt := state.GetMatchingFunctionalityResultID(inspection.SeatBelt)
	if seatBelt == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &seatBelt, inspection.SeatBelt)
		if err != nil {
			println("Seat Belt System ... ")
			return nil, err
		}
	}
	doorAndWindow := state.GetMatchingFunctionalityResultID(inspection.DoorAndWindow)
	if doorAndWindow == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &doorAndWindow, inspection.DoorAndWindow)
		if err != nil || doorAndWindow == 0 {
			println("Door and Window System ... ")
			return nil, err
		}
	}
	dashBoardLight := state.GetMatchingFunctionalityResultID(inspection.DashBoardLight)
	if dashBoardLight == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &dashBoardLight, inspection.DashBoardLight)
		if err != nil {
			println("Dash Board Light ... ")
			return nil, err
		}
	}
	windshield := state.GetMatchingFunctionalityResultID(inspection.WindShield)
	if windshield == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &windshield, inspection.WindShield)
		if err != nil {
			println("Wind Shield  ... ")
			return nil, err
		}
	}
	baggageDoorWindow := state.GetMatchingFunctionalityResultID(inspection.BaggageDoorWindow)
	if baggageDoorWindow == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &baggageDoorWindow, inspection.BaggageDoorWindow)
		if err != nil {
			println("Baggage Door Window  ... ")
			return nil, err
		}
	}
	gearBox := state.GetMatchingFunctionalityResultID(inspection.GearBox)
	if gearBox == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &gearBox, inspection.GearBox)
		if err != nil {
			println(" Gear Box  ... ")
			return nil, err
		}
	}
	shockAbsorber := state.GetMatchingFunctionalityResultID(inspection.ShockAbsorber)
	if shockAbsorber == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &shockAbsorber, inspection.ShockAbsorber)
		if err != nil {
			println(" Shock Absorber  ... ")
			return nil, err
		}
	}
	frontHighAndLowBeamLight := state.GetMatchingFunctionalityResultID(inspection.FrontHighAndLowBeamLight)
	if frontHighAndLowBeamLight == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &frontHighAndLowBeamLight, inspection.FrontHighAndLowBeamLight)
		if err != nil {
			println(" Front High And Low Beam Light  ... ")
			return nil, err
		}
	}
	rearLightAndBrakeLight := state.GetMatchingFunctionalityResultID(inspection.RearLightAndBrakeLight)
	if rearLightAndBrakeLight == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &rearLightAndBrakeLight, inspection.RearLightAndBrakeLight)
		if err != nil {
			println(" Rear Light And Brake Light  ... ")
			return nil, err
		}
	}
	wiperOperation := state.GetMatchingFunctionalityResultID(inspection.WiperOperation)
	if wiperOperation == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &wiperOperation, inspection.WiperOperation)
		if err != nil {
			println(" Wiper Operation... ")
			return nil, err
		}
	}
	carHorn := state.GetMatchingFunctionalityResultID(inspection.CarHorn)
	if carHorn == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &carHorn, inspection.CarHorn)
		if err != nil {
			println(" Car Horn ... ")
			return nil, err
		}
	}
	sideMirror := state.GetMatchingFunctionalityResultID(inspection.SideMirrors)
	if sideMirror == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &sideMirror, inspection.SideMirrors)
		if err != nil {
			println(" Side Mirror  ... ")
			return nil, err
		}
	}
	generalBodyCondition := state.GetMatchingFunctionalityResultID(inspection.GeneralBodyCondition)
	if generalBodyCondition == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &generalBodyCondition, inspection.GeneralBodyCondition)
		if err != nil {
			println(" General Body Condition  ... ")
			return nil, err
		}
	}
	// inserting the inspection data to the database using the inspection instance and the variables which i instantiate above
	err := insrepo.DB.QueryRow(ctx, `INSERT INTO inspections (  
		garageid,
		inspector_id,
		drivername,
		vehicle_model,
		vehicle_year,
		vehicle_make,
		vehicle_color,
		license_plate,
		front_image,
		left_image,
		right_image,
		back_image,
		signature_image,
		vin_number,
		handbrake,
		steering_system,
		brake_system,
		seat_belt,
		door_and_window,
		dashboard_light,
		windshield,
		baggage_door_window,
		gear_box,
		shock_absorber,
		front_high_and_low_beam_light,
		rear_light_and_brake_light,
		wiper_operation,
		car_horn,
		side_mirror,
		general_body_condition,
		driver_performance,
		balancing,
		hazard,
		signal_light_usage ,
		passed )  VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14,
			$15,
			$16,
			$17,
			$18,
			$19,
			$20,
			$21,
			$22,
			$23,
			$24,
			$25,
			$26,
			$27,
			$28,
			$29,
			$30,
			$31,
			$32,
			$33,
			$34,
			$35
		) RETURNING id `,
		inspection.GarageID,
		inspection.InspectorID,
		inspection.Drivername,
		inspection.VehicleModel,
		inspection.VehicleYear,
		inspection.VehicleMake,
		inspection.VehicleColor,
		inspection.LicensePlate,
		inspection.FrontImage,
		inspection.LeftSideImage,
		inspection.RightSideImage,
		inspection.BackImage,
		inspection.SignatureImage,
		inspection.VinNumber,
		handbrake,
		steeringSystem,
		brakeSystem,
		seatBelt,
		doorAndWindow,
		dashBoardLight,
		windshield,
		baggageDoorWindow,
		gearBox,
		shockAbsorber,
		frontHighAndLowBeamLight,
		rearLightAndBrakeLight,
		wiperOperation,
		carHorn,
		sideMirror,
		generalBodyCondition,
		inspection.DriverPerformance,
		inspection.Balancing,
		inspection.Hazard,
		inspection.SignalLightUsage,
		inspection.Passed,
	).Scan(&(inspection.ID))
	if err != nil {
		print(err.Error())
		return nil, err
	}
	inspection.HandBrake.ID = handbrake
	inspection.SteeringSystem.ID = steeringSystem
	inspection.BrakeSystem.ID = brakeSystem
	inspection.SeatBelt.ID = seatBelt
	inspection.DoorAndWindow.ID = doorAndWindow
	inspection.DashBoardLight.ID = dashBoardLight
	inspection.WindShield.ID = windshield
	inspection.BaggageDoorWindow.ID = baggageDoorWindow
	inspection.GearBox.ID = gearBox
	inspection.ShockAbsorber.ID = shockAbsorber
	inspection.FrontHighAndLowBeamLight.ID = frontHighAndLowBeamLight
	inspection.RearLightAndBrakeLight.ID = rearLightAndBrakeLight
	inspection.WiperOperation.ID = wiperOperation
	inspection.CarHorn.ID = carHorn
	inspection.SideMirrors.ID = sideMirror
	inspection.GeneralBodyCondition.ID = generalBodyCondition
	return inspection, nil
}

// IsInspectionOwner (inspectionID uint, inspectorID uint) (bool, error)  to check whether the Inspection is created by the
// inspector whose id is specified in the parameter or not
//  the method
// returns true  , nil     if the inspector is the owner of the inspection
// return false  , nil     if the inspection is not created by the inspector whose id IS Specified in the parameter
// return false  , error   if the inspection doesn't exist
func (insrepo *InspectionRepo) IsInspectionOwner(ctx context.Context) (bool, error) {
	// getting the Inspection ID and the Inspector ID
	inspectorID := ctx.Value("inspector_id").(uint)
	inspectionID := ctx.Value("inspection_id").(uint)
	// selecting the cretedby column  from the Inspection table
	// if it doesn't have an error the INspection exists else not.
	createdBy := 0
	err := insrepo.DB.QueryRow(ctx, " SELECT inspector_id FROM inspections WHERE id=$1", inspectionID).Scan(&createdBy)
	if err != nil || createdBy == 0 {
		return false, err
	} else if createdBy != int(inspectorID) {
		return false, nil
	}
	return true, nil
}

// GetFunctionalityResultID ...
func (insrepo *InspectionRepo) GetFunctionalityResultID(ctx context.Context, val *uint, fr *model.FunctionalityResult) error {
	// try searching a functionality result instance from  the database getting the ID we can re use that ID again as a foreign key
	// as another way of minimizing redundancy
	err := insrepo.DB.QueryRow(ctx, "SELECT id FROM functionality_results WHERE reason ILIKE $1", fr.Reason).Scan(val)
	if err != nil {
		// inserting the new functionality result instance and setting the id to the handbrake table
		err = insrepo.DB.QueryRow(ctx, "INSERT INTO functionality_results ( result , reason ) VALUES ( $1, $2 ) returning id", fr.Result, fr.Reason).Scan(val)
	}
	if err != nil {

		println(err.Error())
		println(fr.Reason, fr.Result)
		return err
	}
	return nil
}

// GetFunctionalityResultIDGeneral ...
func (insrepo *InspectionRepo) GetFunctionalityResultIDGeneral(ctx context.Context, fr *model.FunctionalityResult) uint {
	frID := state.GetMatchingFunctionalityResultID(fr)
	if frID == 0 {
		err := insrepo.GetFunctionalityResultID(ctx, &frID, fr)
		if err != nil {
			return frID
		}
	}
	return frID
}

// GetFunctionalityResult instance by ID
// returning functionality result value searching from the possible Functionality Results list in the
// state package and if it doesn't found any it will search from the database
// the parameter is the ID of the Functionality results table instance
// if it doesn't find any from either of them it will return a nil instance
func (inserepo *InspectionRepo) GetFunctionalityResultByID(ctx context.Context, ID int) *model.FunctionalityResult {
	//  looping over the functionality results if found return the instance
	for _, fr := range state.FunctionalityResultInstances {
		if fr.ID == uint(ID) {
			return fr
		}
	}
	// search from the database using the ID
	funcres := &model.FunctionalityResult{}
	err := inserepo.DB.QueryRow(ctx, "SELECT * FROM  functionality_results WHERE id=$1", ID).Scan(&(funcres.ID), &(funcres.Result), &(funcres.Reason))
	if err != nil {
		return nil
	}
	return funcres

}

// GetInspectionByID(ctx context.Context) (*model.Inspection, error)
// inspection Repository method to fetch the inspection by ID
//  returns (nil  , error)  if the insepection DOesn't Exist
func (insrepo *InspectionRepo) GetInspectionByID(ctx context.Context) (*model.Inspection, error) {
	defer recover()
	inspection := &model.Inspection{}
	inspectionID := ctx.Value("inspection_id").(uint)
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
	err := insrepo.DB.QueryRow(ctx, "SELECT * FROM  inspections WHERE id= $1", inspectionID).Scan(
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
		&(inspection.Passed),
	)
	inspection.HandBrake = insrepo.GetFunctionalityResultByID(ctx, handbrake)
	inspection.SteeringSystem = insrepo.GetFunctionalityResultByID(ctx, steeringSystem)
	inspection.BrakeSystem = insrepo.GetFunctionalityResultByID(ctx, brakeSystem)
	inspection.SeatBelt = insrepo.GetFunctionalityResultByID(ctx, seatBelt)
	inspection.DoorAndWindow = insrepo.GetFunctionalityResultByID(ctx, doorAndWindow)
	inspection.DashBoardLight = insrepo.GetFunctionalityResultByID(ctx, dashBoardLight)
	inspection.WindShield = insrepo.GetFunctionalityResultByID(ctx, windshield)
	inspection.BaggageDoorWindow = insrepo.GetFunctionalityResultByID(ctx, baggageDoorWindow)
	inspection.GearBox = insrepo.GetFunctionalityResultByID(ctx, gearBox)
	inspection.ShockAbsorber = insrepo.GetFunctionalityResultByID(ctx, shockAbsorber)
	inspection.FrontHighAndLowBeamLight = insrepo.GetFunctionalityResultByID(ctx, frontHighAndLowBeamLight)
	inspection.RearLightAndBrakeLight = insrepo.GetFunctionalityResultByID(ctx, rearLightAndBrakeLight)
	inspection.WiperOperation = insrepo.GetFunctionalityResultByID(ctx, wiperOperation)
	inspection.CarHorn = insrepo.GetFunctionalityResultByID(ctx, carHorn)
	inspection.SideMirrors = insrepo.GetFunctionalityResultByID(ctx, sideMirror)
	inspection.GeneralBodyCondition = insrepo.GetFunctionalityResultByID(ctx, generalBodyCondition)

	if err != nil {
		return nil, err
	}
	return inspection, nil
}

// UpdateInspection ...
func (insrepo *InspectionRepo) UpdateInspection(ctx context.Context) (*model.Inspection, error) {
	defer recover()

	inspection := ctx.Value("inspection").(*model.Inspection)

	getMatchingFunctionalityResultID := func(val *uint, fr *model.FunctionalityResult) uint {
		*val = state.GetMatchingFunctionalityResultID(fr)
		if *val == 0 {
			err := insrepo.GetFunctionalityResultID(ctx, val, fr)
			if err != nil {
				return 60
			}
		}
		return *val
	}
	cmd, err := insrepo.DB.Exec(ctx, `UPDATE inspections SET 
	garageid=$36,
	inspector_id =				 	$2 ,
	drivername 						=$3,
	vehicle_model					=$4,
	vehicle_year					=$5,
	vehicle_make					=$6,
	vehicle_color					=$7,
	license_plate					=$8,
	front_image						=$9,
	left_image						=$10,
	right_image						=$11,
	back_image						=$12,
	signature_image					=$13,
	vin_number						=$14,
	handbrake						=$15,
	steering_system					=$16,
	brake_system					=$17,
	seat_belt						=$18,
	door_and_window					=$19,
	dashboard_light					=$20,
	windshield						=$21,
	baggage_door_window 			=$22,
	gear_box						=$23,
	shock_absorber					=$24,
	front_high_and_low_beam_light	=$25,
	rear_light_and_brake_light		=$26,
	wiper_operation					=$27,
	car_horn						=$28,
	side_mirror						=$29,	
	general_body_condition			=$30,
	driver_performance				=$31,
	balancing						=$32,
	hazard							=$33,	
	signal_light_usage 				=$34,
	passed = $35 WHERE id=$1`,
		inspection.ID,
		inspection.InspectorID,
		inspection.Drivername,
		inspection.VehicleModel,
		inspection.VehicleYear,
		inspection.VehicleMake,
		inspection.VehicleColor,
		inspection.LicensePlate,
		inspection.FrontImage,
		inspection.LeftSideImage,
		inspection.RightSideImage,
		inspection.BackImage,
		inspection.SignatureImage,
		inspection.VinNumber,

		getMatchingFunctionalityResultID(&inspection.HandBrake.ID, inspection.HandBrake),
		getMatchingFunctionalityResultID(&inspection.SteeringSystem.ID, inspection.SteeringSystem),
		getMatchingFunctionalityResultID(&inspection.BrakeSystem.ID, inspection.BrakeSystem),
		getMatchingFunctionalityResultID(&inspection.SeatBelt.ID, inspection.SeatBelt),
		getMatchingFunctionalityResultID(&inspection.DoorAndWindow.ID, inspection.DoorAndWindow),
		getMatchingFunctionalityResultID(&inspection.DashBoardLight.ID, inspection.DashBoardLight),
		getMatchingFunctionalityResultID(&inspection.WindShield.ID, inspection.WindShield),
		getMatchingFunctionalityResultID(&inspection.BaggageDoorWindow.ID, inspection.BaggageDoorWindow),
		getMatchingFunctionalityResultID(&inspection.GearBox.ID, inspection.GearBox),
		getMatchingFunctionalityResultID(&inspection.ShockAbsorber.ID, inspection.ShockAbsorber),
		getMatchingFunctionalityResultID(&inspection.FrontHighAndLowBeamLight.ID, inspection.FrontHighAndLowBeamLight),
		getMatchingFunctionalityResultID(&inspection.RearLightAndBrakeLight.ID, inspection.RearLightAndBrakeLight),
		getMatchingFunctionalityResultID(&inspection.WiperOperation.ID, inspection.WiperOperation),
		getMatchingFunctionalityResultID(&inspection.CarHorn.ID, inspection.CarHorn),
		getMatchingFunctionalityResultID(&inspection.SideMirrors.ID, inspection.SideMirrors),
		getMatchingFunctionalityResultID(&inspection.GeneralBodyCondition.ID, inspection.GeneralBodyCondition),

		inspection.DriverPerformance,
		inspection.Balancing,
		inspection.Hazard,
		inspection.SignalLightUsage,
		inspection.Passed,
		inspection.GarageID,
	)
	if err != nil || cmd.RowsAffected() == 0 {
		return nil, err
	}
	return inspection, nil
}
