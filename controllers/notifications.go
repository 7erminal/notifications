package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"notification_service/models"
	"notification_service/structs/requests"
	"notification_service/structs/responses"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// NotificationsController operations for Notifications
type NotificationsController struct {
	beego.Controller
}

// URLMapping ...
func (c *NotificationsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("UpdateNotificationReadStatus", c.UpdateNotificationReadStatus)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAllNoficationCategories", c.GetAllNoficationCategories)
	c.Mapping("AddNotificationCategory", c.AddNotificationCategory)
	c.Mapping("GetAllNoficationMessages", c.GetAllNoficationMessages)
	c.Mapping("AddNotificationMessage", c.AddNotificationMessage)
	c.Mapping("GetAllUserNotifications", c.GetAllUserNotifications)
	c.Mapping("GetUserNotificationCount", c.GetUserNotificationCount)
}

// Post ...
// @Title Post
// @Description create Notifications
// @Param	body		body 	requests.NotificationRequest	true		"body for Notifications content"
// @Success 201 {int} responses.NotificationResponse
// @Failure 403 body is empty
// @router / [post]
func (c *NotificationsController) Post() {
	var v requests.NotificationRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	status := "UNREAD"
	message := "An error occurred adding this audit request"
	statusCode := 308
	var user *models.Users
	if user_, err := models.GetUsersById(v.UserId); err == nil {
		user = user_
	} else {
		logs.Info("Error getting user ", err.Error())
		message = "Error getting user"
		statusCode = 608
		resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
		c.Data["json"] = resp
	}
	if service, err := models.GetServicesByName(v.Service); err == nil {
		category, err := models.GetNotification_categoryByName(v.Category)
		if err != nil {
			logs.Info("Category fetched is ", category)
		}
		if statusC, err := models.GetStatusByName(v.Status); err == nil {
			if nMessage, err := models.GetNotification_messagesByCodeAndStatus(*statusC, *service); err == nil {
				tnMessage := nMessage.Message
				configuredLabels := strings.Split(nMessage.Labels, ",")
				// Insert values
				if v.Params != nil {
					for i, mn := range configuredLabels {
						logs.Info("Label:: ", mn)
						logs.Info("Value:: ", v.Params[i])
						tnMessage = strings.Replace(tnMessage, "["+mn+"]", v.Params[i], -1)
					}
				}

				if notificationStatus, err := models.GetNotification_statusByCode(status); err == nil {
					logs.Info("User is ", user)
					notificationResp := models.Notifications{NotificationMessage: tnMessage, NotificationMessageId: nMessage, Category: category, Status: notificationStatus, Service: service, NotificationFor: user}

					if _, err := models.AddNotifications(&notificationResp); err == nil {
						c.Ctx.Output.SetStatus(200)
						statusCode = 200
						message = "Notification inserted successfully"
						resp := responses.NotificationResponse{StatusCode: statusCode, Notification: &notificationResp, StatusDesc: message}
						c.Data["json"] = resp
					} else {
						logs.Info("Error inserting notification ", err.Error())
						message = "Error inserting notification"
						statusCode = 608
						resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
						c.Data["json"] = resp
					}
				} else {
					logs.Info("Error getting notification ", err.Error())
					message = "Error inserting notification. Invalid status."
					statusCode = 608
					resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
					c.Data["json"] = resp
				}
			}
		} else {
			logs.Info("Error getting notification ", err.Error())
			message = "Error inserting notification. Invalid status."
			statusCode = 608
			resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
			c.Data["json"] = resp
		}
	} else {
		logs.Info("Error getting service ", err.Error())
		message = "Error getting service"
		statusCode = 608
		resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Notifications by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Notifications
// @Failure 403 :id is empty
// @router /:id [get]
func (c *NotificationsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	message := "An error occurred adding this audit request"
	statusCode := 308

	v, err := models.GetNotificationsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
		logs.Info("Error getting user ", err.Error())
		message = "Error getting notification"
		statusCode = 608
		resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		c.Ctx.Output.SetStatus(200)
		statusCode = 200
		message = "Notification fetched successfully"
		resp := responses.NotificationResponse{StatusCode: statusCode, Notification: v, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Notifications
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	exclude	query	string	false	"Exclude. e.g. col1:v1,col2:v2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Notifications
// @Failure 403
// @router / [get]
func (c *NotificationsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var exclude = make(map[string]string)
	var limit int64 = 100
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	// query: k:v,k:v
	if v := c.GetString("exclude"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			exclude[k] = v
		}
	}
	message := "An error occurred adding this audit request"
	statusCode := 308

	l, err := models.GetAllNotifications(query, exclude, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("There was an error fetching notifications ", err.Error())
		message = "Error fetching notifications"
		statusCode = 608
		resp := responses.NotificationsResponse{StatusCode: statusCode, Notifications: nil, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Info("Notifications fetched successfully")
		fmt.Printf("Value of v: %+v\n", l)
		c.Ctx.Output.SetStatus(200)
		statusCode = 200
		message = "Notifications fetched successfully"
		resp := responses.NotificationsResponse{StatusCode: statusCode, Notifications: &l, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAllUserNotifications ...
// @Title Get All User Not
// @Description get Notifications
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	exclude	query	string	false	"Exclude. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.NotificationsResponse
// @Failure 403
// @router /get-user-notifications/:id [get]
func (c *NotificationsController) GetAllUserNotifications() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 50
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	message := "An error occurred adding this audit request"
	statusCode := 308

	if user, err := models.GetUsersById(id); err == nil {
		l, err := models.GetAllUserNotifications(*user, query, fields, sortby, order, offset, limit)
		if err != nil {
			logs.Error("There was an error fetching notifications ", err.Error())
			message = "Error fetching notifications"
			statusCode = 608
			resp := responses.NotificationsResponse{StatusCode: statusCode, Notifications: nil, StatusDesc: message}
			c.Data["json"] = resp
		} else {
			c.Data["json"] = l
			c.Ctx.Output.SetStatus(200)
			statusCode = 200
			message = "Notifications fetched successfully"
			resp := responses.NotificationsResponse{StatusCode: statusCode, Notifications: &l, StatusDesc: message}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("There was an error fetching notifications ", err.Error())
		message = "Error fetching notifications for this user"
		statusCode = 608
		resp := responses.NotificationsResponse{StatusCode: statusCode, Notifications: nil, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// AddNotificationCategory ...
// @Title Post Notification Category
// @Description create Notifications
// @Param	body		body 	requests.NotificationCategoryRequest	true		"body for Notifications content"
// @Success 201 {int} responses.NotificationCategoryResponse
// @Failure 403 body is empty
// @router /add-notification-category [post]
func (c *NotificationsController) AddNotificationCategory() {
	var v requests.NotificationCategoryRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	message := "An error occurred adding this category"
	statusCode := 308

	notificationCategoryResp := models.Notification_category{Category: v.CategoryName, Description: v.Description, Active: 1, CreatedBy: 1, ModifiedBy: 1, DateCreated: time.Now(), DateModified: time.Now()}

	if _, err := models.AddNotification_category(&notificationCategoryResp); err == nil {
		c.Ctx.Output.SetStatus(200)
		statusCode = 200
		message = "Notification category inserted successfully"
		resp := responses.NotificationCategoryResponse{StatusCode: statusCode, Category: &notificationCategoryResp, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Info("Error inserting notification ", err.Error())
		message = "Error inserting notification"
		statusCode = 608
		resp := responses.NotificationCategoryResponse{StatusCode: statusCode, Category: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetAllNoficationCategories ...
// @Title Get All Notification Categories
// @Description get Notifications
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Notifications
// @Failure 403
// @router /get-all-notification-categories [get]
func (c *NotificationsController) GetAllNoficationCategories() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	message := "An error occurred adding this audit request"
	statusCode := 308

	l, err := models.GetAllNotification_category(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("There was an error fetching notifications ", err.Error())
		message = "Error fetching notifications"
		statusCode = 608
		resp := responses.NotificationCategoriesResponse{StatusCode: statusCode, Categories: nil, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		c.Data["json"] = l
		c.Ctx.Output.SetStatus(200)
		statusCode = 200
		message = "Notifications fetched successfully"
		resp := responses.NotificationCategoriesResponse{StatusCode: statusCode, Categories: &l, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// AddNotificationMessage ...
// @Title Post Notification Message
// @Description create Notification Message
// @Param	body		body 	requests.NotificationMessageRequest	true		"body for Notifications content"
// @Success 201 {int} responses.NotificationCategoryResponse
// @Failure 403 body is empty
// @router /add-notification-message [post]
func (c *NotificationsController) AddNotificationMessage() {
	var v requests.NotificationMessageRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	message := "An error occurred adding this category"
	statusCode := 308

	if service, err := models.GetServicesById(v.ServiceId); err == nil {
		if status, err := models.GetStatusById(v.StatusId); err == nil {
			notificationMessageResp := models.Notification_messages{Code: status, Service: service, Message: v.Message, Labels: v.Labels, Active: 1, CreatedBy: 1, ModifiedBy: 1, DateCreated: time.Now(), DateModified: time.Now()}
			if _, err := models.AddNotification_messages(&notificationMessageResp); err == nil {
				c.Ctx.Output.SetStatus(200)
				statusCode = 200
				message = "Notification message inserted successfully"
				resp := responses.NotificationMessageResponse{StatusCode: statusCode, Message: &notificationMessageResp, StatusDesc: message}
				c.Data["json"] = resp
			} else {
				logs.Info("Error inserting notification ", err.Error())
				message = "Error inserting notification message"
				statusCode = 608
				resp := responses.NotificationMessageResponse{StatusCode: statusCode, Message: nil, StatusDesc: message}
				c.Data["json"] = resp
			}
		} else {
			logs.Info("Error inserting notification ", err.Error())
			message = "Error getting status"
			statusCode = 608
			resp := responses.NotificationMessageResponse{StatusCode: statusCode, Message: nil, StatusDesc: message}
			c.Data["json"] = resp
		}
	} else {
		logs.Info("Error inserting notification ", err.Error())
		message = "Error getting service"
		statusCode = 603
		resp := responses.NotificationMessageResponse{StatusCode: statusCode, Message: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetAllNoficationMessages ...
// @Title Get All Notification Messages
// @Description get Notifications
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Notifications
// @Failure 403
// @router /get-all-notification-messages [get]
func (c *NotificationsController) GetAllNoficationMessages() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	message := "An error occurred adding this audit request"
	statusCode := 308

	l, err := models.GetAllNotification_messages(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("There was an error fetching notifications ", err.Error())
		message = "Error fetching notification messages"
		statusCode = 608
		resp := responses.NotificationMessagesResponse{StatusCode: statusCode, Messages: nil, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		c.Data["json"] = l
		c.Ctx.Output.SetStatus(200)
		statusCode = 200
		message = "Notification messages fetched successfully"
		resp := responses.NotificationMessagesResponse{StatusCode: statusCode, Messages: &l, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// UpdateNotificationReadStatus ...
// @Title Update notification read status
// @Description update the Notifications
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.NotificationUpdateRequest	true		"body for Notifications content"
// @Success 200 {object} models.Notifications
// @Failure 403 :id is not int
// @router /update-read-status/:id [put]
func (c *NotificationsController) UpdateNotificationReadStatus() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.NotificationUpdateRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	proceed := false
	message := "An error occurred adding this audit request"
	statusCode := 308
	status := models.Notification_status{}
	if status_, err := models.GetNotification_statusByCode(v.Status); err == nil {
		status = *status_
		proceed = true
	}
	if proceed {
		if notification, err := models.GetNotificationsById(id); err == nil {
			uid, _ := strconv.ParseInt(v.UserId, 10, 64)
			if user, err := models.GetUsersById(uid); err == nil {
				if user == notification.NotificationFor {
					notification.Status = &status
					notification.ReadDate = time.Now()
					if err := models.UpdateNotificationsById(notification); err == nil {
						c.Ctx.Output.SetStatus(200)
						statusCode = 200
						message = "Notification updated successfully"
						resp := responses.NotificationResponse{StatusCode: statusCode, Notification: notification, StatusDesc: message}
						c.Data["json"] = resp
					} else {
						logs.Info("Error updating notification ", err.Error())
						message = "Error updating notification"
						statusCode = 608
						resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
						c.Data["json"] = resp
					}
				} else {
					logs.Info("Error updating notification ", err.Error())
					message = "Error updating notification. User not found."
					statusCode = 608
					resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
					c.Data["json"] = resp
				}
			} else {
				logs.Info("Error updating notification ", err.Error())
				message = "Error updating notification. User not found."
				statusCode = 608
				resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
				c.Data["json"] = resp
			}
		} else {
			logs.Info("Error getting notification ", err.Error())
			message = "Error getting notification"
			statusCode = 608
			resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
			c.Data["json"] = resp
		}
	} else {
		logs.Info("Error getting status ")
		message = "Error getting status"
		statusCode = 608
		resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Notifications
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *NotificationsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteNotifications(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetUserNotificationCount ...
// @Title Get User notification count
// @Description get notification count
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 :id is empty
// @router /count/:id [get]
func (c *NotificationsController) GetUserNotificationCount() {
	// q, err := models.GetItemsById(id)
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	var query = make(map[string]string)

	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	if user, err := models.GetUsersById(id); err == nil {
		v, err := models.GetNotificationCount(query, *user)
		count := strconv.FormatInt(v, 10)

		if err != nil {
			logs.Error("Error fetching count of customers ... ", err.Error())
			resp := responses.StringResponseDTO{StatusCode: 301, Value: "", StatusDesc: err.Error()}
			c.Data["json"] = resp
		} else {
			resp := responses.StringResponseDTO{StatusCode: 200, Value: count, StatusDesc: "Count fetched successfully"}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Error fetching count of notifications ... ", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 301, Value: "", StatusDesc: err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
