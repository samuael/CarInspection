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
	"github.com/samuael/Project/CarInspection/platforms/hash"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

type IAdminHandler interface {
	AdminLogin(http.ResponseWriter, *http.Request, httprouter.Params)
	Logout(http.ResponseWriter, *http.Request, httprouter.Params)
}

// AdminHandler ...
type AdminHandler struct {
	Authenticator auth.Authenticator
	AdminSer      admin.IAdminService
}

func NewAdminHandler(auths auth.Authenticator, adminser admin.IAdminService) IAdminHandler {
	return &AdminHandler{
		AdminSer:      adminser,
		Authenticator: auths,
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
