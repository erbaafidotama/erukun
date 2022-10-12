package lookup

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LookupModel struct {
	gorm.Model
	LookupUuid   uuid.UUID `gorm:"primaryKey;type:uuid" json:"lookup_uuid"`
	LookupCode   string    `json:"lookup_code"`
	Keterangan   string    `json:"keterangan"`
	StatusActive bool      `json:"status_active"`
}

func (LookupModel) TableName() string {
	return "master_lookups"
}
