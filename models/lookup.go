package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lookup struct {
	gorm.Model
	LookupUuid   uuid.UUID `gorm:"type:uuid" json:"lookup_uuid"`
	LookupCode   string    `json:"lookup_code"`
	Keterangan   string    `json:"keterangan"`
	StatusActive bool      `json:"status_active"`
}
