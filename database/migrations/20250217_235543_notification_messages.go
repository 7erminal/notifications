package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type NotificationMessages_20250217_235543 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &NotificationMessages_20250217_235543{}
	m.Created = "20250217_235543"

	migration.Register("NotificationMessages_20250217_235543", m)
}

// Run the migrations
func (m *NotificationMessages_20250217_235543) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE notification_messages(`notification_message_id` int(11) NOT NULL AUTO_INCREMENT,`code` varchar(80) NOT NULL,`message` varchar(500) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`notification_message_id`))")
}

// Reverse the migrations
func (m *NotificationMessages_20250217_235543) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `notification_messages`")
}
