package state

type (
	Permission struct {
		Roles   []string
		Methods []string
	}
	Authority map[string]*Permission
)

// Authorities this map represents a map of pathroutes and their permissions
// and roles that are allowed to
var Authorities = Authority{
	"/api/secretary/new/": &Permission{
		Roles: []string{ADMIN},
	},
	"/api/inspector/new/": &Permission{
		Roles: []string{ADMIN},
	},
	"/api/inspection/new/": &Permission{
		Roles: []string{INSPECTOR},
	},
	"/api/inspection/": &Permission{
		Roles: []string{INSPECTOR},
	},
	"/api/inspection/images/": &Permission{
		Roles: []string{INSPECTOR},
	},
	"/api/password/new/": &Permission{
		Roles: []string{INSPECTOR, ADMIN, SECRETARY},
	},
	"/api/inspector/myinspections/": &Permission{
		Roles: []string{INSPECTOR},
	},
	// "/api/inspector/login/": &Permission{
	// 	Roles: []string{INSPECTOR},
	// },
	// "/api/secretary/login/": &Permission{
	// 	Roles: []string{SECRETARY},
	// },
}
