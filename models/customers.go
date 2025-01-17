package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Customers struct {
	Active             int                 `orm:"column(active);null"`
	CreatedBy          int                 `orm:"column(created_by);null"`
	CustomerCategoryId *CustomerCategories `orm:"column(customer_category_id);rel(fk)"`
	Id                 int                 `orm:"column(customer_id);auto"`
	DateCreated        time.Time           `orm:"column(date_created);type(datetime);null;auto_now_add"`
	DateModified       time.Time           `orm:"column(date_modified);type(datetime);null"`
	ModifiedBy         int                 `orm:"column(modified_by);null"`
	Nickname           string              `orm:"column(nickname);size(100);null"`
	ShopId             *Shops              `orm:"column(shop_id);rel(fk)"`
	UserId             *Users              `orm:"column(user_id);rel(fk)"`
}

func (t *Customers) TableName() string {
	return "customers"
}

func init() {
	orm.RegisterModel(new(Customers))
}

// AddCustomers insert a new Customers into database and returns
// last inserted Id on success.
func AddCustomers(m *Customers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCustomersById retrieves Customers by Id. Returns error if
// Id doesn't exist
func GetCustomersById(id int) (v *Customers, err error) {
	o := orm.NewOrm()
	v = &Customers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCustomers retrieves all Customers matches certain condition. Returns empty list if
// no records exist
func GetAllCustomers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Customers))
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

	var l []Customers
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

// UpdateCustomers updates Customers by Id and returns error if
// the record to be updated doesn't exist
func UpdateCustomersById(m *Customers) (err error) {
	o := orm.NewOrm()
	v := Customers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCustomers deletes Customers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCustomers(id int) (err error) {
	o := orm.NewOrm()
	v := Customers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Customers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
