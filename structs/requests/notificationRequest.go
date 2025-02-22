package requests

type NotificationRequest struct {
	UserId   int64
	Service  string
	Status   string
	Category string
	Params   []string
}

type NotificationCategoryRequest struct {
	CategoryName string
	Description  string
}

type NotificationMessageRequest struct {
	ServiceId int64
	StatusId  int64
	Message   string
	Labels    string
}

type NotificationUpdateRequest struct {
	UserId string
	Status string
}
