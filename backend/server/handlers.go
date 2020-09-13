package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mashun4ek/study_group/app/models"
)

func (e *Env) createProfile(w http.ResponseWriter, r *http.Request) {
	// parse and decode the request body into a new Profile instance
	userInput := &models.Profile{}
	err := json.NewDecoder(r.Body).Decode(userInput)
	if err != nil {
		log.Printf("Error decoding user's input createProfile handler: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// insert a new profile in db
	if err = e.DB.CreateNewProfile(userInput); err != nil {
		log.Printf("Error creating/inserting a new profile: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

// get profile info
func (e *Env) getProfile(w http.ResponseWriter, r *http.Request) {
	// get var from path
	vars := mux.Vars(r)
	// get a user profile info from db (type Profile struct)
	userProfile, err := e.DB.GetProfile(vars["firstName"])
	if err != nil {
		log.Printf("Can't get user profile e.DB.GetUserProfile: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&userProfile)
	if err != nil {
		log.Printf("Error marshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// respond to client, set header, write body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

func (e *Env) deleteProfile(w http.ResponseWriter, r *http.Request) {
	var userIDDelete string
	// parse and decode the request body get profileId that should be deleted
	err := json.NewDecoder(r.Body).Decode(userIDDelete)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = e.DB.DeleteProfile(userIDDelete)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}

func (e *Env) getAllProfiles(w http.ResponseWriter, r *http.Request) {
	allProfiles, err := e.DB.GetAllProfiles()
	if err != nil {
		log.Printf("Can't get profiles e.DB.GetAllProfiles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(allProfiles)
	if err != nil {
		log.Printf("Error marshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (e *Env) editProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		log.Printf("Error unmarshalling in editProfile func: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = e.DB.UpdateProfile(profile)
	if err != nil {
		log.Printf("Error updating profile: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
