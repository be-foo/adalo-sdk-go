// Package adalo is a thin wrapper for working with the services offered by the Adalo API.
package adalo

import "errors"

var (
	// ErrorUnauthorized is returned by the API when the ApiKey was invalid.
	ErrorUnauthorized = errors.New("unauthorized")

	// ErrorAppMismatch is returned by the API when the AppID does not exist
	// or does not match with your ApiKey.
	ErrorAppMismatch = errors.New("access token / app mismatch")

	// ErrorResourceNotFound is returned when the API response indicates
	// that a record in the collection could not be found.
	// This usually means the passed record ID is invalid.
	ErrorResourceNotFound = errors.New("resource not found")
)

// apiErrorResponse is a representation of the json returned by the Adalo api
// in order to provide explicit error messages.
type apiErrorResponse struct {
	Error string `json:"error"`
}

// ApiKey is the Adalo API key that is globally used to authenticate requests.
var ApiKey string

// AppID is the Adalo app ID that is globally used to perform API requests.
var AppID string
