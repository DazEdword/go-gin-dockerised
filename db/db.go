package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

//Pooler abstracts connection pool methods
type Pooler interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

//Db creates an exported global variable to hold the database connection pool.
var Db Pooler

/*BuildConnectionString builds a Postgresql connection string.
Extracts its different sections from environmental variables*/
func BuildConnectionString() string {

	portValue, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		portValue = "5432"
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), portValue,
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	return connectionString
}

/*BuildConnectionURL builds a Postgresql connection url.
Extracts its different sections from environmental variables.
Format: postgres://user:password@host:port/dbname */
func BuildConnectionURL() string {

	portValue, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		portValue = "5432"
	}

	connectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), portValue,
		os.Getenv("POSTGRES_DB"))

	return connectionURL
}

/*InitDb starts and returns a connection pool*/
func InitDb(connectionString string) Pooler {
	pool, err := pgxpool.Connect(context.Background(), connectionString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Db = pool

	log.Println("Postgres pool created!")
	return Db
}
