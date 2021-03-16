package state

import (
	"strings"

	"github.com/samuael/Project/CarInspection/pkg/constants/model"
)

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

var (
	// FunctionalityResultInstances  ... representing the possible failure resons
	//  reusing their ID as a foreign key to minimize redundant datas
	// This data will be inserted at the time of database initialization .
	FunctionalityResultInstances = []*model.FunctionalityResult{

		&model.FunctionalityResult{
			ID:     1,
			Result: false,
			Reason: "Lose Handbrake",
		},
		&model.FunctionalityResult{
			ID:     2,
			Result: false,
			Reason: "Rigt rear hand brake not working ",
		},
		&model.FunctionalityResult{
			ID:     3,
			Result: false,
			Reason: "Left rear hand brake not working ",
		},
		&model.FunctionalityResult{
			ID:     4,
			Result: false,
			Reason: "Handbrake not functional ",
		},
		&model.FunctionalityResult{
			ID:     5,
			Result: false,
			Reason: "Front brake not enough ",
		},
		&model.FunctionalityResult{
			ID:     6,
			Result: false,
			Reason: "Back brakes not enough ",
		},
		&model.FunctionalityResult{
			ID:     7,
			Result: false,
			Reason: "ABS warning on ",
		},
		&model.FunctionalityResult{
			ID:     8,
			Result: false,
			Reason: "Steering Shaing ",
		},
		&model.FunctionalityResult{
			ID:     9,
			Result: false,
			Reason: "Sound from steering ",
		},
		&model.FunctionalityResult{
			ID:     10,
			Result: false,
			Reason: "Left Outer tie rod glap ",
		},
		&model.FunctionalityResult{
			ID:     11,
			Result: false,
			Reason: "Right Outer tie rod glap ",
		},
		&model.FunctionalityResult{
			ID:     12,
			Result: false,
			Reason: "Left Inner tie rod glap ",
		},
		&model.FunctionalityResult{
			ID:     13,
			Result: false,
			Reason: "Right Inner tie rod glap ",
		},
		&model.FunctionalityResult{
			ID:     14,
			Result: false,
			Reason: "Driver Seatbelt malfunction ",
		},
		&model.FunctionalityResult{
			ID:     15,
			Result: false,
			Reason: "Passenger seatbelt malfunction ",
		},
		&model.FunctionalityResult{
			ID:     16,
			Result: false,
			Reason: "Driver door malfunciton ",
		},
		&model.FunctionalityResult{
			ID:     17,
			Result: false,
			Reason: "Right Rear passenger door malfunction ",
		},
		&model.FunctionalityResult{
			ID:     18,
			Result: false,
			Reason: "left rear passenger door malfunction ",
		},
		&model.FunctionalityResult{
			ID:     19,
			Result: false,
			Reason: "Left front passenger door malfunction ",
		},
		&model.FunctionalityResult{
			ID:     20,
			Result: false,
			Reason: "Driver window malfunction ",
		},
		&model.FunctionalityResult{
			ID:     21,
			Result: false,
			Reason: "Right rear window malfunction ",
		},
		&model.FunctionalityResult{
			ID:     22,
			Result: false,
			Reason: "ABS warning on ",
		},
		&model.FunctionalityResult{
			ID:     23,
			Result: false,
			Reason: "Check Engine Warning on ",
		},
		&model.FunctionalityResult{
			ID:     24,
			Result: false,
			Reason: "SRS Airbag Warning on ",
		},
		&model.FunctionalityResult{
			ID:     25,
			Result: false,
			Reason: "Windshild cracked /visibility impaired ",
		},
		&model.FunctionalityResult{
			ID:     26,
			Result: false,
			Reason: "Rear Window defect /visibility impaired ",
		},
		&model.FunctionalityResult{
			ID:     27,
			Result: false,
			Reason: "Gearbox defective ",
		},
		&model.FunctionalityResult{
			ID:     28,
			Result: false,
			Reason: " Front shock absorber malfunction ",
		},
		&model.FunctionalityResult{
			ID:     29,
			Result: false,
			Reason: "Back shock absorber malfunction ",
		},
		&model.FunctionalityResult{
			ID:     30,
			Result: false,
			Reason: "Leakage front left shock absorber",
		},
		&model.FunctionalityResult{
			ID:     31,
			Result: false,
			Reason: "Leakage front right shock absorber",
		},
		&model.FunctionalityResult{
			ID:     32,
			Result: false,
			Reason: "Leakage rear left shock absorber",
		},
		&model.FunctionalityResult{
			ID:     33,
			Result: false,
			Reason: "Leakage rear right shock absorber",
		},
		&model.FunctionalityResult{
			ID:     34,
			Result: false,
			Reason: "Left parking light not working ",
		},
		&model.FunctionalityResult{
			ID:     35,
			Result: false,
			Reason: "Right parking light not working",
		},
		&model.FunctionalityResult{
			ID:     36,
			Result: false,
			Reason: "Left low beam parking not working",
		},
		&model.FunctionalityResult{
			ID:     37,
			Result: false,
			Reason: "Right low beam parking not working",
		},
		&model.FunctionalityResult{
			ID:     38,
			Result: false,
			Reason: "Left fog light not working",
		},
		&model.FunctionalityResult{
			ID:     39,
			Result: false,
			Reason: "Right fog light not working",
		},
		&model.FunctionalityResult{
			ID:     40,
			Result: false,
			Reason: "Left turn signal not working",
		},
		&model.FunctionalityResult{
			ID:     41,
			Result: false,
			Reason: "Right turn signal not working",
		},
		&model.FunctionalityResult{
			ID:     42,
			Result: false,
			Reason: "Hazard light not working",
		},
		&model.FunctionalityResult{
			ID:     43,
			Result: false,
			Reason: "Wiper not working",
		},
		&model.FunctionalityResult{
			ID:     44,
			Result: false,
			Reason: "Washer fluid not spraying",
		},
		&model.FunctionalityResult{
			ID:     45,
			Result: false,
			Reason: "Car horn not working",
		},
		&model.FunctionalityResult{
			ID:     46,
			Result: false,
			Reason: "Left turn signal not working",
		},
		&model.FunctionalityResult{
			ID:     47,
			Result: false,
			Reason: "Right turn signal not working",
		},
		&model.FunctionalityResult{
			ID:     48,
			Result: false,
			Reason: "Left parking light not working",
		},
		&model.FunctionalityResult{
			ID:     49,
			Result: false,
			Reason: "Right parking light not working ",
		},
		&model.FunctionalityResult{
			ID:     50,
			Result: false,
			Reason: "Left break light not working",
		},
		&model.FunctionalityResult{
			ID:     51,
			Result: false,
			Reason: "Right break light not working",
		},
		&model.FunctionalityResult{
			ID:     52,
			Result: false,
			Reason: "Third break light not working",
		},
		&model.FunctionalityResult{
			ID:     53,
			Result: false,
			Reason: "Reverse light not working",
		},
		&model.FunctionalityResult{
			ID:     54,
			Result: false,
			Reason: "Hazard light not working",
		},
		&model.FunctionalityResult{
			ID:     55,
			Result: false,
			Reason: " Left side mirror defective",
		},
		&model.FunctionalityResult{
			ID:     56,
			Result: false,
			Reason: " Right side mirror defective",
		},
		&model.FunctionalityResult{
			ID:     57,
			Result: false,
			Reason: " Left side mirror broken/missing",
		},
		&model.FunctionalityResult{
			ID:     58,
			Result: false,
			Reason: " Right side mirror broken/missing",
		},
		&model.FunctionalityResult{
			ID:     59,
			Result: false,
			Reason: " Damage on front/rear/doors ",
		},
		&model.FunctionalityResult{
			ID:     60,
			Result: true,
			Reason: "",
		},
	}
)

func GetMatchingFunctionalityResultID(fr *model.FunctionalityResult) uint {
	for _, efr := range FunctionalityResultInstances {
		if strings.Trim(strings.ToUpper(efr.Reason), " ") == "PASS" {
			return 60
		}
		if strings.Trim(strings.ToLower(efr.Reason), " ") == strings.Trim(strings.ToLower(fr.Reason), " ") {
			return efr.ID
		}
	}
	return 0
}
