package database

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) {
	CreateUserTable(db)
}

func CreateUserTable(db *gorm.DB) {
	type User struct {
		gorm.Model
		Id       int    `gorm:"unique;primaryKey;autoIncrement"`
		Name     string `gorm:"size:150;notnull"`
		Email    string `gorm:"size:100;unique;notnull"`
		Password string `gorm:"size:60;notnull"`
	}

	if !Exists(db, &User{}) {
		db.Migrator().CreateTable(&User{})
	}
}
