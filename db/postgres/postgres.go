package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mskarbe/go-gql-api-server/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	psql "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Postgres struct {
	Database *sql.DB
	name string
}

// Connect() initalises database session
func (pg *Postgres) Connect(c *config.DbConfig) {
	var err error

	params := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    c.Host, c.Port, c.User, c.Password, c.DbName)

	pg.name = c.DbName

	// open the connection
	if pg.Database, err = sql.Open(c.Driver, params); err != nil {
  		log.Fatal(err)
	}

	// ensure that it is reponsive
	if err = pg.Database.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database at "+c.User+"@"+c.Host+":"+c.Port+"/"+c.DbName)

}

// Close connection 
func (pg *Postgres) Close() {
	log.Println("Closing connection")
	pg.Database.Close()
}

// Migrate schema
func (pg *Postgres) Migrate() {
	var err error

	log.Println("Checking connection")
	if err = pg.Database.Ping(); err != nil {
		log.Fatal(err)
	}

	// setup driver for migrate 
	driver, _ := psql.WithInstance(pg.Database, &psql.Config{})
    m, _ := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
        pg.name, driver)

	log.Println("Starting migrations...")
	// ErrNoChange: indicates that the schema is up-to-date
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	log.Println("Schema migrated")
	 
}
