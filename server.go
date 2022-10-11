package main

import (
	"erukunrukun/config"
	"erukunrukun/middleware"
	"erukunrukun/migrations"
	"erukunrukun/modules/general"
	"erukunrukun/modules/lookup"
	"erukunrukun/modules/user"
	"erukunrukun/modules/warga"

	"github.com/gin-gonic/gin"
)

func main() {
	//when server.go start, it will be run function InitDB (connecting to database)
	config.InitDB()
	migrations.Migration()

	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	// router.Use(cors.Middleware(cors.Config{
	// 	Origins:        "*",
	// 	Methods:        "GET, PUT, POST, DELETE",
	// 	RequestHeaders: "Origin, Authorization, Content-Type",
	// 	ExposedHeaders: "",
	// 	// MaxAge: 50 * time.Second,
	// 	Credentials:     false,
	// 	ValidateHeaders: false,
	// }))

	router.POST("/login", general.Login)                              // /api/v1/erukunrukun/login
	router.GET("/users", middleware.IsAuth(), user.GetUser)           // /api/v1/erukunrukun/users
	router.POST("/users", user.PostUser)                              // /api/v1/erukunrukun/users
	router.PUT("/users/:id", middleware.IsAuth(), user.UpdateUser)    // /api/v1/erukunrukun/users/:id
	router.DELETE("/users/:id", middleware.IsAuth(), user.DeleteUser) // /api/v1/erukunrukun/users/:id

	wargaGroup := router.Group("/warga") // /api/v1/erukunrukun
	{
		wargaGroup.GET("/", middleware.IsAuth(), warga.GetListWarga)                // /api/v1/erukunrukun/login
		wargaGroup.GET("/:warga_uuid", middleware.IsAuth(), warga.GetOneAnakByUuid) // /api/v1/erukunrukun/users
		wargaGroup.POST("/", middleware.IsAuth(), warga.PostWarga)                  // /api/v1/erukunrukun/users
		wargaGroup.PUT("/:warga_uuid", middleware.IsAuth(), warga.UpdateWarga)      // /api/v1/erukunrukun/users/:id
		wargaGroup.DELETE("/:warga_uuid", middleware.IsAuth(), warga.DeleteWarga)   // /api/v1/erukunrukun/users/:id
	}

	lookupGroup := router.Group("/lookup") // /api/v1/erukunrukun
	{
		lookupGroup.GET("/", middleware.IsAuth(), lookup.GetListLookup)                  // /api/v1/erukunrukun/login
		lookupGroup.GET("/:lookup_uuid", middleware.IsAuth(), lookup.GetOneLookupByUuid) // /api/v1/erukunrukun/users
		lookupGroup.POST("/", middleware.IsAuth(), lookup.PostLookup)                    // /api/v1/erukunrukun/users
		lookupGroup.PUT("/:lookup_uuid", middleware.IsAuth(), lookup.UpdateLookup)       // /api/v1/erukunrukun/users/:id
		lookupGroup.DELETE("/:lookup_uuid", middleware.IsAuth(), lookup.DeleteLookup)    // /api/v1/erukunrukun/users/:id
	}

	// router.Run(":8080") // if you want to run on port 8080
	router.Run()
}
