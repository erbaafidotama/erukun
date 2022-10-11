package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Warga model
type BalitaPosyanduMaster struct {
	gorm.Model
	BalitaPosyanduUuid uuid.UUID `gorm:"type:uuid" json:"balita_posyandu_uuid"`
	WargaID            int       `json:"warga_id"`
	StatusActive       bool      `json:"status_active"`
	Keterangan         string    `json:"keterangan"`
}
