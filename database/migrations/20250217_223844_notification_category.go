package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type NotificationCategory_20250217_223844 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &NotificationCategory_20250217_223844{}
	m.Created = "20250217_223844"

	migration.Register("NotificationCategory_20250217_223844", m)
}

// Run the migrations
func (m *NotificationCategory_20250217_223844) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE notification_category(`notification_id` int(11) NOT NULL AUTO_INCREMENT,`category` varchar(80) NOT NULL,`description` varchar(80) NOT NULL,`date_created` datetime NOT NULL,`date_modified` datetime NOT NULL,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT NULL,PRIMARY KEY (`notification_id`))")
}

// Reverse the migrations
func (m *NotificationCategory_20250217_223844) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `notification_category`")
}
