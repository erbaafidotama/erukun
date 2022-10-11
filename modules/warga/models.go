package warga

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Warga model
type WargaMasterModel struct {
	gorm.Model
	WargaUuid      uuid.UUID `gorm:"type:uuid" json:"warga_uuid"`
	NamaLengkap    string    `json:"nama_lengkap"`
	NoKk           string    `json:"no_kk"`
	Nik            string    `json:"nik"`
	TempatLahir    string    `json:"tempat_lahir"`
	TanggalLahir   string    `json:"tanggal_lahir"`
	JenisKelaminId int       `json:"jenis_kelamin_id"`
	Alamat         string    `json:"alamat"`
	Rt             string    `json:"rt"`
	Rw             string    `json:"rw"`
	NoRumah        string    `json:"no_rumah"`
}
