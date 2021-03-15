package rest

import (
	// "encoding/json"
	"net/http"
	"os"

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
// INPUT  : MULTIPART FORM FIle
//  Output  :  JSON
/* {
	"success" : "" ,
	"message"  : "" ,
	"inspection"  : {
		"id" : 2 ,

	}
 }*/
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
	input := form.Input{
		Values:  request.Response.Request.MultipartForm.Value,
		VErrors: form.ValidationErrors{},
	}
	request.ParseMultipartForm(999999999999999999)
	input.Required(
		// "garage_id",
		// "inspector_id",
		"driver_name",
		"vehicle_model",
		"vehicle_make",
		"vehicle_color",
		"license_plate",
		// "front_image",
		// "left_side_image",
		// "right_side_image",
		// "back_image",
		// "signature",
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
	frontFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + helper.GetExtension(header.Filename)
	leftFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + helper.GetExtension(bheader.Filename)
	rightFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + helper.GetExtension(cheader.Filename)
	backFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + helper.GetExtension(dheader.Filename)
	signatureFilename := helper.GenerateRandomString(15, helper.CHARACTERS) + helper.GetExtension(eheader.Filename)

	pathToAssetsDirectory := os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY")
	// Creating files for the new comming files
	newFrontImage, era := os.Create(pathToAssetsDirectory + "frontImages/" + frontFilename)
	newLeftImage, erb := os.Create(pathToAssetsDirectory + "leftImages/" + leftFilename)
	newRightmage, erc := os.Create(pathToAssetsDirectory + "rightImages/" + rightFilename)
	newBackImage, erd := os.Create(pathToAssetsDirectory + "backImages/" + backFilename)
	newSignatureImage, ere := os.Create(pathToAssetsDirectory + "signatureImages/" + signatureFilename)

	if era != nil || erb != nil || erc != nil || erd != nil || erd != nil {
		response.WriteHeader(http.StatusInternalServerError)
		input.VErrors.Add("Internal Server Error", " Files Exception ")
		res.Success = false
		res.Message = " Internal Server Error "
		response.Write(helper.MarshalThis(res))
		return
	}

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

func (h InspectionHandler) EditInspection(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// value := params.ByName("id")
	// if value =="" {
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }
}
