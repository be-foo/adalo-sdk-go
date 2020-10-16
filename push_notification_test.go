package adalo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendPushNotification(t *testing.T) {
	t.Run("send to valid user", func(t *testing.T) {
		setup()

		i, err := SendPushNotification(&PushNotificationInput{
			Audience: PushNotificationAudienceInput{
				Email: "john.doe@gmail.com", // a user with this email exists in the Adalo app
			},
			Notification: PushNotificationContentInput{
				Title: "Notification Title",
				Body:  "Click this notification to learn more",
			},
		})

		assert.Nil(t, err)
		assert.Equal(t, 1, i)
	})

	t.Run("send to user that does not exist", func(t *testing.T) {
		setup()

		i, err := SendPushNotification(&PushNotificationInput{
			Audience: PushNotificationAudienceInput{
				Email: "invalid@gmail.com",
			},
			Notification: PushNotificationContentInput{
				Title: "Notification Title",
				Body:  "Click this notification to learn more",
			},
		})

		assert.Error(t, err)
		assert.Equal(t, ErrorUserNotFound, err)
		assert.Equal(t, 0, i)
	})

	t.Run("unauthorized", func(t *testing.T) {
		setup(unauthorized)

		_, err := SendPushNotification(&PushNotificationInput{
			Audience: PushNotificationAudienceInput{
				Email: "john.doe@gmail.com",
			},
			Notification: PushNotificationContentInput{
				Title: "Notification Title",
				Body:  "Click this notification to learn more",
			},
		})

		assert.Equal(t, ErrorUnauthorized, err)
	})

	t.Run("invalid app", func(t *testing.T) {
		setup(invalidApp)

		_, err := SendPushNotification(&PushNotificationInput{
			Audience: PushNotificationAudienceInput{
				Email: "john.doe@gmail.com",
			},
			Notification: PushNotificationContentInput{
				Title: "Notification Title",
				Body:  "Click this notification to learn more",
			},
		})

		assert.Equal(t, ErrorAppMismatch, err)
	})
}
