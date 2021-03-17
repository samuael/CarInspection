package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/samuael/Project/CarInspection/api"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/middleware"
)

// Route returns an http handler for the api.
func Route(rules middleware.Rules, adminhandler IAdminHandler, secretaryhandler ISecretaryHandler, inspectorhandler IInspectorHandler, inspectionhandler IInspectionHandler) http.Handler {
	router := httprouter.New()

	router.POST("/api/admin/login/", adminhandler.AdminLogin)
	router.POST("/api/inspector/login/", inspectorhandler.InspectorLogin)
	router.POST("/api/secretary/login/", secretaryhandler.SecretaryLogin)

	router.GET("/api/logout/", rules.Authenticated(adminhandler.Logout))

	router.POST("/api/secretary/new/", rules.Authorized(rules.Authenticated(secretaryhandler.Create)))
	router.POST("/api/inspector/new/", rules.Authorized(rules.Authenticated(inspectorhandler.CreateInspector)))
	router.POST("/api/inspection/new/", rules.Authorized(rules.Authenticated(inspectionhandler.CreateInspection)))

	router.PUT("/api/inspection/", rules.Authorized(rules.Authenticated(inspectionhandler.EditInspection)))
	router.PUT("/api/inspection/images/", rules.Authorized(rules.Authenticated(inspectionhandler.UpdateInspectionFiles)))
	router.DELETE("/api/inspection/", rules.Authorized(rules.Authenticated(inspectionhandler.DeleteInspection)))
	router.GET("/api/inspector/myinspections/", rules.Authorized(rules.Authenticated(inspectorhandler.GetMyInspections)))
	router.GET("/api/inspection/:id", rules.Authenticated(inspectionhandler.GetInspectionByID))
	router.GET("/inspection/report/:id", rules.Authenticated(inspectionhandler.GetInspectionAsPDF))

	router.PUT("/api/password/new/", rules.Authorized(rules.Authenticated(adminhandler.ChangePassword)))

	http.ListenAndServe(":8080", router)
	return router
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func respHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)
		if err != nil {
			data = struct {
				Error string `json: "error"`
			}{
				Error: err.Error(),
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}
