package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
	"github.com/samuael/Project/CarInspection/pkg/login"
	"github.com/samuael/Project/CarInspection/platforms/hash"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

type IAdminHandler interface {
	AdminLogin(http.ResponseWriter, *http.Request, httprouter.Params)
	Logout(http.ResponseWriter, *http.Request, httprouter.Params)
}

// AdminHandler ...
type AdminHandler struct {
	LoginService  login.Service
	Authenticator auth.Authenticator
}

func NewAdminHandler(auths auth.Authenticator, l login.Service) IAdminHandler {
	return &AdminHandler{
		LoginService:  l,
		Authenticator: auths,
	}
}

func (adminhr *AdminHandler) AdminLogin(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")
	var admin *model.Admin

	resp := &model.AdminLoginResponse{}
	resp.Success = false
	resp.Admin = nil
	jdecode := json.NewDecoder(request.Body)
	err := jdecode.Decode(&admin)
	if err != nil || admin.Email == "" || admin.Password == "" {
		response.WriteHeader(http.StatusUnauthorized)
		resp.Message = os.Getenv("INVALID_INPUT")
		response.Write(helper.MarshalThis(resp))
		return
	}
	// _, eer := hash.HashPassword(admin.Password)
	// if eer != nil {
	// 	response.WriteHeader(http.StatusUnauthorized)
	// 	resp.Message = os.Getenv("INVALID_INPUT")
	// 	response.Write(helper.MarshalThis(resp))
	// 	return
	// }
	ctx := request.Context()
	ctx = context.WithValue(ctx, "email", admin.Email)
	newAdmin, err := adminhr.LoginService.AdminByEmail(ctx)
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
		matches := hash.ComparePassword(newAdmin.Password , admin.Password)
		if !matches {
			goto InvalidUsernameOrPassword
		}

		session := &model.Session{
			UserID: admin.ID,
			Email:  newAdmin.Email,
			Role:   state.ADMIN,
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
		resp.Admin = newAdmin
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
func (adminhr *AdminHandler) Logout( response http.ResponseWriter , request *http.Request , params httprouter.Params ){
	adminhr.Authenticator.DeleteSession(response , request  )
}
