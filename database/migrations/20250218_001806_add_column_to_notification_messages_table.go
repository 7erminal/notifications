package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToNotificationMessagesTable_20250218_001806 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToNotificationMessagesTable_20250218_001806{}
	m.Created = "20250218_001806"

	migration.Register("AddColumnToNotificationMessagesTable_20250218_001806", m)
}

// Run the migrations
func (m *AddColumnToNotificationMessagesTable_20250218_001806) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE notification_messages MODIFY COLUMN code int NOT NULL, ADD FOREIGN KEY (code) REFERENCES status(status_id) ON UPDATE CASCADE ON DELETE CASCADE")

}

// Reverse the migrations
func (m *AddColumnToNotificationMessagesTable_20250218_001806) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
