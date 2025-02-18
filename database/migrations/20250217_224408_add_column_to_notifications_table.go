package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToNotificationsTable_20250217_224408 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToNotificationsTable_20250217_224408{}
	m.Created = "20250217_224408"

	migration.Register("AddColumnToNotificationsTable_20250217_224408", m)
}

// Run the migrations
func (m *AddColumnToNotificationsTable_20250217_224408) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE notification_category RENAME COLUMN notification_id TO notification_category_id")
}

// Reverse the migrations
func (m *AddColumnToNotificationsTable_20250217_224408) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
