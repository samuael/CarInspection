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
