package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Notifications_20250206_112004 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Notifications_20250206_112004{}
	m.Created = "20250206_112004"

	migration.Register("Notifications_20250206_112004", m)
}

// Run the migrations
func (m *Notifications_20250206_112004) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE notifications(`notification_id` int(11) NOT NULL AUTO_INCREMENT,`notification_message` varchar(500) NOT NULL,`status` int(11) NOT NULL,`service_id` int(11) NOT NULL,`notification_for` int(11) NOT NULL,`read_date` datetime NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`notification_id`), FOREIGN KEY (service_id) REFERENCES actions(action_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (status) REFERENCES notification_status(notification_status_id) ON UPDATE CASCADE ON DELETE CASCADE, FOREIGN KEY (notification_for) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *Notifications_20250206_112004) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `notifications`")
}
