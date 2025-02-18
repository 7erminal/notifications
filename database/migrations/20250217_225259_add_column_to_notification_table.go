package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToNotificationTable_20250217_225259 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToNotificationTable_20250217_225259{}
	m.Created = "20250217_225259"

	migration.Register("AddColumnToNotificationTable_20250217_225259", m)
}

// Run the migrations
func (m *AddColumnToNotificationTable_20250217_225259) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE notifications ADD COLUMN category_id int DEFAULT NULL, ADD FOREIGN KEY (category_id) REFERENCES notification_category(notification_category_id) ON UPDATE CASCADE ON DELETE CASCADE")
}

// Reverse the migrations
func (m *AddColumnToNotificationTable_20250217_225259) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
