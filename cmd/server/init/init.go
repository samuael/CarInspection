package main

import (
	"database/sql"
	"os"
	"sync"

	"github.com/gchaincl/dotsql"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/storage/sql_db"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load("../.env")
}
var once sync.Once

var conn *sql.DB
var connError error

func main() {
	once.Do(func() {
		conn, connError = sql_db.NewStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("CAR_INSPECTION_DB_NAME"))
		if connError != nil {
			println(connError.Error())
			os.Exit(1)
		}
	})
	admin := &model.Admin{
		Email:           "admin@inspection.com",
		Password:        "$2a$04$cC9mpgc7Z8FuB5T83/QO9.HagJdD9AAFqSBJGX7/C6LvAbx7/5tDe",
		Firstname:       "admin",
		Lastname:        "admiin",
		Middlename:      "admin",
		InspectorsCount: 0,
	}
	// Loads queries from file
	if dot, err := dotsql.LoadFromFile("../../../pkg/constants/query/tables.sql"); err == nil {
		if _, err = dot.Exec(conn, "create-functionality-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_FUNCTIONALITY"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-inspections-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_INSPECTIONS"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-admins-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_ADMINS"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-secretaries-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_SECRETARIS"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-inspectors-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_INSPECTORS"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-address-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_ADDRESSES"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-garage-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_GARAGE"))
			os.Exit(1)
		}
		// insert-admin-table
		if _, err = dot.Exec(conn, "insert-admin-table", admin.Email, admin.Firstname, admin.Middlename, admin.Lastname, admin.Password); err != nil {
			println(err.Error())
			println(os.Getenv("ERROR_INSERTING_DEFAULT_ADMIN"))
			os.Exit(1)
		}
	}
	println("\nDatabase Tables succesfuly Initialized ... \n")
	defer conn.Close()
}
