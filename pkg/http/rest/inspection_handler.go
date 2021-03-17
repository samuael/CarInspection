package rest

import (
	// "encoding/json"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	// "strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/garage"
	"github.com/samuael/Project/CarInspection/pkg/inspection"
	"github.com/samuael/Project/CarInspection/pkg/inspector"
	"github.com/samuael/Project/CarInspection/platforms/form"
	"github.com/samuael/Project/CarInspection/platforms/helper"
	"github.com/samuael/Project/CarInspection/platforms/pdf"
)

// PostHandler provides access to Post api methods.
type IInspectionHandler interface {
	CreateInspection(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	EditInspection(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateInspectionFiles(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteInspection(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetInspectionByID(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetInspectionAsPDF(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	// GetInspections(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// AddInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// DeleteInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// EditInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type InspectionHandler struct {
	InspectionSer inspection.IInspectionService
	Template      *template.Template
	InspectorSer  inspector.IInspectorService
	GarageSer     garage.IGarageService
}

// NewInspectionHandler ...
func NewInspectionHandler(inser inspection.IInspectionService, temp *template.Template, insor inspector.IInspectorService, garageser garage.IGarageService) IInspectionHandler {
	return &InspectionHandler{
		InspectionSer: inser,
		Template:      temp,
		InspectorSer:  insor,
		GarageSer:     garageser,
	}
}

func (h InspectionHandler) GetInspections(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// w.Header().Set("Content-Type", "application/json")
	// list, err := h.l.GetMyInspections(9)
	// if err != nil {
	// 	http.Error(w, "Failed to get posts", http.StatusBadRequest)
	// 	return
	// }
	// json.NewEncoder(w).Encode(list)
}

// AddPost handler for POST /api/post requests
// INPUT  : MULTI-PART FORM FILE
//  Output  :  JSON
/* {
    "success": true,
    "has_error": false,
    "message": " Succesfuly Created Inspection ",
    "inspection": {
        "id": 1,
        "garage_id": 1,
        "inspector_id": 1,
        "driver_name": "teferi",
        "vehicle_model": "model223122",
        "vehicle_year": "1997",
        "vehicle_make": "china",
        "vehicle_color": "red",
        "license_plate": "23424243",
        "front_image": "4QNLKG2pSghqM9u.png",
        "left_side_image": "2XdwcFpHKef4GWV.png",
        "right_side_image": "yqV8UuBs3w4h5Cm.png",
        "back_image": "hHYRUDhBDZjKcXr.jpg",
        "signature": "hHYRUDhBDZjKcXr.jpg",
        "vin_number": "sdfhsfkjahsdkjfhasdjk",
        "hand_brake": {
            "ID": 0,
            "result": true,
            "reason": ""
        },
        "steering_system": {
            "ID": 0,
            "result": true,
            "reason": ""
        },
        "brake_system": {
            "ID": 0,
            "result": true,
            "reason": ""
        },
        "seat_belt": {
            "ID": 0,
            "result": false,
            "reason": "Driver Seatbelt malfunction "
        },
        "door_and_window": {
            "ID": 0,
            "result": false,
            "reason": "Door And Window are not fine man"
        },
        "dash_board_light": {
            "ID": 0,
            "result": false,
            "reason": "ABS warning on"
        },
        "wind_shield": {
            "ID": 0,
            "result": false,
            "reason": "Windshild cracked /visibility impaired"
        },
        "baggage_door_window": {
            "ID": 0,
            "result": false,
            "reason": "Rear Window defect /visibility impaired "
        },
        "gear_box": {
            "ID": 0,
            "result": false,
            "reason": "gear box not worling well "
        },
        "shock_absorber": {
            "ID": 0,
            "result": false,
            "reason": "shock absorber not working well "
        },
        "high_and_low_beam_light": {
            "ID": 0,
            "result": false,
            "reason": "Hazard light not working"
        },
        "rear_light_and_break_light": {
            "ID": 0,
            "result": false,
            "reason": "Hazard light not working"
        },
        "wiper_operation": {
            "ID": 0,
            "result": false,
            "reason": "Wiper not working"
        },
        "car_horn": {
            "ID": 0,
            "result": false,
            "reason": "Car horn not working"
        },
        "side_mirrors": {
            "ID": 0,
            "result": false,
            "reason": "Left side mirror defective"
        },
        "general_body_condition": {
            "ID": 0,
            "result": false,
            "reason": "Damage on front/rear/doors "
        },
        "driver_performance": true,
        "balancing": true,
        "hazard": true,
        "signal_light_usage": true
    },
    "errors": null
}
 }*/
//  COnsideration :: if any criterions are passed it must include a Strig Passed
func (h InspectionHandler) CreateInspection(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)

	res := &model.InspectionCreationResponse{
		Success:    false,
		Message:    "",
		HasError:   false,
		Inspection: nil,
	}
	request.ParseMultipartForm(999999999999999999)
	input := form.Input{
		Values:  request.MultipartForm.Value,
		VErrors: form.ValidationErrors{},
	}
	input.Required(
		"driver_name",
		"vehicle_model",
		"vehicle_make",
		"vehicle_color",
		"license_plate",
		"vin_number",

		"hand_brake",
		"steering_system",
		"brake_system",
		"seat_belt",
		"door_and_window",
		"dash_board_light",
		"wind_shield",
		"baggage_door_window",
		"gear_box",
		"shock_absorber",
		"high_and_low_beam_light",
		"rear_light_and_break_light",
		"wiper_operation",
		"car_horn",
		"side_mirrors",
		"general_body_condition",
		"driver_performance",
		"balancing",
		"hazard",
		"signal_light_usage")
	driverPerformance := input.ParseBoolean("driver_performance")
	balancing := input.ParseBoolean("balancing")
	hazard := input.ParseBoolean("hazard")
	signalLightUsage := input.ParseBoolean("signal_light_usage")

	// checking the presence of image files
	// those are "front_image","left_side_image" ,"right_side_image" , "back_image" , "signature" ,
	frontImage, header, eraa := input.GetFormFile(request, "front_image")
	if eraa == nil {
		defer frontImage.Close()
	}
	leftImage, bheader, erab := input.GetFormFile(request, "left_side_image")
	if erab == nil {
		defer leftImage.Close()
	}
	rightImage, cheader, erac := input.GetFormFile(request, "right_side_image")
	if erac == nil {
		defer rightImage.Close()
	}
	backImage, dheader, erad := input.GetFormFile(request, "back_image")
	if erad == nil {
		defer backImage.Close()
	}
	signature, eheader, erae := input.GetFormFile(request, "signature")
	if erae == nil {
		defer signature.Close()
	}
	// close them if they are not nul for the future
	if !input.Valid() {
		res.HasError = true
		res.Errors = input.VErrors
		response.WriteHeader(http.StatusBadRequest)
		res.Message = os.Getenv("INVALID_INPUT")
		response.Write(helper.MarshalThis(res))
		return
	}

	licensePlate := request.FormValue("license_plate")
	vinNumber := request.FormValue("vin_number")

	ctx = context.WithValue(ctx, "license_plate", licensePlate)
	ctx = context.WithValue(ctx, "vin_number", vinNumber)

	exists := h.InspectionSer.DoesThisVahicheWithLicensePlateExist(ctx)
	if exists {
		res.HasError = false
		res.Errors = nil
		// Status Code conflict
		response.WriteHeader(409)
		res.Message = fmt.Sprintf("Inspection With Specified License Plate '%s' Exists!", licensePlate)
		response.Write(helper.MarshalThis(res))
		return
	}
	exists = h.InspectionSer.DoesThisVehicleWithVinNumberExists(ctx)
	if exists {
		res.HasError = false
		res.Errors = nil
		// Status Code conflict
		response.WriteHeader(409)
		res.Message = fmt.Sprintf("Inspection With Specified Vin Number '%s' Exists!", vinNumber)
		response.Write(helper.MarshalThis(res))
		return
	}
	//  Generate random names for the files filenames
	frontFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + "." + helper.GetExtension(header.Filename)
	leftFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + "." + helper.GetExtension(bheader.Filename)
	rightFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + "." + helper.GetExtension(cheader.Filename)
	backFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + "." + helper.GetExtension(dheader.Filename)
	signatureFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + "." + helper.GetExtension(eheader.Filename)

	pathToAssetsDirectory := os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY")
	// Creating files for the new comming files
	newFrontImage, era := os.Create(pathToAssetsDirectory + "frontImages/" + frontFilename)
	if era == nil {
		defer newFrontImage.Close()
	}
	newLeftImage, erb := os.Create(pathToAssetsDirectory + "leftImages/" + leftFilename)
	if erb != nil {
		defer newLeftImage.Close()
	}
	newRightmage, erc := os.Create(pathToAssetsDirectory + "rightImages/" + rightFilename)
	if erc != nil {
		defer newRightmage.Close()
	}
	newBackImage, erd := os.Create(pathToAssetsDirectory + "backImages/" + backFilename)
	if erd != nil {
		defer newBackImage.Close()
	}
	newSignatureImage, ere := os.Create(pathToAssetsDirectory + "signatureImages/" + signatureFilename)
	if ere != nil {
		defer newSignatureImage.Close()
	}
	if era != nil || erb != nil || erc != nil || erd != nil || erd != nil {
		response.WriteHeader(http.StatusInternalServerError)
		input.VErrors.Add("Internal Server Error", " Files Exception ")
		res.Success = false
		res.Message = " Internal Server Error "
		response.Write(helper.MarshalThis(res))
		return
	}
	if era != nil || erb != nil || erc != nil || erd != nil || ere != nil {
		// if any of the file creation faiils it will creation of the inspection will fail
		response.WriteHeader(http.StatusInternalServerError)
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		res.HasError = true
		response.Write(helper.MarshalThis(res))
		return
	}
	// copying the image files into the newly created files
	_, ea := io.Copy(newFrontImage, frontImage)
	_, eb := io.Copy(newLeftImage, leftImage)
	_, ec := io.Copy(newRightmage, rightImage)
	_, ed := io.Copy(newBackImage, backImage)
	_, ee := io.Copy(newSignatureImage, signature)
	if ea != nil || eb != nil || ec != nil || ed != nil || ee != nil {
		os.Remove(pathToAssetsDirectory + "frontImages/" + frontFilename)
		os.Remove(pathToAssetsDirectory + "leftImages/" + leftFilename)
		os.Remove(pathToAssetsDirectory + "rightImages/" + rightFilename)
		os.Remove(pathToAssetsDirectory + "backImages/" + backFilename)
		os.Remove(pathToAssetsDirectory + "signatureImages/" + backFilename)

		response.WriteHeader(http.StatusInternalServerError)
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		response.Write(helper.MarshalThis(res))
		return
	}
	passed := true
	getFunctionResult := func(valueName string) *model.FunctionalityResult {
		if strings.ToUpper(valueName) == "PASS" {
			return &model.FunctionalityResult{
				Result: true,
				Reason: "",
			}
		}
		passed = false
		return &model.FunctionalityResult{
			Result: false,
			Reason: valueName,
		}
	}
	if driverPerformance == false ||
		balancing == false ||
		hazard == false ||
		signalLightUsage == false {
		passed = false
	}
	// craeting the inspection instance ,
	inspection := &model.Inspection{
		GarageID:                 session.GarageID,
		InspectorID:              uint(session.ID),
		Drivername:               request.FormValue("driver_name"),
		VehicleModel:             request.FormValue("vehicle_model"),
		VehicleYear:              request.FormValue("vehicle_year"),
		VehicleMake:              request.FormValue("vehicle_make"),
		VehicleColor:             request.FormValue("vehicle_color"),
		LicensePlate:             request.FormValue("license_plate"),
		VinNumber:                request.FormValue("vin_number"),
		HandBrake:                getFunctionResult(request.FormValue("hand_brake")),
		SteeringSystem:           getFunctionResult(request.FormValue("steering_system")),
		BrakeSystem:              getFunctionResult(request.FormValue("brake_system")),
		SeatBelt:                 getFunctionResult(request.FormValue("seat_belt")),
		DoorAndWindow:            getFunctionResult(request.FormValue("door_and_window")),
		DashBoardLight:           getFunctionResult(request.FormValue("dash_board_light")),
		WindShield:               getFunctionResult(request.FormValue("wind_shield")),
		BaggageDoorWindow:        getFunctionResult(request.FormValue("baggage_door_window")),
		GearBox:                  getFunctionResult(request.FormValue("gear_box")),
		ShockAbsorber:            getFunctionResult(request.FormValue("shock_absorber")),
		FrontHighAndLowBeamLight: getFunctionResult(request.FormValue("high_and_low_beam_light")),
		RearLightAndBrakeLight:   getFunctionResult(request.FormValue("rear_light_and_break_light")),
		WiperOperation:           getFunctionResult(request.FormValue("wiper_operation")),
		CarHorn:                  getFunctionResult(request.FormValue("car_horn")),
		SideMirrors:              getFunctionResult(request.FormValue("side_mirrors")),
		GeneralBodyCondition:     getFunctionResult(request.FormValue("general_body_condition")),
		DriverPerformance:        driverPerformance,
		Balancing:                balancing,
		Hazard:                   hazard,
		SignalLightUsage:         signalLightUsage,
		FrontImage:               "public/frontImages/" + frontFilename,
		LeftSideImage:            "public/leftImages/" + leftFilename,
		RightSideImage:           "public/rightImages/" + rightFilename,
		BackImage:                "public/backImages/" + backFilename,
		SignatureImage:           "public/signatureImages/" + signatureFilename,
		Passed:                   passed,
	}
	ctx = context.WithValue(ctx, "inspection", inspection)
	inspection, err := h.InspectionSer.CreateInspection(ctx)
	if err != nil || inspection == nil {
		println(" Error In the Inspection :::    |", err.Error(), "|\n")
		// If the inspection creation was not succesful i have to delete all the Image Files i have created
		os.Remove(pathToAssetsDirectory + "frontImages/" + frontFilename)
		os.Remove(pathToAssetsDirectory + "leftImages/" + leftFilename)
		os.Remove(pathToAssetsDirectory + "rightImages/" + rightFilename)
		os.Remove(pathToAssetsDirectory + "backImages/" + backFilename)
		os.Remove(pathToAssetsDirectory + "signatureImages/" + signatureFilename)

		response.WriteHeader(http.StatusInternalServerError)
		input.VErrors.Add("Internal Server Error", " Files Exception ")
		res.Success = false
		res.Message = " Internal Server Error "
		response.Write(helper.MarshalThis(res))
		return
	}
	res.Success = true
	res.Message = func() string {
		return fmt.Sprintf(" Succesfuly Created Inspection %s ", func() string {
			if passed {
				return " Vahicle passed the Test "
			}
			return " Vehicle don't Passed the Test "
		}())
	}()

	res.HasError = false
	res.Inspection = inspection
	response.WriteHeader(201)
	response.Write(helper.MarshalThis(res))
	return
}

// EditInspection method to edit an existing inspection
// METHOD : PUT
// INPUT   : JSON
/*
	INPUT : {
		{
    "success": true,
    "has_error": false,
    "message": " Succesfuly Created Inspection  Vehicle don't Passed the Test  ",
    "inspection": {
        "id": 1,
        "driver_name": "teferi",
        "vehicle_model": "model223122",
        "vehicle_year": "1997",
        "vehicle_make": "Europe",
        "vehicle_color": "red",
        "hand_brake": "PASS",
        "steering_system": "PASS",
        "brake_system": "PASS",
        "seat_belt": "PASS",
        "door_and_window": "PASS",
        "dash_board_light": "PASS",
        "wind_shield":"PASS",
        "baggage_door_window": "Another Reason",
        "gear_box": "Pass " ,
        "shock_absorber": "Another Reason" ,
        "high_and_low_beam_light":""Another Reason",
        "rear_light_and_break_light":"PASS",
        "wiper_operation": "PASS",
        "car_horn": "PASS" ,
        "side_mirrors": "Another reason",
        "general_body_condition": "PASS" ,
        "driver_performance": true,
        "balancing": true,
        "hazard": true,
        "signal_light_usage": true,
        "passed": false
    },
    "errors": null
}
	}

*/
// OUTPUT   : JSON
func (h InspectionHandler) EditInspection(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)

	res := &model.InspectionUpdateResponse{
		Success:    false,
		Message:    " Invalid Input ",
		Inspection: nil,
	}
	jdecode := json.NewDecoder(request.Body)
	inspection := &model.InspectionUpdate{}
	err := jdecode.Decode(inspection)
	if err != nil {
		fmt.Println(err.Error())
		response.WriteHeader(http.StatusBadRequest)
		res.Message = os.Getenv("INVALID_INPUT")
		response.Write(helper.MarshalThis(res))
		return
	}
	if inspection.ID == 0 {
		val, er := strconv.Atoi(params.ByName("id"))
		if er != nil || val == 0 {
			if er != nil {
				fmt.Println(er.Error())
			}
			response.WriteHeader(http.StatusBadRequest)
			res.Message = os.Getenv("INVALID_INPUT") + "\n The ID of Inspection Must be specified "
			response.Write(helper.MarshalThis(res))
			return
		}
		inspection.ID = (val)
	}
	ctx = context.WithValue(ctx, "inspection_id", uint(inspection.ID))
	ctx = context.WithValue(ctx, "inspector_id", uint(session.ID))
	// checking the whether the inspector is the owner of the inspection or not
	success, erra := h.InspectionSer.IsInspectionOwner(ctx)
	if erra != nil {
		response.WriteHeader(http.StatusNotModified)
		res.Message = os.Getenv("RESOURCE_NOT_FOUND") + "\n Inspection With this ID not Found ..."
		response.Write(helper.MarshalThis(res))
		return
	} else if !success {
		response.WriteHeader(http.StatusUnauthorized)
		res.Message = os.Getenv("UNAUTHORIZED_ACCESS") + "\n You are not authorized to update this Inspection!"
		response.Write(helper.MarshalThis(res))
		return
	}
	// The Inspection os Confirmed that . it is created by the inspector ..

	// getting the variables that are to be included in the updated inspection

	formerInspection, erra := h.InspectionSer.GetInspectionByID(ctx)
	if erra != nil || formerInspection == nil {
		response.WriteHeader(http.StatusInternalServerError)
		res.Success = false
		res.Message = " Internal Server Error "
		response.Write(helper.MarshalThis(res))
		return
	}

	passed := true
	changeCount := 0
	changeStringIfValidAndNew := func(new, old string) string {
		if new == old || new == "" || len(new) <= 2 {
			return old
		}
		changeCount++
		return new
	}
	// changing the new Inspection instance to include the chage
	formerInspection.Drivername = changeStringIfValidAndNew(inspection.Drivername, formerInspection.Drivername)
	formerInspection.VehicleModel = changeStringIfValidAndNew(inspection.VehicleModel, formerInspection.VehicleModel)
	formerInspection.VehicleYear = changeStringIfValidAndNew(inspection.VehicleYear, formerInspection.VehicleYear)
	formerInspection.VehicleMake = changeStringIfValidAndNew(inspection.VehicleMake, formerInspection.VehicleMake)
	formerInspection.VehicleColor = changeStringIfValidAndNew(inspection.VehicleColor, formerInspection.VehicleColor)

	getFunctionResult := func(valueName, oldreason string) *model.FunctionalityResult {
		if strings.ToUpper(valueName) == strings.ToUpper(oldreason) {
			if strings.ToUpper(valueName) == "PASS" {
				return &model.FunctionalityResult{
					Result: true,
					Reason: "",
				}
			} else {
				passed = false
				return &model.FunctionalityResult{
					Result: false,
					Reason: valueName,
				}
			}
		}
		changeCount++
		if strings.ToUpper(valueName) == "PASS" {
			return &model.FunctionalityResult{
				Result: true,
				Reason: "",
			}
		}
		passed = false
		return &model.FunctionalityResult{
			Result: false,
			Reason: valueName,
		}
	}

	getValidBooleanValue := func(new bool, oldval bool) bool {
		if new == oldval {
			if oldval == false {
				passed = false
			}
			return oldval
		}
		changeCount++
		if new == false {
			passed = false
		}
		return new
	}

	// getting the functionality result valued files
	formerInspection.HandBrake = getFunctionResult(inspection.HandBrake, formerInspection.HandBrake.Reason)
	formerInspection.SteeringSystem = getFunctionResult(inspection.SteeringSystem, formerInspection.SteeringSystem.Reason)
	formerInspection.BrakeSystem = getFunctionResult(inspection.BrakeSystem, formerInspection.BrakeSystem.Reason)
	formerInspection.SeatBelt = getFunctionResult(inspection.SeatBelt, formerInspection.SeatBelt.Reason)
	formerInspection.DoorAndWindow = getFunctionResult(inspection.DoorAndWindow, formerInspection.DoorAndWindow.Reason)
	formerInspection.DashBoardLight = getFunctionResult(inspection.DashBoardLight, formerInspection.DashBoardLight.Reason)
	formerInspection.WindShield = getFunctionResult(inspection.WindShield, formerInspection.WindShield.Reason)
	formerInspection.BaggageDoorWindow = getFunctionResult(inspection.BaggageDoorWindow, formerInspection.BaggageDoorWindow.Reason)
	formerInspection.GearBox = getFunctionResult(inspection.GearBox, formerInspection.GearBox.Reason)
	formerInspection.ShockAbsorber = getFunctionResult(inspection.ShockAbsorber, formerInspection.ShockAbsorber.Reason)
	formerInspection.FrontHighAndLowBeamLight = getFunctionResult(inspection.FrontHighAndLowBeamLight, formerInspection.FrontHighAndLowBeamLight.Reason)
	formerInspection.RearLightAndBrakeLight = getFunctionResult(inspection.RearLightAndBrakeLight, formerInspection.RearLightAndBrakeLight.Reason)
	formerInspection.WiperOperation = getFunctionResult(inspection.WiperOperation, formerInspection.WiperOperation.Reason)
	formerInspection.CarHorn = getFunctionResult(inspection.CarHorn, formerInspection.CarHorn.Reason)
	formerInspection.SideMirrors = getFunctionResult(inspection.SideMirrors, formerInspection.SideMirrors.Reason)
	formerInspection.GeneralBodyCondition = getFunctionResult(inspection.GeneralBodyCondition, formerInspection.GeneralBodyCondition.Reason)

	formerInspection.DriverPerformance = getValidBooleanValue(inspection.DriverPerformance, formerInspection.DriverPerformance)
	formerInspection.Balancing = getValidBooleanValue(inspection.Balancing, formerInspection.Balancing)
	formerInspection.Hazard = getValidBooleanValue(inspection.Hazard, formerInspection.Hazard)
	formerInspection.SignalLightUsage = getValidBooleanValue(inspection.SignalLightUsage, formerInspection.SignalLightUsage)

	formerInspection.Passed = passed

	if changeCount == 0 {
		response.WriteHeader(304)
		res.Inspection = formerInspection
		res.Success = true
		res.Message = " No Change is made! "
		response.Write(helper.MarshalThis(res))
		return
	}

	ctx = context.WithValue(ctx, "inspection", formerInspection)
	// saving the change
	savedInspection, err := h.InspectionSer.UpdateInspection(ctx)
	if err != nil || savedInspection == nil {
		println(err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		res.Inspection = nil
		res.Success = false
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		response.Write(helper.MarshalThis(res))
		return
	}
	response.WriteHeader(http.StatusOK)
	res.Inspection = savedInspection
	res.Success = true
	res.Message = os.Getenv("UPPDATED_SUCCESFULY")
	response.Write(helper.MarshalThis(res))
	return
}

// UpdateInspectionFiles   ... route to update the files that the request may hold
// METHOD : PUT
// INPUT : MULTIPART FORM VALUE
// OUTPUT : JSON
// AUTHORIZATION : INSPECTOR ONLY
func (h *InspectionHandler) UpdateInspectionFiles(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)
	// PARSING THE MULTIPART FORM FILE
	res := &model.InspectionUpdateResponse{
		Success:    false,
		Message:    os.Getenv("BAD_REQUEST_BODY"),
		Inspection: nil,
	}
	err := request.ParseMultipartForm(9999999999999999)
	inspectionID, err := strconv.Atoi(request.FormValue("inspection_id"))
	if err != nil {
		// return that the request is bad request
		response.WriteHeader(400)
		res.Message = os.Getenv("BAD_REQUEST_BODY")
		response.Write(helper.MarshalThis(res))
		return
	}
	ctx = context.WithValue(ctx, "inspector_id", uint(session.ID))
	ctx = context.WithValue(ctx, "inspection_id", uint(inspectionID))
	// checking the whether the inspector is the owner of the inspection or not
	success, erra := h.InspectionSer.IsInspectionOwner(ctx)
	if erra != nil {
		response.WriteHeader(http.StatusNotModified)
		res.Message = os.Getenv("RESOURCE_NOT_FOUND") + "\n Inspection With this ID not Found ..."
		response.Write(helper.MarshalThis(res))
		return
	} else if !success {
		response.WriteHeader(http.StatusUnauthorized)
		res.Message = os.Getenv("UNAUTHORIZED_ACCESS") + "\n You are not authorized to update this Inspection!"
		response.Write(helper.MarshalThis(res))
		return
	}
	// The Inspection os Confirmed that . it is created by the inspector ..
	// getting the variables that are to be included in the updated inspection
	// try fetching the five image files if not any one of them exist , then i will exit the loop
	var a, b, c, d, e bool
	frontImage, frontHeader, eraa := request.FormFile("front_image")
	if eraa == nil {
		a = true
		defer frontImage.Close()
	}
	leftImage, leftHeader, erbb := request.FormFile("left_side_image")
	if erbb == nil {
		b = true
		defer leftImage.Close()
	}
	rightImage, rightHeader, ercc := request.FormFile("right_side_image")
	if ercc == nil {
		c = true
		defer rightImage.Close()
	}
	backImage, backHeader, erdd := request.FormFile("back_image")
	if erdd == nil {
		d = true
		defer backImage.Close()
	}
	signatureImage, signatureHeader, eree := request.FormFile("signature")
	if eree == nil {
		e = true
		defer signatureImage.Close()
	}
	var afilename, bfilename, cfilename, dfilename, efilename string

	if a {
		afilename = helper.GenerateRandomString(10, helper.CHARACTERS) + "." + helper.GetExtension(frontHeader.Filename)
	}
	if b {
		bfilename = helper.GenerateRandomString(10, helper.CHARACTERS) + "." + helper.GetExtension(leftHeader.Filename)
	}
	if c {
		cfilename = helper.GenerateRandomString(10, helper.CHARACTERS) + "." + helper.GetExtension(rightHeader.Filename)
	}
	if d {
		dfilename = helper.GenerateRandomString(10, helper.CHARACTERS) + "." + helper.GetExtension(backHeader.Filename)
	}
	if e {
		efilename = helper.GenerateRandomString(10, helper.CHARACTERS) + "." + helper.GetExtension(signatureHeader.Filename)
	}

	formerInspection, erra := h.InspectionSer.GetInspectionByID(ctx)
	formerFrontImageFilename := formerInspection.FrontImage
	formerLeftImageFilename := formerInspection.LeftSideImage
	formerRightImageFilename := formerInspection.RightSideImage
	formerBackImageFilename := formerInspection.BackImage
	formerSignatureFilename := formerInspection.SignatureImage

	pathToAssetsDirectory := os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY")
	var anewf, bnewf, cnewf, dnewf, enewf *os.File
	var era, erb, erc, erd, ere error
	if a {
		anewf, era = os.Create(pathToAssetsDirectory + "frontImages/" + afilename)
		if era == nil {
			defer anewf.Close()
		}
		_, era = io.Copy(anewf, frontImage)
	}
	if b {
		bnewf, erb = os.Create(pathToAssetsDirectory + "leftImages/" + bfilename)
		if erb == nil {
			defer bnewf.Close()
		}
		_, erb = io.Copy(bnewf, leftImage)
	}
	if c {
		cnewf, erc = os.Create(pathToAssetsDirectory + "rightImages/" + cfilename)
		if erc == nil {
			defer cnewf.Close()
		}
		_, erc = io.Copy(cnewf, rightImage)
	}
	if d {
		dnewf, erd = os.Create(pathToAssetsDirectory + "backImages/" + dfilename)
		if erd == nil {
			defer dnewf.Close()
		}
		_, erd = io.Copy(dnewf, backImage)
	}
	if e {
		enewf, ere = os.Create(pathToAssetsDirectory + "signatureImages/" + efilename)
		if ere == nil {
			defer enewf.Close()
		}
		_, ere = io.Copy(enewf, signatureImage)
	}

	if era != nil || erb != nil || erc != nil || erd != nil || ere != nil {
		response.WriteHeader(http.StatusInternalServerError)
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		res.Success = false
		response.Write(helper.MarshalThis(res))
		return
	}
	if a {
		formerInspection.FrontImage = "public/frontImages/" + afilename
	}
	if b {
		formerInspection.LeftSideImage = "public/leftImages/" + bfilename
	}
	if c {
		formerInspection.RightSideImage = "public/rightImages/" + cfilename
	}
	if d {
		formerInspection.BackImage = "public/backImages/" + dfilename
	}
	if e {
		formerInspection.SignatureImage = "public/signatureImages/" + efilename
	}

	ctx = context.WithValue(ctx, "inspection", formerInspection)
	// saving the change
	savedInspection, err := h.InspectionSer.UpdateInspection(ctx)
	if err != nil || savedInspection == nil {
		// deleting the newly created files
		os.Remove(pathToAssetsDirectory + "frontImages/" + afilename)
		os.Remove(pathToAssetsDirectory + "leftImages/" + bfilename)
		os.Remove(pathToAssetsDirectory + "rightImages/" + cfilename)
		os.Remove(pathToAssetsDirectory + "backImages/" + dfilename)
		os.Remove(pathToAssetsDirectory + "signatureImages/" + efilename)

		println(err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		res.Inspection = nil
		res.Success = false
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		response.Write(helper.MarshalThis(res))
		return
	}
	// the updated inspection is saved succesfully
	os.Remove(pathToAssetsDirectory + strings.TrimPrefix(formerFrontImageFilename, "public/"))
	os.Remove(pathToAssetsDirectory + strings.TrimPrefix(formerLeftImageFilename, "public/"))
	os.Remove(pathToAssetsDirectory + strings.TrimPrefix(formerRightImageFilename, "public/"))
	os.Remove(pathToAssetsDirectory + strings.TrimPrefix(formerBackImageFilename, "public/"))
	os.Remove(pathToAssetsDirectory + strings.TrimPrefix(formerSignatureFilename, "public/"))

	response.WriteHeader(200)
	res.Success = true
	res.Message = os.Getenv("UPPDATED_SUCCESFULY")
	res.Inspection = savedInspection
	response.Write(helper.MarshalThis(res))
}

func (h InspectionHandler) DeleteInspection(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)

	res := &model.SimpleSuccessNotifier{
		Success: false, Message: "Bad Request ",
	}

	inspectionID, era := strconv.Atoi(request.FormValue("inspection_id"))
	if era != nil || inspectionID <= 0 {
		// return that the request is bad request
		response.WriteHeader(400)
		res.Message = os.Getenv("BAD_REQUEST_BODY")
		response.Write(helper.MarshalThis(res))
		return
	}
	ctx = context.WithValue(ctx, "inspection_id", uint(inspectionID))
	ctx = context.WithValue(ctx, "inspector_id", uint(session.ID))
	// checking the whether the inspector is the owner of the inspection or not
	success, erra := h.InspectionSer.IsInspectionOwner(ctx)
	if erra != nil {
		response.WriteHeader(http.StatusNotModified)
		res.Message = os.Getenv("RESOURCE_NOT_FOUND") + "\n Inspection With this ID not Found ..."
		response.Write(helper.MarshalThis(res))
		return
	} else if !success {
		response.WriteHeader(http.StatusUnauthorized)
		res.Message = os.Getenv("UNAUTHORIZED_ACCESS") + "\n You are not authorized to update this Inspection!"
		response.Write(helper.MarshalThis(res))
		return
	}
	// The Inspection os Confirmed that . it is created by the inspector ..

	// calling the delete inspection method to delete  the inspection from the databse
	success, era = h.InspectionSer.DeleteInspection(ctx)
	if !success || era != nil {
		response.WriteHeader(http.StatusInternalServerError)
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		res.Success = false
		response.Write(helper.MarshalThis(res))
		return
	}
	res.Message = "Inspection Deleted Succesfuly!"
	res.Success = true
	response.WriteHeader(204)
	response.Write(helper.MarshalThis(res))
}

// GetInspectionByID method to get inspection by id
func (h *InspectionHandler) GetInspectionByID(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	res := &struct {
		Success    bool              `json:"success"`
		ID         uint              `json:"id"`
		Inspection *model.Inspection `json:"inspection"`
	}{
		Success: false,
	}
	inspectionID, era := strconv.Atoi(params.ByName("id"))
	if era != nil || inspectionID <= 0 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	res.ID = uint(inspectionID)

	ctx = context.WithValue(ctx, "inspection_id", uint(inspectionID))
	inspection, er := h.InspectionSer.GetInspectionByID(ctx)
	if er != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write(helper.MarshalThis(res))
		return
	}
	res.Inspection = inspection
	res.Success = true
	response.WriteHeader(200)
	response.Write(helper.MarshalThis(res))
}

// GetInspectionAsPDf ... methdo to print the inspection data as a pdf format
func (h *InspectionHandler) GetInspectionAsPDF(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	inspectionID, era := strconv.Atoi(params.ByName("id"))
	if era != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := request.Context()
	ctx = context.WithValue(ctx, "inspection_id", uint(inspectionID))
	inspection, era := h.InspectionSer.GetInspectionByID(ctx)
	if era != nil || inspection == nil {
		println("Inspection Not Found ")
		response.WriteHeader(http.StatusNotFound)
		return
	}
	ctx = context.WithValue(ctx, "inspector_id", inspection.InspectorID)
	inspector, era := h.InspectorSer.GetInspectorID(ctx)
	if era != nil || inspector == nil {
		println("Inspector Not Found ")
		response.WriteHeader(http.StatusNotFound)
		return
	}
	ctx = context.WithValue(ctx, "garage_id", uint(inspection.GarageID))
	garage, era := h.GarageSer.GetGarageByID(ctx)
	if era != nil || garage == nil {
		println("Garage Not Found ")
		response.WriteHeader(http.StatusNotFound)
		return
	}

	in := &struct {
		AssetsDirectory string
		Garage          *model.Garage
		Inspection      *model.Inspection
		Inspector       *model.Inspector
	}{
		AssetsDirectory: os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY"),
		Garage:          garage,
		Inspection:      inspection,
		Inspector:       inspector,
	}
	print(in)

	fileDirectory := os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY") + "html/" + helper.GenerateRandomString(5, helper.CHARACTERS) + ".html"
	// craete file to save the image file
	zhtml, er := os.Create(fileDirectory)
	if er != nil || zhtml == nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer zhtml.Close()
	er = h.Template.ExecuteTemplate(zhtml, "inspection.html", in)
	if er != nil {
		os.Remove(fileDirectory)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	pdfFileDirectory := pdf.GetThePdf(fileDirectory)
	if pdfFileDirectory == "" {
		println("  Pdf File Directory ...  ")
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	pdfFile, era := os.Open(pdfFileDirectory)
	if era != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.Copy(response, pdfFile)
}
