package server

import "net/http"

// Route struct
type Route struct {
	Description string
	Path        string
	Method      string
	Handler     http.HandlerFunc
}

// Routes slice
type Routes []Route

func buildRoutes(e *Env) Routes {
	return Routes{
		{
			Description: "Create profile",
			Path:        "/profiles/{profileId}",
			Method:      "POST",
			Handler:     e.createProfile,
		},
		{
			Description: "Get profile",
			Path:        "/profiles/{profileId}",
			Method:      "GET",
			Handler:     e.getProfile,
		},
		{
			Description: "Get all profiles",
			Path:        "/profiles",
			Method:      "GET",
			Handler:     e.getAllProfiles,
		},
		{
			Description: "Update profile",
			Path:        "/profiles/{id}",
			Method:      "UPDATE",
			Handler:     e.editProfile,
		},
		{
			Description: "Delete profile",
			Path:        "/profiles/{profileId}",
			Method:      "DELETE",
			Handler:     e.deleteProfile,
		},
	}
}
