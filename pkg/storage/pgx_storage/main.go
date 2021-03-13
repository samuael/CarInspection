package pgx_storage

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx"
)

func NewStorage(username, password, host, dbname string) ( *pgx.Conn , error)  {
	// Preparing the statement 
	postgresStatment := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, dbname )
	conn, err := pgx.Connect(context.Background(), postgresStatment )
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil  , err
	}
	print(" pgx : DB Connected Succesfuly ... \n")
	return conn , err 
}