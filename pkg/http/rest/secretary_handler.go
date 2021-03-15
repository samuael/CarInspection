package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
	"github.com/samuael/Project/CarInspection/pkg/secretary"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

// ISecretaryHandler ... interface
type ISecretaryHandler interface {
	Create(http.ResponseWriter, *http.Request, httprouter.Params)
}

// SecretaryHandler ...
type SecretaryHandler struct {
	Authenticator auth.Authenticator
	SecretSer  secretary.ISecretaryService
}

func NewSecretaryHandler( auths auth.Authenticator , secretser secretary.ISecretaryService) ISecretaryHandler {
	return &SecretaryHandler{
		SecretSer : secretser,
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
	ctx = context.WithValue(ctx  , "email"  , secretary.Email)
	exist := secreth.SecretSer.DoesThisEmailExist(ctx )
	if exist {
		response.WriteHeader(409)
		sucnti.Success= false 
		sucnti.Message = " Secretary With this email Exists "
		response.Write(helper.MarshalThis(sucnti))
		return 
	}

	password := helper.GenerateRandomString(4 , helper.NUMBERS)
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
