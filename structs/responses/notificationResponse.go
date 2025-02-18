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

type NotificationCategoryResponse struct {
	StatusCode int
	Category   *models.Notification_category
	StatusDesc string
}

type NotificationCategoriesResponse struct {
	StatusCode int
	Categories *[]interface{}
	StatusDesc string
}

type NotificationMessageResponse struct {
	StatusCode int
	Message    *models.Notification_messages
	StatusDesc string
}

type NotificationMessagesResponse struct {
	StatusCode int
	Messages   *[]interface{}
	StatusDesc string
}
