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
	"/api/admin/logout/": &Permission{
		Roles:   []string{ ADMIN, SECRETARY },
	},
}
