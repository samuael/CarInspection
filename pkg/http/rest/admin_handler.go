package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/admin"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
	"github.com/samuael/Project/CarInspection/pkg/inspector"
	"github.com/samuael/Project/CarInspection/pkg/secretary"
	"github.com/samuael/Project/CarInspection/platforms/hash"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

type IAdminHandler interface {
	AdminLogin(http.ResponseWriter, *http.Request, httprouter.Params)
	Logout(http.ResponseWriter, *http.Request, httprouter.Params)
	ChangePassword(response http.ResponseWriter, request *http.Request, params httprouter.Params)
}

// AdminHandler ...
type AdminHandler struct {
	Authenticator auth.Authenticator
	AdminSer      admin.IAdminService
	InspectorSer  inspector.IInspectorService
	Secretser     secretary.ISecretaryService
}

func NewAdminHandler(auths auth.Authenticator, adminser admin.IAdminService, secretser secretary.ISecretaryService, inspectorser inspector.IInspectorService) IAdminHandler {
	return &AdminHandler{
		AdminSer:      adminser,
		Authenticator: auths,
		InspectorSer:  inspectorser,
		Secretser:     secretser,
	}
}

// AdminLogin to handle a login request for an admin ....
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
func (adminhr *AdminHandler) AdminLogin(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	var admin *model.Admin

	resp := &model.LoginResponse{}
	resp.Success = false
	resp.User = nil
	jdecode := json.NewDecoder(request.Body)
	err := jdecode.Decode(&admin)
	if err != nil || admin.Email == "" || admin.Password == "" {
		response.WriteHeader(http.StatusUnauthorized)
		resp.Message = os.Getenv("INVALID_INPUT")
		response.Write(helper.MarshalThis(resp))
		return
	}
	ctx := request.Context()
	ctx = context.WithValue(ctx, "email", admin.Email)
	newAdmin, err := adminhr.AdminSer.AdminByEmail(ctx)
	if err != nil {
		resp.Success = false
		resp.Message = " No Record Found By this id "
		response.WriteHeader(401)
		response.Write(helper.MarshalThis(resp))
		return
	} else {
		if newAdmin == nil {
			goto InvalidUsernameOrPassword
		}

		// comparing the hashed password and the password
		matches := hash.ComparePassword(newAdmin.Password, admin.Password)
		if !matches {
			goto InvalidUsernameOrPassword
		}

		session := &model.Session{
			ID:       newAdmin.ID,
			Email:    newAdmin.Email,
			Role:     state.ADMIN,
			GarageID: newAdmin.GarageID,
		}

		success := adminhr.Authenticator.SaveSession(response, session)
		if !success {
			resp.Message = os.Getenv("INTERNAL_SERVER_ERROR")
			resp.Success = false
			response.WriteHeader(http.StatusInternalServerError)
			response.Write(helper.MarshalThis(resp))
			return
		}
		resp.Success = true
		resp.Message = state.SuccesfulyLoggedIn
		resp.User = newAdmin
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

// Logout || method GET /for an admin to log out
func (adminhr *AdminHandler) Logout(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	adminhr.Authenticator.DeleteSession(response, request)
}

// ChangePassword ... method to change the password for all the three roles
// METHOD  : PUT
// INPUT : JSON
/*
	{
		"old_password" : "theoldpassword" ,
		"new_password" : "new_password " ,
		"confirm_password" : "new_password_here"
	}

	OUTPUT : JSON

	{
		"success" : true ,
		"message" : "Password changed succesfuly "
	}
*/
func (adminhr *AdminHandler) ChangePassword(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	ctx := request.Context()
	session := ctx.Value(os.Getenv("CAR_INSPECTION_COOKIE_NAME")).(*model.Session)

	res := &model.SimpleSuccessNotifier{
		Success: false,
	}
	input := &struct {
		Oldpassword     string `json:"old_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	jdecoder := json.NewDecoder(request.Body)
	era := jdecoder.Decode(input)
	if era != nil || input.Oldpassword == "" || input.NewPassword == "" || input.ConfirmPassword == "" {
		response.WriteHeader(http.StatusBadRequest)
		res.Message = os.Getenv("BAD_REQUEST_BODY")
		response.Write(helper.MarshalThis(res))
		return
	}
	if input.ConfirmPassword != input.NewPassword {
		res.Message = os.Getenv("RE_CONFIRM_PASSWORD")
		response.Write(helper.MarshalThis(res))
		return
	}
	if len(input.NewPassword) < 4 {
		response.WriteHeader(http.StatusBadRequest)
		res.Message = "Password Length Must exceed 4 characters! "
		response.Write(helper.MarshalThis(res))
		return
	}
	var changesuccess bool
	ctx = context.WithValue(ctx, "user_id", uint(session.ID))

	switch session.Role {
	case state.ADMIN:
		{
			hashed, era := hash.HashPassword(input.NewPassword)
			if era != nil {
				res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
				res.Success = false
				response.WriteHeader(http.StatusInternalServerError)
				response.Write(helper.MarshalThis(res))
				return
			}
			ctx = context.WithValue(ctx, "password", hashed)
			changesuccess, era = adminhr.AdminSer.ChangePassword(ctx)
		}
	case state.INSPECTOR:
		{
			ctx = context.WithValue(ctx, "password", input.NewPassword)
			changesuccess , era = adminhr.InspectorSer.ChangePassword(ctx)
		}
	case state.SECRETARY:
		{
			ctx = context.WithValue(ctx, "password", input.NewPassword)
			changesuccess , era = adminhr.Secretser.ChangePassword(ctx)

		}
	}

	if era != nil || !changesuccess {
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		res.Success = false
		response.WriteHeader(http.StatusInternalServerError)
		response.Write(helper.MarshalThis(res))
	}
	response.WriteHeader(200)
	res.Message = "Password Changed Succesfuly!"
	res.Success = true
	response.Write(helper.MarshalThis(res))	
}


// 