package router

import (
	"apartment/app/auth"
	"apartment/app/customer"
	"apartment/app/room"
	"apartment/middleware/authorizetion"
	"apartment/utility"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func V1(e *echo.Echo, db *gorm.DB) {

	g := e.Group("/api/v1")

	secretKey := utility.GetSecretKey()

	authorizetionRepository := authorizetion.NewRepository(db)
	authorizetionService := authorizetion.NewService(authorizetionRepository, secretKey)
	authorizetionHandler := authorizetion.NewHandler(authorizetionService)

	// auth
	{
		authRepository := auth.NewRepository(db)
		authService := auth.NewService(authRepository, secretKey)
		authHandler := auth.NewHandler(authService)

		authRoute := g.Group("/auth")
		authRoute.POST("/login", authHandler.HandlerLogin)
		authRoute.POST("/register", authHandler.HandlerRegister)
		authRoute.POST("/logout", authHandler.HandlerLogout, authorizetionHandler.Handler)
	}

	// customer
	{
		customerRepository := customer.NewRepository(db)
		customerService := customer.NewService(customerRepository)
		customerHandler := customer.NewHandler(customerService)

		customerRoute := g.Group("/customer")

		customerRoute.Use(authorizetionHandler.Handler)

		customerRoute.GET("", customerHandler.HandlerFindAll)
		customerRoute.GET("/:id", customerHandler.HandlerFindOne)
		customerRoute.POST("", customerHandler.HandlerCreate)
		customerRoute.PUT("/:id", customerHandler.HandlerUpdate)
		customerRoute.PUT("/profile/:id", customerHandler.HandlerUploadProfile)
		customerRoute.DELETE("/:id", customerHandler.HandlerRemove)
		customerRoute.DELETE("/destory/:id", customerHandler.HandlerDestory)
	}

	//rooms
	{
		roomRepository := room.NewRepository(db)
		roomService := room.NewService(roomRepository)
		roomHandler := room.NewHandler(roomService)

		roomRoute := g.Group("/room")

		roomRoute.Use(authorizetionHandler.Handler)

		roomRoute.GET("/", roomHandler.HandlerFindAll)
		roomRoute.GET("/:id", roomHandler.HandlerFindOne)
		roomRoute.POST("/", roomHandler.HandlerCreate)
		roomRoute.PUT("/:id", roomHandler.HandlerUpdate)
		roomRoute.PUT("/picture/:id", roomHandler.HandlerUploadPicture)
		roomRoute.DELETE("/:id", roomHandler.HandlerRemove)
		roomRoute.DELETE("/destory/:id", roomHandler.HandlerDestory)
	}
}
