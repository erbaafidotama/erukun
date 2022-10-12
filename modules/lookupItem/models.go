package lookupItem

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LookupItemModel struct {
	gorm.Model
	LookupItemUuid uuid.UUID `gorm:"primaryKey;type:uuid" json:"lookup_item_uuid"`
	LookupID       int       `json:"lookup_id"`
	LookupItemCode string    `json:"lookup_item_code"`
	LookupItemName string    `json:"lookup_item_name"`
	Keterangan     string    `json:"keterangan"`
	StatusActive   bool      `json:"status_active"`
}

func (LookupItemModel) TableName() string {
	return "master_lookup_items"
}
