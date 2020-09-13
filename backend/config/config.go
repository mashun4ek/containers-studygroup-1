package config

import (
	"github.com/mashun4ek/study_group/app/db"
	"github.com/pkg/errors"
)

// ConfigurationVars pass in cli to connect to db
type ConfigurationVars struct {
	IsProd bool
	DBUser string
	DBPass string
	DBName string
	DBPort int
	DBHost string
}

// Env struct for database
type Env struct {
	DB db.DatabaseInterface
}

// NewEnv creates a new environment object managing common state across the the service
func NewEnv(v *ConfigurationVars) (*Env, error) {
	dbConn, err := db.NewDBInstance(v.DBUser, v.DBPass, v.DBHost, v.DBPort, v.DBName)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &Env{dbConn}, nil
}
