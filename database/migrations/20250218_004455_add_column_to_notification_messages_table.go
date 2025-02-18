package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToNotificationMessagesTable_20250218_004455 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToNotificationMessagesTable_20250218_004455{}
	m.Created = "20250218_004455"

	migration.Register("AddColumnToNotificationMessagesTable_20250218_004455", m)
}

// Run the migrations
func (m *AddColumnToNotificationMessagesTable_20250218_004455) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE notification_messages ADD COLUMN labels varchar(255) DEFAULT NULL AFTER message")

}

// Reverse the migrations
func (m *AddColumnToNotificationMessagesTable_20250218_004455) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
