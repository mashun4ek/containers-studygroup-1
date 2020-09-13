package db

import (
	"fmt"

	"github.com/mashun4ek/study_group/app/models"
)

// AddDatabase is doesn't exist
// func (d *DB) AddDatabase(dbName string) error {
// 	// check in metadata about databases if db already exists
// 	sqlStatement := fmt.Sprint(`SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1);`)
// 	row := d.QueryRow(sqlStatement, dbName)
// 	var exists bool
// 	err := row.Scan(&exists)
// 	if exists == false {
// 		sqlStatement = `CREATE DATABASE ${dbName};`
// 		_, err = d.Exec(sqlStatement)
// 		if err != nil {
// 			log.Fatal("Can't create DB")
// 			return err
// 		}
// 	}
// 	return nil
// }

// CreateNewProfile inserts new person to db
func (d *DB) CreateNewProfile(userInput *models.Profile) error {
	// TODO: use returned id
	defer d.Close()
	sqlStatement := fmt.Sprint(`INSERT INTO contacts (dob, email, first_name, last_name, phone)
	VALUES ($1, $2, $3, $4)
	RETURNING id;`)
	_, err := d.Exec(sqlStatement, userInput.DOB, userInput.Email, userInput.FirstName, userInput.LastName, userInput.Phone)
	return err
}

// GetProfile returns person's info by Name
func (d *DB) GetProfile(firstName string) (*models.Profile, error) {
	sqlStatement := fmt.Sprint(`SELECT FROM contacts WHERE first_name = $1`)
	var person models.Profile
	row, err := d.Query(sqlStatement, firstName)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	if err = row.Scan(&person.ID, &person.DOB, &person.Email, &person.FirstName, &person.LastName, &person.Phone); err != nil {
		return nil, err
	}
	return &person, nil
}

// GetAllProfiles returns contact information for all profiles
func (d *DB) GetAllProfiles() ([]models.Profile, error) {
	// TODO: use limit
	sqlStatement := fmt.Sprint(`SELECT * FROM contacts`)
	var allProfiles []models.Profile
	var person models.Profile
	rows, err := d.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&person.ID, &person.DOB, &person.FirstName, &person.LastName, &person.Email, &person.Phone); err != nil {
			return nil, err
		}
		allProfiles = append(allProfiles, person)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return allProfiles, nil
}

// UpdateProfile updates profile info in db
func (d *DB) UpdateProfile(profile models.Profile) error {
	defer d.Close()
	sqlStatement := fmt.Sprint(`UPDATE contacts SET email=$2, first_name=$3, last_name=$4, phone=$5 WHERE id = $1`)
	_, err := d.Exec(sqlStatement, profile.ID, profile.Email, profile.FirstName, profile.LastName, profile.Phone)
	return err
}

// DeleteProfile delete profile from db
func (d *DB) DeleteProfile(profileID string) error {
	defer d.Close()
	// DELETE FROM contacts WHERE id=1;
	sqlStatement := fmt.Sprint(`DELETE FROM contacts WHERE id = $1`)
	_, err := d.Exec(sqlStatement, profileID)
	return err
}
