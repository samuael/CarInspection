package main

import (
	// "log"
	"context"
	"os"
	"sync"

	"github.com/jackc/pgx"
	pgxplatform "github.com/samuael/Project/CarInspection/platforms/pgx"
	"github.com/subosito/gotenv"
)



func init() {
	gotenv.Load()
}

var once sync.Once
var conn *pgx.Conn
var connError error

func main(){
	print(os.Getenv("DB_HOST"))
	once.Do(func(){
		conn  , connError = pgxplatform.NewStorage(os.Getenv("DB_USER") , os.Getenv("DB_PASSWORD") , os.Getenv("DB_HOST") , os.Getenv("CAR_INSPECTION_DB_NAME"))
		if connError != nil {
			os.Exit(1)
		}
	})
	defer conn.Close(context.Background())
}