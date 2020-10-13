package adalo

import "os"

func init() {
	if os.Getenv("TEST_API_KEY") == "" {
		panic("environment variable TEST_API_KEY is not set")
	}
	ApiKey = os.Getenv("TEST_API_KEY")

	if os.Getenv("TEST_APP_ID") == "" {
		panic("environment variable TEST_APP_ID is not set")
	}
	AppID = os.Getenv("TEST_APP_ID")
}
