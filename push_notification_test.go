package adalo

import (
	"log"
	"testing"
)

func TestSendPushNotification(t *testing.T) {
	resp, err := SendPushNotification(&PushNotificationInput{
		Audience: PushNotificationAudienceInput{
			Email: "john.doe@gmail.com",
		},
		Notification: PushNotificationContentInput{
			Title: "Notification Title",
			Body:  "Click this notification to learn more",
		},
	})
	log.Println(resp, err)
}
