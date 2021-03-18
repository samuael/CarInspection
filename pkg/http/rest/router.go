package rest

import (
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	_ "github.com/samuael/Project/CarInspection/api"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/middleware"
)

// Route returns an http handler for the api.
func Route(rules middleware.Rules, adminhandler IAdminHandler, secretaryhandler ISecretaryHandler, inspectorhandler IInspectorHandler, inspectionhandler IInspectionHandler) http.Handler {
	router := httprouter.New()

	router.NotFound = http.StripPrefix("/public/", http.FileServer(http.Dir(os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY"))))
	router.GET("/public/", FilterDirectory(http.StripPrefix("/public/", router.NotFound)))

	router.POST("/api/admin/login/", accessControl(adminhandler.AdminLogin))
	router.GET("/api/admin/inspectors/", rules.Authorized(rules.Authenticated(accessControl(inspectorhandler.GetMyInspectors))))

	router.POST("/api/inspector/login/", accessControl(inspectorhandler.InspectorLogin))
	router.POST("/api/secretary/login/", accessControl(secretaryhandler.SecretaryLogin))
	router.GET("/api/logout/", rules.Authenticated(accessControl(adminhandler.Logout)))

	router.DELETE("/api/secretary/", rules.Authorized(rules.Authenticated(accessControl(secretaryhandler.DeleteSecretary))))
	router.POST("/api/secretary/new/", rules.Authorized(rules.Authenticated(accessControl(secretaryhandler.Create))))

	router.DELETE("/api/inspector/", accessControl(inspectorhandler.DeleteInspectorByID))
	router.POST("/api/inspector/new/", rules.Authorized(rules.Authenticated(accessControl(inspectorhandler.CreateInspector))))
	router.PUT("/inspector/profile/image/new/", rules.Authorized(rules.Authenticated(accessControl(inspectorhandler.InspectorProfileImageChange))))
	router.POST("/api/inspection/new/", rules.Authorized(rules.Authenticated(accessControl(inspectionhandler.CreateInspection))))

	router.PUT("/api/inspection/", rules.Authorized(rules.Authenticated(accessControl(inspectionhandler.EditInspection))))
	router.PUT("/inspection/images/", rules.Authorized(rules.Authenticated(accessControl(inspectionhandler.UpdateInspectionFiles))))
	router.DELETE("/api/inspection/", rules.Authorized(rules.Authenticated(accessControl(inspectionhandler.DeleteInspection))))
	router.GET("/api/inspector/myinspections/", rules.Authorized(rules.Authenticated(accessControl(inspectorhandler.GetMyInspections))))
	router.GET("/api/inspection/:id", rules.Authenticated(accessControl(inspectionhandler.GetInspectionByID)))
	router.GET("/inspection/report/:id", rules.Authenticated(accessControl(inspectionhandler.GetInspectionAsPDF)))

	router.GET("/api/inspections/search/", rules.Authenticated(accessControl(inspectionhandler.SearchInspections)))
	router.GET("/api/inspections/search/:inspectorID", rules.Authenticated(accessControl(inspectionhandler.SearchInspections)))

	router.PUT("/api/password/new/", rules.Authorized(rules.Authenticated(accessControl(adminhandler.ChangePassword))))

	http.ListenAndServe(":8080", router)
	return router
}

func accessControl(h httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		h(w, r, params)
	})
}

func FilterDirectory(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if strings.HasSuffix(r.URL.Path, "/") {
			return
		}
		next.ServeHTTP(w, r)
	}
}
