package main

import (
	"erukunrukun/config"
	"erukunrukun/middleware"
	"erukunrukun/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//when server.go start, it will be run function InitDB (connecting to database)
	config.InitDB()

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

	router.POST("/login", routes.Login)                                 // /api/v1/erukunrukun/login
	router.GET("/users", middleware.IsAuth(), routes.GetUser)           // /api/v1/erukunrukun/users
	router.POST("/users", routes.PostUser)                              // /api/v1/erukunrukun/users
	router.PUT("/users/:id", middleware.IsAuth(), routes.UpdateUser)    // /api/v1/erukunrukun/users/:id
	router.DELETE("/users/:id", middleware.IsAuth(), routes.DeleteUser) // /api/v1/erukunrukun/users/:id

	warga := router.Group("/warga") // /api/v1/erukunrukun
	{
		warga.GET("/", middleware.IsAuth(), routes.GetListWarga)                // /api/v1/erukunrukun/login
		warga.GET("/:warga_uuid", middleware.IsAuth(), routes.GetOneAnakByUuid) // /api/v1/erukunrukun/users
		warga.POST("/", middleware.IsAuth(), routes.PostWarga)                  // /api/v1/erukunrukun/users
		warga.PUT("/:warga_uuid", middleware.IsAuth(), routes.UpdateWarga)      // /api/v1/erukunrukun/users/:id
		warga.DELETE("/:warga_uuid", middleware.IsAuth(), routes.DeleteWarga)   // /api/v1/erukunrukun/users/:id
	}

	lookup := router.Group("/lookup") // /api/v1/erukunrukun
	{
		lookup.GET("/", middleware.IsAuth(), routes.GetListLookup)                  // /api/v1/erukunrukun/login
		lookup.GET("/:lookup_uuid", middleware.IsAuth(), routes.GetOneLookupByUuid) // /api/v1/erukunrukun/users
		lookup.POST("/", middleware.IsAuth(), routes.PostLookup)                    // /api/v1/erukunrukun/users
		lookup.PUT("/:lookup_uuid", middleware.IsAuth(), routes.UpdateLookup)       // /api/v1/erukunrukun/users/:id
		lookup.DELETE("/:lookup_uuid", middleware.IsAuth(), routes.DeleteLookup)    // /api/v1/erukunrukun/users/:id
	}

	// router.Run(":8080") // if you want to run on port 8080
	router.Run()
}
