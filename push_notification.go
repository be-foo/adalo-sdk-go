package adalo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// pushNotificationApiURL is the base url for push notification api calls on Adalo
const pushNotificationApiURL = "https://api.adalo.com/notifications"

// PushNotificationInput is a representation of the input expected by the Adalo API
type PushNotificationInput struct {
	// (optional) if not specified, global AppID is being taken
	AppID *string `json:"appId"`

	// Audience of this push notification
	Audience PushNotificationAudienceInput `json:"audience"`

	// Notification content
	Notification PushNotificationContentInput `json:"notification"`
}

// PushNotificationAudienceInput is a representation of the audience input in the api request
type PushNotificationAudienceInput struct {
	// Email of the recipient of this push notification
	// It must belong to a user in your Adalo app
	Email string `json:"email"`
}

// PushNotificationAudienceInput is a representation of the notification input in the api request
type PushNotificationContentInput struct {
	// Title of the notification message to be displayed
	Title string `json:"titleText"`

	// Body will be displayed as a description below the Title in the notification message
	Body string `json:"bodyText"`
}

// SendPushNotification requests the Adalo API to send a push notification
func SendPushNotification(input *PushNotificationInput) (interface{}, error) {
	if input.AppID == nil {
		// using global app id
		input.AppID = &AppID
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewReader(inputBytes)

	client := &http.Client{}
	req, err := http.NewRequest("POST", pushNotificationApiURL, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response interface{}
	err = json.Unmarshal(body, &response)
	return response, err
}
