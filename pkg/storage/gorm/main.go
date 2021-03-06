package gorm
import (
	"fmt"
	// _ "github.com/lib/pq"
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var dbs *sql.DB

var postgresStatmente string
var errors error

// NewStorage ...
func NewStorage(username, password, host, dbname string) (*gorm.DB, error) {
	// Preparing the statmente
	postgresStatmente = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, dbname)
	db, errors = gorm.Open("postgres", postgresStatmente)
	if errors != nil {
		panic(errors)
	}
	fmt.Println("DB Connected Succesfully ")
	return db, nil
}
