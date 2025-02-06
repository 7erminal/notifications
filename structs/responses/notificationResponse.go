package responses

import "notification_service/models"

type NotificationResponse struct {
	StatusCode   int
	Notification *models.Notifications
	StatusDesc   string
}

type NotificationsResponse struct {
	StatusCode    int
	Notifications *[]interface{}
	StatusDesc    string
}
