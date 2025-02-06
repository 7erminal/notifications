package controllers

import (
	"encoding/json"
	"errors"
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
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
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
	if user, err := models.GetUsersById(v.UserId); err == nil {
		if service, err := models.GetServicesById(v.ServiceId); err == nil {
			if notification, err := models.GetNotification_statusByCode(status); err == nil {
				notificationResp := models.Notifications{NotificationMessage: v.Message, Status: notification, ServiceId: service, NotificationFor: user}

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
		} else {
			logs.Info("Error getting service ", err.Error())
			message = "Error getting service"
			statusCode = 608
			resp := responses.NotificationResponse{StatusCode: statusCode, Notification: nil, StatusDesc: message}
			c.Data["json"] = resp
		}
	} else {
		logs.Info("Error getting user ", err.Error())
		message = "Error getting user"
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
	v, err := models.GetNotificationsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Notifications
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
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

	l, err := models.GetAllNotifications(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
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
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Notifications
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.NotificationUpdateRequest	true		"body for Notifications content"
// @Success 200 {object} models.Notifications
// @Failure 403 :id is not int
// @router /:id [put]
func (c *NotificationsController) Put() {
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
