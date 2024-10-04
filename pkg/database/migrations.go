package database

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) {
	CreateUserTable(db)
}

func CreateUserTable(db *gorm.DB) {
	type User struct {
		gorm.Model
		Id       uint `gorm:"unique;primaryKey;autoIncrement"`
		Name     string
		Email    string `gorm:"unique"`
		Password string
	}

	if !Exists(db, &User{}) {
		db.Migrator().CreateTable(&User{})
	}
}
