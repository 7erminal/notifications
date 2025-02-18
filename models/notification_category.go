package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Notification_category struct {
	NotificationCategoryId int64     `orm:"auto"`
	Category               string    `orm:"size(80)"`
	Description            string    `orm:"size(80)"`
	DateCreated            time.Time `orm:"type(datetime)"`
	DateModified           time.Time `orm:"type(datetime)"`
	CreatedBy              int
	ModifiedBy             int
	Active                 int
}

func init() {
	orm.RegisterModel(new(Notification_category))
}

// AddNotification_category insert a new Notification_category into database and returns
// last inserted Id on success.
func AddNotification_category(m *Notification_category) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNotification_categoryById retrieves Notification_category by Id. Returns error if
// Id doesn't exist
func GetNotification_categoryById(id int64) (v *Notification_category, err error) {
	o := orm.NewOrm()
	v = &Notification_category{NotificationCategoryId: id}
	if err = o.QueryTable(new(Notification_category)).Filter("NotificationCategoryId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNotification_category retrieves all Notification_category matches certain condition. Returns empty list if
// no records exist
func GetAllNotification_category(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Notification_category))
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

	var l []Notification_category
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

// UpdateNotification_category updates Notification_category by Id and returns error if
// the record to be updated doesn't exist
func UpdateNotification_categoryById(m *Notification_category) (err error) {
	o := orm.NewOrm()
	v := Notification_category{NotificationCategoryId: m.NotificationCategoryId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNotification_category deletes Notification_category by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNotification_category(id int64) (err error) {
	o := orm.NewOrm()
	v := Notification_category{NotificationCategoryId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Notification_category{NotificationCategoryId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
