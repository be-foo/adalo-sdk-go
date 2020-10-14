package adalo

import (
	"os"
)

// testConfig defines a config argument for the setup function
type testConfig int

// list of accepted values for testConfig
const (
	// unauthorized will set an invalid ApiKey
	unauthorized = 0

	// authorized will set a valid ApiKey (set by default)
	authorized = 1

	// invalidApp will set an invalid AppID
	invalidApp = 2

	// validApp will set a valid AppID (set by default)
	validApp = 3
)

// validApiKey is a copy of the valid api key used in the tests.
// We created this because we want to make tests where ApiKey is invalid and be able to reset it back to the valid api key.
var validApiKey string

// validAppID is a copy of the valid app id used in the tests.
// We created this because we want to make tests where AppKey is invalid and be able to reset it back to the valid app id.
var validAppID string

func init() {
	if os.Getenv("TEST_API_KEY") == "" {
		panic("environment variable TEST_API_KEY is not set")
	}
	validApiKey = os.Getenv("TEST_API_KEY")

	if os.Getenv("TEST_APP_ID") == "" {
		panic("environment variable TEST_APP_ID is not set")
	}
	validAppID = os.Getenv("TEST_APP_ID")
}

// setup is meant to be called at the beginning of each test to setup some conditions.
// By default, so with no args passed, the test is setup to be with valid ApiKey and AppID.
func setup(args ...testConfig) {
	// when no args are passed, setup with valid credentials
	ApiKey = validApiKey
	AppID = validAppID

	for _, arg := range args {
		switch arg {
		case unauthorized:
			ApiKey = "invalid-api-key"
		case authorized:
			ApiKey = validApiKey
		case invalidApp:
			AppID = "invalid-app-id"
		case validApp:
			AppID = validAppID
		}
	}
}
