package entity

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Base is base model for every entity in database that define every what every model should have
type Base struct {
	ID string `gorm:"type:uuid;primary_key;" json:"id"`
}

// BeforeCreate function is defining uuid
func (base *Base) BeforeCreate(db *gorm.DB) error {
	uuid, err := uuid.NewV4()

	if err != nil {
		return err
	}

	base.ID = uuid.String()
	return nil
}
