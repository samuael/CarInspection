package rest

import (
	// "encoding/json"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	// "strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/inspection"
	"github.com/samuael/Project/CarInspection/platforms/form"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

// PostHandler provides access to Post api methods.
type IInspectionHandler interface {
	CreateInspection(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	EditInspection(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	// GetInspections(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// AddInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// DeleteInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	// EditInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type InspectionHandler struct {
	InspectionSer inspection.IInspectionService
}

// NewInspectionHandler ...
func NewInspectionHandler(inser inspection.IInspectionService) IInspectionHandler {
	return &InspectionHandler{
		InspectionSer: inser,
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
	if erra != nil  || formerInspection==nil  {
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

func (h InspectionHandler) DeleteInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// postID, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	// if err != nil || postID == 0 {
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }

	// credentials := r.Context().Value("credentials").(*auth.AppClaims)
	// userID := credentials.ID

	// if allowed := h.u.IsOwnerOfPost(userID, uint(postID)); allowed == false {
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }
	// if err := h.d.DeleteInspection(uint(postID)); err != nil {
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode("post deleted.")
}
