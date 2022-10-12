package user

import (
	"erukunrukun/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserReqest struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	DateBirth string `json:"date_birth"`
	AdminRole bool   `json:"admin_role"`
}

func GetUser(c *gin.Context) {
	db := config.InitDB()
	users := []UserModel{}

	// select * from User
	if err := db.Find(&users).Error; err != nil {
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
		"message": "GET data user",
		"data":    users,
	})
}

func PostUser(c *gin.Context) {
	db := config.InitDB()
	var userReq UserReqest

	if err := c.BindJSON(&userReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	user := UserModel{
		UserUuid:  uuid.New(),
		Username:  userReq.Username,
		Password:  userReq.Password,
		Email:     userReq.Email,
		DateBirth: userReq.DateBirth,
		AdminRole: userReq.AdminRole,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal create",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil create",
			"data":   user,
		})
	}
}

func UpdateUser(c *gin.Context) {
	db := config.InitDB()
	// var roleAdmin bool
	var userReq UserReqest

	if err := c.BindJSON(&userReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	// get id from url
	userId := c.Param("id")

	var dataUser UserModel
	if err := db.Where("id = ?", userId).First(&dataUser).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	db.Model(&dataUser).Where("id = ?", userId).Updates(UserModel{
		FullName:  userReq.FullName,
		DateBirth: userReq.DateBirth,
		AdminRole: userReq.AdminRole,
		Username:  userReq.Username,
		Password:  userReq.Password,
		Email:     userReq.Email,
	})

	c.JSON(200, gin.H{
		"status": "Success",
		"data":   dataUser,
	})
}

func DeleteUser(c *gin.Context) {
	db := config.InitDB()
	// get id from url
	userId := c.Param("id")

	var dataUser UserModel
	if err := db.Where("id = ?", userId).First(&dataUser).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	db.Where("id = ?", userId).Delete(&dataUser)

	c.JSON(200, gin.H{
		"status": "Success Delete",
		"data":   dataUser,
	})
}
