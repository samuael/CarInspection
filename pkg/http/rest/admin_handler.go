package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/login"
	"github.com/samuael/Project/CarInspection/platforms/hash"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)



type IAdminHandler interface{
	AdminLogin(http.ResponseWriter , *http.Request  , httprouter.Params )
}

// AdminHandler ...
type AdminHandler struct {
	LoginService login.Service
}


func NewAdminHandler(l login.Service) IAdminHandler { 
	return &AdminHandler{
		LoginService : l ,
	}
}

func (adminhr *AdminHandler) AdminLogin(response http.ResponseWriter , request *http.Request  , params httprouter.Params){
	response.Header().Set("Content-Type", "application/json")
	var admin *model.Admin

	resp := &model.AdminLoginResponse{}
	resp.Success= false 
	resp.Admin = nil 
	jdecode := json.NewDecoder(request.Body)
	err := jdecode.Decode(&admin)
	if err != nil || admin.Email=="" || admin.Password==""  {
		response.WriteHeader(http.StatusUnauthorized)
		resp.Message = " Invalid  Input "
		response.Write(helper.MarshalThis(resp))
		return 
	}
	hash  , eer := hash.HashPassword(admin.Password)
	if eer != nil {	
		response.WriteHeader(http.StatusUnauthorized)
		resp.Message = " Invalid  Input "
		response.Write(helper.MarshalThis(resp))
		return 
	}

	ctx := request.Context()
	ctx = context.WithValue(ctx ,"email" , admin.Email )
	ctx = context.WithValue(ctx ,"password" ,  hash )
	admin  , err = adminhr.LoginService.AdminLogin(ctx)
	if err != nil  {
		resp.Success = false
		resp.Message= " No Admin Found By this id "
		response.WriteHeader(401)
		response.Write(helper.MarshalThis(resp))
		return 
	}
	if admin == nil {
		goto InvalidUsernameOrPassword
	}else {
		admin = ctx.Value("admin").(*model.Admin)
		if admin==nil {
			goto InvalidUsernameOrPassword
		}
		resp.Success = true
		resp.Message = state.SuccesfulyLoggedIn
		resp.Admin = admin
		response.WriteHeader(200)
		response.Write(helper.MarshalThis(resp))
		return 
	}
	InvalidUsernameOrPassword : {
		resp.Success = false
		resp.Message= state.InvalidUsernameORPassword
		response.WriteHeader(401)
		response.Write(helper.MarshalThis(resp))	
		return
	};
	
}