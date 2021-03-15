package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
)

type Rules interface {
	Authenticated(next httprouter.Handle) httprouter.Handle
	Authorized(next httprouter.Handle) httprouter.Handle
	HasPermission(path, role, method string) bool
}

type rules struct {
	auth auth.Authenticator
}

func NewRules(auth auth.Authenticator) Rules {
	return &rules{auth}
}

// LoggedIn simple middleware to push value to the context
func (m rules) Authenticated(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		t, err := m.auth.GetSession(r)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), os.Getenv("CAR_INSPECTION_COOKIE_NAME"), t)
		next(w, r.WithContext(ctx), params)
	})
}

// Authorized checks if a user has proper authority to access a give route
func (m *rules) Authorized(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		session, err := m.auth.GetSession(r)
		if err != nil || session == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		role := session.Role
		permitted := m.HasPermission(r.URL.Path, role, r.Method)
		if !permitted {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		if r.Method == http.MethodPost {
			erro := r.ParseForm()
			if erro != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
		next(w, r, params)
	})
}

func (m *rules) HasPermission(path, role, method string) bool {
	path = strings.ToLower(path)
	role = strings.ToUpper(role)
	method = strings.ToUpper(method)

	if permission := state.Authorities[path]; permission != nil {
		for _, rl := range permission.Roles {
			if strings.ToUpper(rl) == role {
				return true
			}
		}
		return false
	}
	return false
}
