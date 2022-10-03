package main

import (
	"erukunrukun/middleware"
	"erukunrukun/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//when server.go start, it will be run function InitDB (connecting to database)
	// config.InitDB()

	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	router.POST("/login", routes.Login)                              // /api/v1/erukunrukun/login
	router.GET("/users", routes.GetUser)                             // /api/v1/erukunrukun/users
	router.POST("/users", routes.PostUser)                           // /api/v1/erukunrukun/users
	router.PUT("/users/:id", middleware.IsAuth(), routes.UpdateUser) // /api/v1/erukunrukun/users/:id
	// router.DELETE("/users/:id", middleware.IsAuth(), routes.DeleteUser) // /api/v1/erukunrukun/users/:id

	// warga := router.Group("/warga") // /api/v1/erukunrukun
	// {
	// 	warga.GET("/", routes.GetListWarga)                                     // /api/v1/erukunrukun/login
	// 	warga.GET("/:warga_uuid", middleware.IsAuth(), routes.GetOneAnakByUuid) // /api/v1/erukunrukun/users
	// 	warga.POST("/", middleware.IsAuth(), routes.PostWarga)                  // /api/v1/erukunrukun/users
	// 	warga.PUT("/:warga_uuid", middleware.IsAuth(), routes.DeleteWarga)      // /api/v1/erukunrukun/users/:id
	// 	warga.DELETE("/:warga_uuid", middleware.IsAuth(), routes.DeleteWarga)   // /api/v1/erukunrukun/users/:id
	// }

	// router.Run(":8080") // if you want to run on port 8080
	router.Run()
}
