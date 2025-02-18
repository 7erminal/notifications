package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToNotificationsTable_20250218_000054 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToNotificationsTable_20250218_000054{}
	m.Created = "20250218_000054"

	migration.Register("AddColumnToNotificationsTable_20250218_000054", m)
}

// Run the migrations
func (m *AddColumnToNotificationsTable_20250218_000054) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE notifications ADD COLUMN notification_message_id int DEFAULT NULL AFTER notification_id, ADD FOREIGN KEY (notification_message_id) REFERENCES notification_messages(notification_message_id) ON UPDATE CASCADE ON DELETE CASCADE")

}

// Reverse the migrations
func (m *AddColumnToNotificationsTable_20250218_000054) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
