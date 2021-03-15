package main

import (
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx"
	"github.com/samuael/Project/CarInspection/pkg/admin"
	"github.com/samuael/Project/CarInspection/pkg/http/rest"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/auth"
	"github.com/samuael/Project/CarInspection/pkg/http/rest/middleware"
	"github.com/samuael/Project/CarInspection/pkg/inspector"
	"github.com/samuael/Project/CarInspection/pkg/secretary"
	pgxstorage "github.com/samuael/Project/CarInspection/pkg/storage/pgx_storage"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

var once sync.Once
var conn *pgx.Conn
var connError error

func main() {
	once.Do(func() {
		conn, connError = pgxstorage.NewStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("CAR_INSPECTION_DB_NAME"))
		if connError != nil {
			os.Exit(1)
		}
	})
	defer conn.Close(context.Background())

	authenticator := auth.NewAuthenticator()
	rules := middleware.NewRules(authenticator)

	adminrepo := pgxstorage.NewAdminRepo(conn)
	secretaryrepo := pgxstorage.NewSecretaryRepo(conn)
	inspectorrepo := pgxstorage.NewInspectorRepo(conn)

	adminservice := admin.NewAdminService(adminrepo)
	secretaryservice := secretary.NewSecretaryService(secretaryrepo)
	inspectorservice := inspector.NewInspectorService(inspectorrepo)

	inspectorhadnler := rest.NewInspectorHandler(authenticator, inspectorservice)
	secretaryhandler := rest.NewSecretaryHandler(authenticator, secretaryservice)
	adminhandler := rest.NewAdminHandler(authenticator, adminservice)
	rest.Route(rules, adminhandler, secretaryhandler, inspectorhadnler)
}
