package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type NotificationStatus_20250206_111844 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &NotificationStatus_20250206_111844{}
	m.Created = "20250206_111844"

	migration.Register("NotificationStatus_20250206_111844", m)
}

// Run the migrations
func (m *NotificationStatus_20250206_111844) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE notification_status(`notification_status_id` int(11) NOT NULL AUTO_INCREMENT,`status` varchar(80) NOT NULL,`status_code` varchar(80) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`notification_status_id`))")
}

// Reverse the migrations
func (m *NotificationStatus_20250206_111844) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `notification_status`")
}
