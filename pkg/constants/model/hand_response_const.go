// Package model ...
// this model file  holds structs that are to be used by the admin handler.
package model

import "github.com/samuael/Project/CarInspection/platforms/form"

// AdminLoginResponse to be usedby the admin response class
type AdminLoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Admin   *Admin `json:"admin"`
}

// LoginResponse to be usedby the admin response class
type LoginResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	User    interface{} `json:"user"`
}

// SimpleSuccessNotifier ...
type SimpleSuccessNotifier struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// InspectionCreationResponse ....
type InspectionCreationResponse struct {
	Success    bool                  `json:"success"`
	HasError   bool                  `json:"has_error"`
	Message    string                `json:"message"`
	Inspection *Inspection           `json:"inspection"`
	Errors     form.ValidationErrors `json:"errors"`
}

// InspectionUpdateResponse   model to represent output message
type InspectionUpdateResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Inspection *Inspection `json:"inspection"`
}
