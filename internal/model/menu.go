package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	ParentID  *uuid.UUID     `gorm:"type:uuid;index" json:"parent_id"`
	Order     int            `gorm:"type:int;default:0;index" json:"order"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Children []Menu `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

func (m *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
