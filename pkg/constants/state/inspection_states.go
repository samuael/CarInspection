package state

const (
	// HandBrakFunctionality ...
	Reason_HandBrake = iota
	// Reason_BrakeSystem  ...
	Reason_Brake_System
	// Reason_Steering_System ...
	Reason_Steering_System
	// Reason_Seatbelt ...
	Reason_Seatbelt
	// Reason_Door_AND_Window ...
	Reason_Door_AND_Window
	// Reason_Dashboard_light ...
	Reason_Dashboard_light
	// Reason_Windshield ...
	Reason_Windshield
	// Reason_Baggage_Door_window ...
	Reason_Baggage_Door_window
	// Reason_Gearbox ...
	Reason_Gearbox
	// Reason_Shock_Absorber ...
	Reason_Shock_Absorber
	// Reason_High_And_Low_Beam_Light ...
	Reason_High_And_Low_Beam_Light
	// Reason_Tail_Light_And_Break_Light ...
	Reason_Tail_Light_And_Break_Light
	// Reason_Wiper_Operation ...
	Reason_Wiper_Operation
	// Reason_Car_Horn ...
	Reason_Car_Horn
	// Reason_Side_Mirrors ...
	Reason_Side_Mirrors
	// Reason_General_Body_Condition ...
	Reason_General_Body_Condition
)

var FailureReasonMap = map[int][]string{
	Reason_HandBrake: []string{
		"Lose Handbrake",
		"Rigt rear hand brake not working ",
		"Left rear hand brake not working ",
		"Handbrake not functional ",
	},

	Reason_Brake_System: []string{
		"Front brake not enough ",
		"Back brakes not enough ",
		"ABS warning on ",
	},
	Reason_Steering_System: []string{
		"Steering Shaing ",
		"Sound from steering ",
		"Left Outer tie rod glap ",
		"Right Outer tie rod glap ",
		"Left Inner tie rod glap ",
		"Right Inner tie rod glap ",
	},
	Reason_Seatbelt: []string{
		"Driver Seatbelt malfunction ",
		"Passenger seatbelt malfunction ",
	},
	Reason_Door_AND_Window: []string{
		"Driver door malfunciton ",
		"Right Rear passenger door malfunction ",
		"left rear passenger door malfunction ",
		"Left front passenger door malfunction ",
		"Driver window malfunction ",
		"Right rear window malfunction ",
	},
	Reason_Dashboard_light: []string{
		"ABS warning on ",
		"Check Engine Warning on ",
		"SRS Airbag Warning on ",
	},
	Reason_Windshield: []string{
		"Windshild cracked /visibility impaired ",
	},
	Reason_Baggage_Door_window: []string{
		"Rear Window defect /visibility impaired ",
	},
	Reason_Gearbox: []string{
		"Gearbox defective ",
	},
	Reason_Shock_Absorber: []string{
		" Front shock absorber malfunction ",
		"Back shock absorber malfunction ",
		"Leakage front left shock absorber",
		"Leakage front right shock absorber",
		"Leakage rear left shock absorber",
		"Leakage rear right shock absorber",
	},

	Reason_High_And_Low_Beam_Light: []string{
		"Left parking light not working ",
		"Right parking light not working",
		"Left low beam parking not working",
		"Right low beam parking not working",
		"Left fog light not working",
		"Right fog light not working",
		"Left turn signal not working",
		"Right turn signal not working",
		"Hazard light not working",
	},
	Reason_Tail_Light_And_Break_Light: []string{
		"Left turn signal not working",
		"Right turn signal not working",
		"Left parking light not working",
		"Right parking light not working ",
		"Left break light not working",
		"Right break light not working",
		"Third break light not working",
		"Reverse light not working",
		"Hazard light not working",
	},
	Reason_Wiper_Operation: []string{
		"Wiper not working",
		"Washer fluid not spraying",
	},
	Reason_Car_Horn: []string{
		"Car horn not working",
	},
	Reason_Side_Mirrors: []string{
		" Left side mirror defective",
		" Right side mirror defective",
		" Left side mirror broken/missing",
		" Right side mirror broken/missing",
	},
	Reason_General_Body_Condition: []string{
		" Damage on front/rear/doors ",
	},
}
