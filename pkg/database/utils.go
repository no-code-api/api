package database

import "gorm.io/gorm"

func Exists(db *gorm.DB, schema ...interface{}) bool {
	return db.Migrator().HasTable(schema)
}
