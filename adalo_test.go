package adalo

import (
	"os"
)

type testConfig int

const (
	unauthorized = 0
	authorized   = 1
	invalidApp   = 2
	validApp     = 3
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
