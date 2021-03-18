package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
	"github.com/samuael/Project/CarInspection/pkg/inspector"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

type IInspectorHandler interface {
	CreateInspector(http.ResponseWriter, *http.Request, httprouter.Params)
	InspectorLogin(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetMyInspections(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	InspectorProfileImageChange(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteInspectorByID(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetMyInspectors(response http.ResponseWriter, request *http.Request, params httprouter.Params)
}

// InspetorHandler inspector handler for
// handling inspector related functionalities
type InspectorHandler struct {
	Authenticator auth.Authenticator
	InspectorSer  inspector.IInspectorService
}

func NewInspectorHandler(authenticator auth.Authenticator, inser inspector.IInspectorService) IInspectorHandler {
	return &InspectorHandler{
		Authenticator: authenticator,
		InspectorSer:  inser,
	}
}

// CreateInspector  method to create new instance of Inspector
// METHOD : POST
/*
	INPUT : JSON

	OUTPUT : JSON

*/
// AUTHORIZATION : ADMINS ONLY
func (insorh *InspectorHandler) CreateInspector(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	jdecoder := json.NewDecoder(request.Body)
	inspectr := &model.Inspector{}
	sucnti := &model.SimpleSuccessNotifier{}

	ctx := request.Context()
	decErr := jdecoder.Decode(inspectr)
	if decErr != nil || inspectr.Email == "" || inspectr.Firstname == "" || inspectr.Middlename == "" || inspectr.Lastname == "" {
		sucnti.Success = false
		mess := ""
		if decErr == nil {
			if inspectr.Email == "" {
				mess += func() string {
					if len(mess) == 0 {
						return " email "
					} else {
						return ", email "
					}
				}()
			}
			if inspectr.Firstname == "" {
				mess += func() string {
					if len(mess) == 0 {
						return " first_name "
					} else {
						return ", first_name "
					}
				}()
			}
			if inspectr.Email == "" {
				mess += func() string {
					if len(mess) == 0 {
						return " middle_name "
					} else {
						return ", middle_name "
					}
				}()
			}
			if inspectr.Email == "" {
				mess += func() string {
					if len(mess) == 0 {
						return " last_name "
					} else {
						return ", last_name "
					}
				}()
			}
			if len(mess) > 0 {
				mess = " Missing  : " + mess
			}
		}
		sucnti.Message = fmt.Sprintf(os.Getenv("ERROR_INVALID_INPUT"), mess)
		response.WriteHeader(http.StatusBadRequest)
		response.Write(helper.MarshalThis(sucnti))
		return
	}
	adminsess := (ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session))
	// print(adminsess.ID , adminsess.Email , adminsess.GarageID , adminsess.Role )
	if adminsess == nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	// checking whether the email exist or not in the database

	//  Adding the email in the context values list and passing the context to the method below
	ctx = context.WithValue(ctx, "email", inspectr.Email)
	exist := insorh.InspectorSer.DoesThisEmailExist(ctx)
	if exist {
		response.WriteHeader(409)
		sucnti.Success = false
		sucnti.Message = " Inspector With this email Exists "
		response.Write(helper.MarshalThis(sucnti))
		return
	}

	password := helper.GenerateRandomString(4, helper.NUMBERS)
	inspectr.GarageID = adminsess.GarageID
	inspectr.InspectionCount = 0
	inspectr.Password = password
	inspectr.Imageurl = ""
	inspectr.Createdby = uint(adminsess.ID)
	ctx = context.WithValue(ctx, "inspector", inspectr)
	inspectr, era := insorh.InspectorSer.CreateInspector(ctx)
	if era != nil || inspectr == nil {
		println(era.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	sucnti.Success = true
	sucnti.Message = " Inspector Created Succesfuly "
	response.WriteHeader(http.StatusCreated)
	response.Write(helper.MarshalThis(sucnti))
	recover()
}

// InspectorLogin to handle a login request for an admin ....
// METHOD : POST
// INPUT  : JSON
/*
	INPUT : {
		"email"  : "email" ,
		"password"  : "passs"
	}

	OUTPUT : {
		"success" : true ,
		"message" : "Success message" ,
		"admin" : {
			"id" : 3 ,
			"email" : ""
		}
	}
*/
func (insorh *InspectorHandler) InspectorLogin(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	var inspector *model.Inspector

	resp := &model.LoginResponse{}
	resp.Success = false
	resp.User = nil
	jdecode := json.NewDecoder(request.Body)
	err := jdecode.Decode(&inspector)
	if err != nil || inspector.Email == "" || inspector.Password == "" {
		response.WriteHeader(http.StatusUnauthorized)
		resp.Message = os.Getenv("INVALID_INPUT")
		response.Write(helper.MarshalThis(resp))
		return
	}
	ctx := request.Context()
	ctx = context.WithValue(ctx, "email", inspector.Email)
	newInspector, err := insorh.InspectorSer.InspectorByEmail(ctx)
	if err != nil {
		resp.Success = false
		resp.Message = " No Record Found By this id "
		response.WriteHeader(401)
		response.Write(helper.MarshalThis(resp))
		return
	} else {
		if newInspector == nil {
			goto InvalidUsernameOrPassword
		}

		// comparing the hashed password and the password
		// matches := hash.ComparePassword(newAdmin.Password, admin.Password)
		matches := newInspector.Password == inspector.Password
		if !matches {
			goto InvalidUsernameOrPassword
		}

		session := &model.Session{
			ID:       uint64(newInspector.ID),
			Email:    newInspector.Email,
			Role:     state.INSPECTOR,
			GarageID: newInspector.GarageID,
		}

		success := insorh.Authenticator.SaveSession(response, session)
		if !success {
			resp.Message = os.Getenv("INTERNAL_SERVER_ERROR")
			resp.Success = false
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(helper.MarshalThis(resp))
			return
		}
		resp.Success = true
		resp.Message = state.SuccesfulyLoggedIn
		resp.User = newInspector
		response.WriteHeader(200)
		response.Write(helper.MarshalThis(resp))
		return
	}
InvalidUsernameOrPassword:
	{
		resp.Success = false
		resp.Message = state.InvalidUsernameORPassword
		response.WriteHeader(401)
		response.Write(helper.MarshalThis(resp))
		return
	}
}

// GetMyInspections  returns list of inspections that an inspector took
/*
	OUTPUT : JSON

	{
		"success" : true  ,
		"message" : "succesfuly fetched -- messages "
		"inspections" : [
			{
				id : 1,
				.
				.
				.
			}
		]
	}

*/
func (insorh *InspectorHandler) GetMyInspections(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)

	res := &struct {
		Success     bool   `json:"success"`
		Message     string `json:"message"`
		Inspections []*model.Inspection
	}{
		Success: false,
		Message: "",
	}

	ctx = context.WithValue(ctx, "inspector_id", uint(session.ID))

	inspections, er := insorh.InspectorSer.GetInspectionsByInspectorID(ctx)
	if er != nil || inspections == nil {
		res.Message = "Unable to Found Inspections For Specified User!"
		res.Success = false
		response.WriteHeader(404)
		res.Inspections = []*model.Inspection{}
		response.Write(helper.MarshalThis(res))
	}
	res.Success = true
	res.Inspections = inspections
	res.Message = fmt.Sprintf(" Succesfuly Found %d  Inspections ", len(inspections))
	response.WriteHeader(200)
	response.Write(helper.MarshalThis(res))
}

// InspectorProfileImageChange (response http.ResponseWriter, request *http.Request, params httprouter.Params)
// METHOD : PUT
// INPUT : MULTIPART FORM DATA
// OUTPUT : JSON of the inspector Profile
// INPUT PARAMETER NAME : profile_image
func (insorh *InspectorHandler) InspectorProfileImageChange(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)
	// Since i have the inspector Session i don't have to check the presence of the Inspector
	// getting the image files
	res := &struct {
		Success   bool             `json:"success"`
		Inspector *model.Inspector `json:"inspector"`
	}{
		Success: false,
	}
	// parsing the formdata since it is a multipart form data
	request.ParseMultipartForm(999999999999)
	file, header, era := request.FormFile("image")
	if era != nil || file == nil || header == nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	// get the inspector Instance from the database
	ctx = context.WithValue(ctx, "inspector_id", uint(session.ID))
	inspector, era := insorh.InspectorSer.GetInspectorByID(ctx)
	if era != nil || inspector == nil {
		response.WriteHeader(http.StatusNotModified)
		return
	}
	// "public/inspectorsImage/"+
	randomFilename := helper.GenerateRandomString(5, helper.CHARACTERS) + "." + helper.GetExtension(header.Filename)
	newFile, er := os.Create(os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY") + "inspectors/" + randomFilename)
	if er != nil {
		println(er.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer newFile.Close()
	_, er = io.Copy(newFile, file)
	if er != nil {
		println(er.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	oldImage := inspector.Imageurl
	inspector.Imageurl = "public/inspectors/" + randomFilename
	ctx = context.WithValue(ctx, "inspector", inspector)
	er = insorh.InspectorSer.UpdateProfileImage(ctx)

	if er != nil {
		println(er.Error())
		response.WriteHeader(http.StatusInternalServerError)
		os.Remove(os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY") + "inspectors/" + randomFilename)
		return
	}
	os.Remove(os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY") + strings.TrimPrefix(oldImage, "public/"))
	response.WriteHeader(200)
	res.Success = true
	res.Inspector = inspector
	response.Write(helper.MarshalThis(res))
}

// DeleteInspectorByID (response http.ResponseWriter, request *http.Request, params httprouter.Params)
func (insorh *InspectorHandler) DeleteInspectorByID(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)
	// Since i have the inspector Session i don't have to check the presence of the Inspector
	// getting the image files
	res := &struct {
		Success     bool   `json:"success"`
		Message     string `json:"message"`
		InspectorID uint   `json:"inspector_id"`
	}{
		Success: false,
	}
	// Getting the inspector ID from the request
	inspectorID, era := strconv.Atoi(request.FormValue("inspector_id"))
	if era != nil || inspectorID == 0 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	res.InspectorID = uint(inspectorID)

	ctx = context.WithValue(ctx, "inspector_id", uint(inspectorID))
	inspector, era := insorh.InspectorSer.GetInspectorByID(ctx)
	if era != nil {
		response.WriteHeader(http.StatusNotFound)
		res.Message = " Inspector With Specified ID not found!"
		response.Write(helper.MarshalThis(res))
		return
	}
	if session.GarageID != inspector.GarageID {
		res.Message = fmt.Sprintf("You are Not Authorized to delete Inspector with ID ", inspectorID)
		response.WriteHeader(http.StatusUnauthorized)
		response.Write(helper.MarshalThis(res))
		return
	}
	// deleting the secretary
	era = insorh.InspectorSer.DeleteInspectorByID(ctx)
	if era != nil {
		response.WriteHeader(http.StatusInternalServerError)
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		response.Write(helper.MarshalThis(res))
		return
	}
	res.Message = "Deletion was succesful!"
	res.Success = true
	response.WriteHeader(http.StatusOK)
	response.Write(helper.MarshalThis(res))
}

// GetMyInspectors ... a method to get the list of inspectors the admin have created
// METHOD : GET
// INPUT  : --
// OUTPUT : JSON
// AUTHORIZATION : ADMINS ONLY
func (insorh *InspectorHandler) GetMyInspectors(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)

	// putting the inspector ID in the context
	ctx = context.WithValue(ctx, "admin_id", uint(session.ID))
	inspectors, err := insorh.InspectorSer.GetInspectorsByAdminID(ctx)

	res := &struct {
		Success    bool               `json:"success"`
		Inspectors []*model.Inspector `json:"inspectors"`
	}{
		Success:    false,
		Inspectors: nil,
	}
	if err != nil || inspectors == nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write(helper.MarshalThis(res))
		return
	}
	response.WriteHeader(http.StatusOK)
	res.Inspectors = inspectors
	res.Success = true
	response.Write(helper.MarshalThis(res))
}
