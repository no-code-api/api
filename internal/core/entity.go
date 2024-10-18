package core

import "time"

type Entity struct {
	CreatedAt time.Time `gorm:"notnull"`
	UpdatedAt time.Time `gorm:"notnull"`
}

func (e *Entity) SetCreatedAt() {
	e.CreatedAt = time.Now()
}

func (e *Entity) SetUpdatedAt() {
	e.UpdatedAt = time.Now()
}
