package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mashun4ek/study_group/app/models"
)

// DatabaseInterface implements following methods
type DatabaseInterface interface {
	CreateNewProfile(userInput *models.Profile) error
	DeleteProfile(userID string) error
	GetAllProfiles() ([]models.Profile, error)
	GetProfile(userInput string) (*models.Profile, error)
	UpdateProfile(profile models.Profile) error
}

// DB struct
type DB struct {
	*sql.DB
}

// NewDBInstance create a new db instance, print username, password, host, port, dbname
func NewDBInstance(dbUser string, dbPassword string, dbEndpoint string, dbPort int, dbName string) (*DB, error) {
	fmt.Println("My database - new instance")
	uri := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbEndpoint, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("\nCannot ping db: %s\n", err)
		panic(err)
	}
	fmt.Println("Connection established")
	return &DB{db}, nil
}
