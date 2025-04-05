package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type Notifications struct {
	NotificationId        int64                  `orm:"auto"`
	NotificationMessage   string                 `orm:"size(500)"`
	Status                *Notification_status   `orm:"rel(fk);column(status);"`
	Category              *Notification_category `orm:"rel(fk);column(category_id)"`
	Service               *Services              `orm:"rel(fk);column(service_id)"`
	NotificationFor       *Users                 `orm:"rel(fk);column(notification_for);null"`
	Role                  *Roles                 `orm:"rel(fk);column(role_id);null"`
	NotificationMessageId *Notification_messages `orm:"rel(fk);column(notification_message_id)"`
	ReadDate              time.Time              `orm:"type(datetime)"`
	DateCreated           time.Time              `orm:"type(datetime)"`
	DateModified          time.Time              `orm:"type(datetime)"`
	CreatedBy             int
	ModifiedBy            int
}

func init() {
	orm.RegisterModel(new(Notifications))
}

// AddNotifications insert a new Notifications into database and returns
// last inserted Id on success.
func AddNotifications(m *Notifications) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNotificationsById retrieves Notifications by Id. Returns error if
// Id doesn't exist
func GetNotificationsById(id int64) (v *Notifications, err error) {
	o := orm.NewOrm()
	v = &Notifications{NotificationId: id}
	if err = o.QueryTable(new(Notifications)).Filter("NotificationId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetOrderCount retrieves Items by Id. Returns error if
// Id doesn't exist
func GetNotificationCount(query map[string]string, search map[string]string, user Users) (c int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Notifications))

	if len(search) > 0 {
		cond := orm.NewCondition()
		for k, v := range search {
			// rewrite dot-notation to Object__Attribute
			k = strings.Replace(k, ".", "__", -1)
			if strings.Contains(k, "isnull") {
				qs = qs.Filter(k, (v == "true" || v == "1"))
			} else {
				logs.Info("Adding or statement")
				cond = cond.Or(k+"__icontains", v)

				// qs = qs.Filter(k+"__icontains", v)

			}
		}
		logs.Info("Condition set ", qs)
		qs = qs.SetCond(cond)
	}

	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}

	qs.Filter("NotificationFor", user)

	if c, err = qs.RelatedSel().Count(); err == nil {
		return c, nil
	}
	return 0, err
}

// GetAllNotifications retrieves all Notifications matches certain condition. Returns empty list if
// no records exist
func GetAllNotifications(query map[string]string, exclude map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("notification_for__isnull", true)
	o := orm.NewOrm()
	qs := o.QueryTable(new(Notifications))

	qs.SetCond(cond1)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Exclude(k, v)
		logs.Info("Condition set")
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Notifications
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// GetAllNotifications retrieves all Notifications matches certain condition. Returns empty list if
// no records exist
func GetAllUserNotifications(user Users, query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Notifications))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Notifications
	qs = qs.Filter("NotificationFor", user)
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateNotifications updates Notifications by Id and returns error if
// the record to be updated doesn't exist
func UpdateNotificationsById(m *Notifications) (err error) {
	o := orm.NewOrm()
	v := Notifications{NotificationId: m.NotificationId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNotifications deletes Notifications by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNotifications(id int64) (err error) {
	o := orm.NewOrm()
	v := Notifications{NotificationId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Notifications{NotificationId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
