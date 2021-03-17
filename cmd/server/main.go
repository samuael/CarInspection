package main

import (
	// "context"
	"html/template"
	"os"
	"sync"

	// "github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
	"github.com/samuael/Project/CarInspection/pkg/admin"
	"github.com/samuael/Project/CarInspection/pkg/garage"
	"github.com/samuael/Project/CarInspection/pkg/http/rest"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/middleware"
	"github.com/samuael/Project/CarInspection/pkg/inspection"
	"github.com/samuael/Project/CarInspection/pkg/inspector"
	"github.com/samuael/Project/CarInspection/pkg/secretary"
	pgxstorage "github.com/samuael/Project/CarInspection/pkg/storage/pgx_storage"
	// "github.com/samuael/Project/CarInspection/platforms/pdf"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

var once sync.Once
var conn *pgxpool.Pool
var connError error

var templates *template.Template

func main() {
	once.Do(func() {
		conn, connError = pgxstorage.NewStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("CAR_INSPECTION_DB_NAME"))
		if connError != nil {
			os.Exit(1)
		}
		templates = template.Must(template.ParseGlob(os.Getenv("PATH_TO_TEMPLATES") + "*.html"))
	})
	defer conn.Close()

	authenticator := auth.NewAuthenticator()
	rules := middleware.NewRules(authenticator)

	adminrepo := pgxstorage.NewAdminRepo(conn)
	secretaryrepo := pgxstorage.NewSecretaryRepo(conn)
	inspectorrepo := pgxstorage.NewInspectorRepo(conn)
	inspectionrepo := pgxstorage.NewInspectionRepo(conn)
	garagerepo := pgxstorage.NewGarageRepo(conn)

	garageservice := garage.NewGarageService(garagerepo)
	adminservice := admin.NewAdminService(adminrepo)
	secretaryservice := secretary.NewSecretaryService(secretaryrepo)
	inspectorservice := inspector.NewInspectorService(inspectorrepo)
	inspectionservice := inspection.NewInspectionService(inspectionrepo)

	inspectionhandler := rest.NewInspectionHandler(inspectionservice, templates, inspectorservice, garageservice)
	inspectorhadnler := rest.NewInspectorHandler(authenticator, inspectorservice)
	secretaryhandler := rest.NewSecretaryHandler(authenticator, secretaryservice)
	adminhandler := rest.NewAdminHandler(authenticator, adminservice, secretaryservice, inspectorservice)
	rest.Route(rules, adminhandler, secretaryhandler, inspectorhadnler, inspectionhandler)
}
