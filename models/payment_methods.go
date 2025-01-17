package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type PaymentMethods struct {
	Active        int       `orm:"column(active);null"`
	CreatedBy     int       `orm:"column(created_by);null"`
	DateCreated   time.Time `orm:"column(date_created);type(datetime);null;auto_now_add"`
	DateModified  time.Time `orm:"column(date_modified);type(datetime);null"`
	Description   string    `orm:"column(description);size(255);null"`
	ModifiedBy    int       `orm:"column(modified_by);null"`
	PaymentMethod string    `orm:"column(payment_method);size(128)"`
	Id            int       `orm:"column(payment_method_id);auto"`
}

func (t *PaymentMethods) TableName() string {
	return "payment_methods"
}

func init() {
	orm.RegisterModel(new(PaymentMethods))
}

// AddPaymentMethods insert a new PaymentMethods into database and returns
// last inserted Id on success.
func AddPaymentMethods(m *PaymentMethods) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPaymentMethodsById retrieves PaymentMethods by Id. Returns error if
// Id doesn't exist
func GetPaymentMethodsById(id int) (v *PaymentMethods, err error) {
	o := orm.NewOrm()
	v = &PaymentMethods{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPaymentMethods retrieves all PaymentMethods matches certain condition. Returns empty list if
// no records exist
func GetAllPaymentMethods(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PaymentMethods))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []PaymentMethods
	qs = qs.OrderBy(sortFields...)
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

// UpdatePaymentMethods updates PaymentMethods by Id and returns error if
// the record to be updated doesn't exist
func UpdatePaymentMethodsById(m *PaymentMethods) (err error) {
	o := orm.NewOrm()
	v := PaymentMethods{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePaymentMethods deletes PaymentMethods by Id and returns error if
// the record to be deleted doesn't exist
func DeletePaymentMethods(id int) (err error) {
	o := orm.NewOrm()
	v := PaymentMethods{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PaymentMethods{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
