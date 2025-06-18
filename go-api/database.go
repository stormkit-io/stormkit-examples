package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB

	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
	schema   string
}

func (d *Database) String() string {
	return fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s`,
		d.host,
		d.port,
		d.user,
		d.password,
		d.dbname,
		d.sslmode,
		d.schema,
	)
}

var dbmux sync.Mutex
var cachedDB *Database

// Connection returns the database connection or an error.
func Connection() (*Database, error) {
	dbmux.Lock()
	defer dbmux.Unlock()

	if cachedDB != nil {
		return cachedDB, nil
	}

	cachedDB = &Database{
		Connection: nil,
		host:       GetString(os.Getenv("POSTGRES_HOST"), "localhost"),
		port:       GetString(os.Getenv("POSTGRES_PORT"), "5432"),
		sslmode:    GetString(os.Getenv("POSTGRES_SSL"), "disable"),
		schema:     GetString(os.Getenv("POSTGRES_SCHEMA"), "public"),
		user:       os.Getenv("POSTGRES_USER"),
		password:   os.Getenv("POSTGRES_PASSWORD"),
		dbname:     os.Getenv("POSTGRES_DB"),
	}

	db, err := sql.Open("postgres", cachedDB.String())

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetConnMaxLifetime(0) // Unlimited
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(2)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	cachedDB.Connection = db

	return cachedDB, nil
}

func (d *Database) Close() {
	if err := d.Connection.Close(); err != nil {
		log.Printf("failed to close database connection: %v", err)
	}
}

// GetString returns the first non-empty string.
func GetString(s ...string) string {
	for _, str := range s {
		if str != "" {
			return str
		}
	}

	return ""
}
