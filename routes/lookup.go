package routes

import (
	"erukunrukun/config"
	"erukunrukun/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LookupRequest struct {
	LookupCode   string `json:"lookup_code"`
	Keterangan   string `json:"keterangan"`
	StatusActive bool   `json:"status_active"`
}

func GetListLookup(c *gin.Context) {
	db := config.InitDB()
	lookups := []models.Lookup{}

	// select * from lookups
	if err := db.Find(&lookups).Error; err != nil {
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
		"data":    lookups,
	})
}

func GetOneLookupByUuid(c *gin.Context) {
	db := config.InitDB()
	var lookup models.Lookup

	uuid := c.Param("lookup_uuid")
	if err := db.Where("lookup_uuid = ?", uuid).First(&lookup).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal get",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil get",
			"data":   lookup,
		})
	}
}

func PostLookup(c *gin.Context) {
	db := config.InitDB()
	var LookupReq LookupRequest

	if err := c.BindJSON(&LookupReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	lookup := models.Lookup{
		LookupUuid:   uuid.New(),
		LookupCode:   LookupReq.LookupCode,
		Keterangan:   LookupReq.Keterangan,
		StatusActive: LookupReq.StatusActive,
	}

	if err := db.Create(&lookup).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal create",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil create",
			"data":   lookup,
		})
	}
}

func UpdateLookup(c *gin.Context) {
	db := config.InitDB()
	var lookupReq LookupRequest
	// get id from url
	lookupUuid := c.Param("lookup_uuid")

	var dataLookup models.Lookup
	if err := db.Where("lookup_uuid = ?", lookupUuid).First(&lookupReq).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	db.Model(&lookupReq).Where("lookup_uuid = ?", lookupUuid).Updates(models.Lookup{
		LookupCode:   lookupReq.LookupCode,
		Keterangan:   lookupReq.Keterangan,
		StatusActive: lookupReq.StatusActive,
	})

	c.JSON(200, gin.H{
		"status": "Success",
		"data":   dataLookup,
	})
}

func DeleteLookup(c *gin.Context) {
	db := config.InitDB()
	// get id from url
	lookupUuid := c.Param("lookup_uuid")

	var dataLookup models.Lookup
	if err := db.Where("lookup_uuid = ?", lookupUuid).First(&dataLookup).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	db.Where("lookup_uuid = ?", lookupUuid).Delete(&dataLookup)

	c.JSON(200, gin.H{
		"status": "Success Delete",
		"data":   lookupUuid,
	})
}
