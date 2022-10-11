package warga

import (
	"erukunrukun/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type wargaMasterRequest struct {
	NamaLengkap    string `json:"nama_lengkap"`
	NoKk           string `json:"no_kk"`
	Nik            string `json:"nik"`
	TempatLahir    string `json:"tempat_lahir"`
	TanggalLahir   string `json:"tanggal_lahir"`
	JenisKelaminId int    `json:"jenis_kelamin_id"`
	Alamat         string `json:"alamat"`
	Rt             string `json:"rt"`
	Rw             string `json:"rw"`
	NoRumah        string `json:"no_rumah"`
}

func GetListWarga(c *gin.Context) {
	db := config.InitDB()
	wargas := []WargaMasterModel{}

	// select * from User
	if err := db.Find(&wargas).Error; err != nil {
		// return error
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	// return complete
	c.JSON(200, gin.H{
		"message": "GET list data user success",
		"data":    wargas,
	})
}

func GetOneAnakByUuid(c *gin.Context) {
	db := config.InitDB()
	var warga WargaMasterModel

	uuid := c.Param("warga_uuid")
	if err := db.Where("warga_uuid = ?", uuid).First(&warga).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal get",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil get",
			"data":   warga,
		})
	}
}

func PostWarga(c *gin.Context) {
	db := config.InitDB()
	var wargaMasterReq wargaMasterRequest

	if err := c.BindJSON(&wargaMasterReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	warga := WargaMasterModel{
		WargaUuid:      uuid.New(),
		NamaLengkap:    wargaMasterReq.NamaLengkap,
		NoKk:           wargaMasterReq.NoKk,
		Nik:            wargaMasterReq.Nik,
		TempatLahir:    wargaMasterReq.TempatLahir,
		TanggalLahir:   wargaMasterReq.TanggalLahir,
		JenisKelaminId: wargaMasterReq.JenisKelaminId,
		Alamat:         wargaMasterReq.Alamat,
		Rt:             wargaMasterReq.Rt,
		Rw:             wargaMasterReq.Rw,
		NoRumah:        wargaMasterReq.NoRumah,
	}

	if err := db.Create(&warga).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal create",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil create",
			"data":   warga,
		})
	}
}

func UpdateWarga(c *gin.Context) {
	db := config.InitDB()
	var wargaMasterReq wargaMasterRequest
	// get id from url
	wargaUuid := c.Param("warga_uuid")

	var dataWarga WargaMasterModel
	if err := db.Where("warga_uuid = ?", wargaUuid).First(&dataWarga).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	db.Model(&dataWarga).Where("warga_uuid = ?", wargaUuid).Updates(WargaMasterModel{
		NamaLengkap:    wargaMasterReq.NamaLengkap,
		NoKk:           wargaMasterReq.NoKk,
		Nik:            wargaMasterReq.Nik,
		TempatLahir:    wargaMasterReq.TempatLahir,
		TanggalLahir:   wargaMasterReq.TanggalLahir,
		JenisKelaminId: wargaMasterReq.JenisKelaminId,
		Alamat:         wargaMasterReq.Alamat,
		Rt:             wargaMasterReq.Rt,
		Rw:             wargaMasterReq.Rw,
		NoRumah:        wargaMasterReq.NoRumah,
	})

	c.JSON(200, gin.H{
		"status": "Success",
		"data":   dataWarga,
	})
}

func DeleteWarga(c *gin.Context) {
	db := config.InitDB()
	// get id from url
	wargaUuid := c.Param("warga_uuid")

	var dataWarga WargaMasterModel
	if err := db.Where("warga_uuid = ?", wargaUuid).First(&dataWarga).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	db.Where("warga_uuid = ?", wargaUuid).Delete(&dataWarga)

	c.JSON(200, gin.H{
		"status": "Success Delete",
		"data":   wargaUuid,
	})
}
