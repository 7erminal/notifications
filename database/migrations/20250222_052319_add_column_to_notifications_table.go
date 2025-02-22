package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToNotificationsTable_20250222_052319 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToNotificationsTable_20250222_052319{}
	m.Created = "20250222_052319"

	migration.Register("AddColumnToNotificationsTable_20250222_052319", m)
}

// Run the migrations
func (m *AddColumnToNotificationsTable_20250222_052319) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE notifications ADD COLUMN role_id int DEFAULT NULL, ADD FOREIGN KEY (role_id) REFERENCES roles(role_id) ON UPDATE CASCADE ON DELETE NO ACTION")
}

// Reverse the migrations
func (m *AddColumnToNotificationsTable_20250222_052319) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
