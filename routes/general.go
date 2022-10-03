package routes

import (
	"erukunrukun/config"
	"erukunrukun/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	DateBirth time.Time `json:"date_birth"`
	AdminRole bool      `gorm:"default:0"`
}

func Login(c *gin.Context) {
	db := config.InitDB()
	// var userData models.User

	// username := c.PostForm("username")
	// password := c.PostForm("password")

	var loginReq models.User

	if err := c.BindJSON(&loginReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	username := loginReq.Username
	password := loginReq.Password

	fmt.Println(username)
	fmt.Println(password)

	if err := db.Where("username = ? AND password = ?", username, password).First(&loginReq).Error; err != nil {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	fmt.Println(loginReq)
	var jwtToken = createToken(&models.User{})

	c.JSON(200, gin.H{
		"data":    loginReq,
		"token":   jwtToken,
		"message": "Berhasil Login",
	})
}

func createToken(user *models.User) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.ID,
		"admin_role": user.AdminRole,
		"exp":        time.Now().AddDate(0, 0, 1).Unix(),
		"iat":        time.Now().Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
