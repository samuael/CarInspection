package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	// "strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
	"github.com/samuael/Project/CarInspection/pkg/secretary"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

// ISecretaryHandler ... interface
type ISecretaryHandler interface {
	Create(http.ResponseWriter, *http.Request, httprouter.Params)
	SecretaryLogin(response http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteSecretary(response http.ResponseWriter, request *http.Request, parms httprouter.Params)
}

// SecretaryHandler ...
type SecretaryHandler struct {
	Authenticator auth.Authenticator
	SecretSer     secretary.ISecretaryService
}

func NewSecretaryHandler(auths auth.Authenticator, secretser secretary.ISecretaryService) ISecretaryHandler {
	return &SecretaryHandler{
		SecretSer:     secretser,
		Authenticator: auths,
	}
}

// CreateSecretary method to create a secretary
// method post :
// Authorization : Only ROles with admin
// INPUT : JSON
/*
	Input  {
		"email"  : "" ,
		"first_name" :"" ,
		"middle_name" : "" ,
		"last_name" :""
	}

	eg.  StatusCode : 201
	OutPut : {
		"success" : true ,
		"message" : ""
	}
*/
// OutPut : JSON
func (secreth *SecretaryHandler) Create(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	jdecoder := json.NewDecoder(request.Body)

	secretary := &model.Secretary{}
	sucnti := &model.SimpleSuccessNotifier{}

	ctx := request.Context()
	decErr := jdecoder.Decode(secretary)
	if decErr != nil || secretary.Email == "" || secretary.Firstname == "" || secretary.Middlename == "" || secretary.Lastname == "" {
		sucnti.Success = false
		mess := ""
		if decErr == nil {
			if secretary.Email == "" {
				mess += func() string {
					if len(mess) == 0 {
						return " email "
					} else {
						return ", email "
					}
				}()
			}
			if secretary.Firstname == "" {
				mess += func() string {
					if len(mess) == 0 {
						return " first_name "
					} else {
						return ", first_name "
					}
				}()
			}
			if secretary.Email == "" {
				mess += func() string {
					if len(mess) == 0 {
						return " middle_name "
					} else {
						return ", middle_name "
					}
				}()
			}
			if secretary.Email == "" {
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
	ctx = context.WithValue(ctx, "email", secretary.Email)
	exist := secreth.SecretSer.DoesThisEmailExist(ctx)
	if exist {
		response.WriteHeader(409)
		sucnti.Success = false
		sucnti.Message = " Secretary With this email Exists "
		response.Write(helper.MarshalThis(sucnti))
		return
	}

	password := helper.GenerateRandomString(4, helper.NUMBERS)
	secretary.GarageID = adminsess.GarageID
	secretary.Createdby = uint(adminsess.ID)
	secretary.Password = password
	ctx = context.WithValue(ctx, "secretary", secretary)
	secretary, era := secreth.SecretSer.CreateSecretary(ctx)
	if era != nil {
		println(era.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	sucnti.Success = true
	sucnti.Message = " Secretary Created Succesfuly "
	response.WriteHeader(http.StatusCreated)
	response.Write(helper.MarshalThis(sucnti))
	recover()
}

// SecretaryLogin to handle a login request for an admin ....
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
func (secreth *SecretaryHandler) SecretaryLogin(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	var inspector *model.Secretary

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
	newInspector, err := secreth.SecretSer.SecretaryByEmail(ctx)
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
			Role:     state.SECRETARY,
			GarageID: newInspector.GarageID,
		}

		success := secreth.Authenticator.SaveSession(response, session)
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

// DeleteSecretary (response http.ResponseWriter  , request http.Request  , parms httprouter.Params)
// METHOD : DELETE
// INPUT  : INSPECTOR ID AS A PARAMETER
// OUTPUT  : JSON
// AUTHORIZATION : INSPECTOR ONLY
func (secreth *SecretaryHandler) DeleteSecretary(response http.ResponseWriter, request *http.Request, parms httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	secretaryID, era := strconv.Atoi(request.FormValue("secretary_id"))
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)

	if era != nil || secretaryID == 0 {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	res := &struct {
		Success     bool   `json:"success"`
		Message     string `json:"message"`
		SecretaryID uint   `json:"secretary_id"`
	}{
		Success:     false,
		SecretaryID: uint(secretaryID),
	}
	// Is The Admin From Same Garage with secretary
	// chacking whether the admin is from the same garage or not
	// gettingTheSecretary
	ctx = context.WithValue(ctx, "secretary_id", uint(secretaryID))
	secretary, era := secreth.SecretSer.GetSecretaryByID(ctx)
	if era != nil {
		response.WriteHeader(http.StatusNotFound)
		res.Message = " Secretary With Specified ID not found!"
		response.Write(helper.MarshalThis(res))
		return
	}
	if session.GarageID != secretary.GarageID {
		res.Message = fmt.Sprintf("You are Not Authorized to delete Secretary with ID ", secretaryID)
		response.WriteHeader(http.StatusUnauthorized)
		response.Write(helper.MarshalThis(res))
		return
	}
	// deleting the secretary
	era = secreth.SecretSer.DeleteSecretaryByID(ctx)
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

// DeleteSecretary ....  method to delete a secretary
// METHOD : DELETE
// INPUT  : secretary_id  <<< variable
// OUTPUT  : JSON
// AUTHORIZATION : ADMIN ONLY
// func (secreth *SecretaryHandler)   DeleteSecretary(response http.ResponseWriter  , request *http.Request  , params httprouter.Params ){
// 	response.Header().Set("Content-Type", "application")
// 	ctx := request.Context()
// 	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)

// 	secretaryID := strconv.Atoi(request.FormValue("secretary_id"))

// }
