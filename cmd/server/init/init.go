package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/gchaincl/dotsql"
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/pkg/constants/state"
	"github.com/samuael/Project/CarInspection/pkg/storage/sql_db"

	// "github.com/samuael/Project/CarInspection/platforms/hash"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load("../.env")
}

var once sync.Once

var conn *sql.DB
var connError error

// func main() {
// 	hashes, er := hash.HashPassword("asmin")
// 	if er != nil {
// 		print("Error")
// 	} else {
// 		print(hashes)
// 	}
// }

func main() {
	once.Do(func() {
		conn, connError = sql_db.NewStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("CAR_INSPECTION_DB_NAME"))
		if connError != nil {
			println(connError.Error())
			os.Exit(1)
		}
	})
	garage := TakeDefaultGarageData()
	admin := &model.Admin{
		Email:           "admin@inspection.com",
		Password:        "$2a$04$cC9mpgc7Z8FuB5T83/QO9.HagJdD9AAFqSBJGX7/C6LvAbx7/5tDe",
		Firstname:       "admin",
		Lastname:        "admiin",
		Middlename:      "admin",
		InspectorsCount: 0,
		GarageID:        1,
	}
	// Loads queries from file
	if dot, err := dotsql.LoadFromFile("../../../pkg/constants/query/tables.sql"); err == nil {
		if _, err = dot.Exec(conn, "create-functionality-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_FUNCTIONALITY"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-address-table"); err != nil {
			println(err.Error())
			println(os.Getenv("ERROR_CREATING_TABLE_ADDRESSES"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-garage-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_GARAGE"))
			os.Exit(1)
		}
		if _, err = dot.Exec(conn, "create-inspections-table"); err != nil {
			println(os.Getenv("ERROR_CREATING_TABLE_INSPECTIONS"))
			println(err.Error())
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
		// insert-ADDRESS-table
		if _, err = dot.Exec(conn, "insert-address-table", garage.Address.Country, garage.Address.Region, garage.Address.Zone, garage.Address.Woreda, garage.Address.City, garage.Address.Kebele); err != nil {
			println(os.Getenv("ERROR_INSERTING_ADDRESS_DATA"))
			os.Exit(1)
		}
		// insert-GARAGE-table
		if _, err = dot.Exec(conn, "insert-garage-table", garage.Name, 1); err != nil {
			println(os.Getenv("ERROR_INSERTING_GARAGE_DATA"))
			os.Exit(1)
		}
		// insert-admin-table
		if _, err = dot.Exec(conn, "insert-admin-table", admin.Email, admin.Firstname, admin.Middlename, admin.Lastname, admin.Password, admin.GarageID); err != nil {
			println(os.Getenv("ERROR_INSERTING_DEFAULT_ADMIN"))
			os.Exit(1)
		}
		for _, fr := range state.FunctionalityResultInstances {
			// insert-functionality-results-table
			if _, err = dot.Exec(conn, "insert-functionality-results-table", fr.Result, fr.Reason); err != nil {
				println(os.Getenv("ERROR_INSERTING_FUNCTIONALITY_RESULTS_DATA"), "   At Index   "+strconv.Itoa(int(fr.ID)))
				os.Exit(1)
			}
		}
	}
	println("\nDatabase Tables succesfuly Initialized ... \n")
	defer conn.Close()
}

// Taking Default Garage Data for the first Instance of Garage
func TakeDefaultGarageData() *model.Garage {
	garage := &model.Garage{}
	valid := false
	for !valid {
		print("\nEnter Default Garage Name :   ")
		var garagename string
		fmt.Scan(&garagename)

		println(" -------------------- Fill the Address Information of the Garage below --------------------------- ")

		print("\nEnter Country :   ")
		// taking the garage address
		var country = ""

		fmt.Scanln(&country)
		print("\nEnter Region :   ")
		var region = ""
		fmt.Scanln(&region)
		print("\nEnter Zone :   ")
		var zone = ""
		fmt.Scanln(&zone)

		print("\nEnter Woreda :   ")
		var woreda = ""
		fmt.Scanln(&woreda)

		print("\nEnter City :   ")
		var city = ""
		fmt.Scanln(&city)

		print("\nEnter Kebele :   ")
		var kebele = ""
		fmt.Scanln(&kebele)
		if !(kebele == "" || woreda == "" || zone == "" || region == "" || country == "") {
			valid = true
			garage = &model.Garage{
				Name: garagename,
				Address: &model.Address{
					ID:      1,
					Country: country,
					Region:  region,
					Zone:    zone,
					Woreda:  woreda,
					Kebele:  kebele,
					City:    city,
				},
			}
			println()
		} else {
			println("\n\n---------------------------Re-Taking --------------------------------\n\n")
		}

	}
	return garage
}
