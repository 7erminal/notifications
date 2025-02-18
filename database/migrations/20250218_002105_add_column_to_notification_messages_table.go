package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToNotificationMessagesTable_20250218_002105 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToNotificationMessagesTable_20250218_002105{}
	m.Created = "20250218_002105"

	migration.Register("AddColumnToNotificationMessagesTable_20250218_002105", m)
}

// Run the migrations
func (m *AddColumnToNotificationMessagesTable_20250218_002105) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE notification_messages ADD COLUMN service_id int NOT NULL AFTER code, ADD FOREIGN KEY (service_id) REFERENCES service(service_id) ON UPDATE CASCADE ON DELETE CASCADE")

}

// Reverse the migrations
func (m *AddColumnToNotificationMessagesTable_20250218_002105) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
