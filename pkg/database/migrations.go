package database

import (
	"time"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	CreateUserTable(db)
	CreateProjectsTable(db)
	CreateResourcesTable(db)
}

type User struct {
	Id        uint      `gorm:"unique;primaryKey;autoIncrement"`
	Name      string    `gorm:"size:150;notnull"`
	Email     string    `gorm:"size:100;unique;notnull"`
	Password  string    `gorm:"size:60;notnull"`
	CreatedAt time.Time `gorm:"notnull"`
	UpdatedAt time.Time `gorm:"notnull"`
}

type Project struct {
	Id        string    `gorm:"size:32;unique;primaryKey"`
	UserId    uint      `gorm:"notnull"`
	User      User      `gorm:"foreignKey:UserId;references:id"`
	Name      string    `gorm:"size:30;notnull"`
	CreatedAt time.Time `gorm:"notnull"`
	UpdatedAt time.Time `gorm:"notnull"`
}

type Resource struct {
	Id        string      `gorm:"size:32;notnull;primaryKey"`
	ProjectId string      `gorm:"size:32;notnull"`
	Path      string      `gorm:"size:50;notnull"`
	Project   Project     `gorm:"foreignKey:ProjectId;references:id"`
	Endpoints []*Endpoint `gorm:"foreignKey:ResourceId"`
	CreatedAt time.Time   `gorm:"notnull"`
	UpdatedAt time.Time   `gorm:"notnull"`
}

type Endpoint struct {
	Id         uint      `gorm:"primaryKey;autoIncrement"`
	Path       string    `gorm:"size:50;notnull"`
	Method     string    `gorm:"size:10;notnull"`
	ResourceId string    `gorm:"size:32;notnull"`
	Resource   Resource  `gorm:"foreignKey:ResourceId;references:id"`
	CreatedAt  time.Time `gorm:"notnull"`
	UpdatedAt  time.Time `gorm:"notnull"`
}

func CreateUserTable(db *gorm.DB) {
	if !Exists(db, &User{}) {
		db.Migrator().CreateTable(&User{})
	}
}

func CreateProjectsTable(db *gorm.DB) {
	if !Exists(db, &Project{}) {
		db.Migrator().CreateTable(&Project{})
	}
}

func CreateResourcesTable(db *gorm.DB) {
	if !Exists(db, &Resource{}) {
		db.Migrator().CreateTable(&Resource{})
	}
	CreateEndpointsTable(db)
}

func CreateEndpointsTable(db *gorm.DB) {
	if !Exists(db, &Endpoint{}) {
		db.Migrator().CreateTable(&Endpoint{})
	}
}
