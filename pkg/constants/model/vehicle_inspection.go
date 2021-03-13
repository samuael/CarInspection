package model

// Inspection ...
type Inspection struct {
	ID                       uint                 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	InspectorID              uint                 `json:"inspector_id"`
	Drivername               string               `json:"driver_name"`
	VehicleModel             string               `json:"vehicle_model"`
	VehicleYear              string               `json:"vehicle_year"`
	VehicleMake              string               `json:"vehicle_make"`
	VehicleColor             string               `json:"vehicle_color"`
	LicensePlate             string               `json:"license_plate"`
	FrontImage               string               `json:"front_image"`
	LeftSideImage            string               `json:"lest_side_image"`
	RightSideImage           string               `json:"right_side_image"`
	BackImage                string               `json:"back_image"`
	SignatureImage           string               `json:"signature"`
	VinNumber                string               `json:"vin_number"`
	HandBrake                *FunctionalityResult `json:"hand_brake"  					gorm:"hand_brake,	foreignkey:HandBrakeRefer"`
	SteeringSystem           *FunctionalityResult `json:"steering_system"  				gorm:"steering_system,	foreignkey:SteeringSystemRefer"`
	BrakeSystem              *FunctionalityResult `json:"brake_system" 					gorm:"brake_system,	foreignkey:BrakeSystemRefer"`
	SeatBelt                 *FunctionalityResult `json:"seat_belt" 					gorm:"seat_belt,	foreignkey:SeatBeltRefer"`
	DoorAndWindow            *FunctionalityResult `json:"door_and_window"   			gorm:"door_and_window,	foreignkey:DoorAndWindowRefer"`
	DashBoardLight           *FunctionalityResult `json:"dash_board_light"  			gorm:"dash_board_light,	foreignkey:DashBoardLightRefer"`
	WindShield               *FunctionalityResult `json:"wind_shield" 					gorm:"wind_shield,	foreignkey:WindShieldRefer"`
	BaggageDoorWindow        *FunctionalityResult `json:"baggage_door_window" 			gorm:"baggage_door_window,	foreignkey:BaggageDoorWindowRefer"`
	GearBox                  *FunctionalityResult `json:"gear_box" 						gorm:"gear_box,	foreignkey:GearBoxRefer"`
	ShockAbsorber            *FunctionalityResult `json:"shock_absorber" 				gorm:"shock_absorber,	foreignkey:ShockAbsorberRefer"`
	FrontHighAndLowBeamLight *FunctionalityResult `json:"high_and_low_beam_light" 		gorm:"high_and_low_beam_light,	foreignkey:FrontHighAndLowBeamLightRefer"`
	RearLightAndBrakeLight   *FunctionalityResult `json:"rear_light_and_break_light" 	gorm:"rear_light_and_break_light,	foreignkey:RearLightAndBrakeLightRefer"`
	WiperOperation           *FunctionalityResult `json:"wiper_operation" 				gorm:"wiper_operation,	foreignkey:WiperOperationRefer"`
	CarHorn                  *FunctionalityResult `json:"car_horn" 						gorm:"car_horn,	foreignkey:CarHornRefer"`
	SideMirrors              *FunctionalityResult `json:"side_mirrors" 					gorm:"side_mirrors,	foreignkey:SideMirrorsRefer"`
	GeneralBodyCondition     *FunctionalityResult `json:"general_body_condition" 		gorm:"general_body_condition,	foreignkey:GeneralBodyConditionRefer"`
	DriverPerformance        bool                 `json:"driver_performance"`
	Balancing                bool                 `json:"balancing"`
	Hazard                   bool                 `json:"hazard"`
	SignalLightUsage         bool                 `json:"signal_light_usage"` // turn indicator
}

// FunctionalityResult to be used by a list of functionality parameters and their reasons
type FunctionalityResult struct {
	ID     uint   `gorm:"primaryKey;autoIncrement:true"`
	Result bool   `json:"result" gorm:"boolean;not null;"`           // representing the functionality result
	Reason string `json:"reason" gorm:"type:varchar(255);not null;"` // To represent the failure reason
}
