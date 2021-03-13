// Package model ...
// this model file  holds structs that are to be used by the admin handler.
package model


// AdminLoginResponse to be usedby the admin response class
type AdminLoginResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Admin   *Admin `json:"admin"`
}
