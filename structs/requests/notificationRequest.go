package requests

type NotificationRequest struct {
	Message   string
	UserId    int64
	ServiceId int64
}

type NotificationUpdateRequest struct {
	NotificationId int64
	Status         string
}
