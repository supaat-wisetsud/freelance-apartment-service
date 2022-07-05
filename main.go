package main

import (
	"apartment/database"
	"apartment/migrations"
	"apartment/router"
	"apartment/utility"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var (
	gormInstance *gorm.DB
	logger       = middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} method=${method}, uri=${uri}, status=${status}, origin=${header:origin}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	})
	cors = middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*",
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderCookie,
			echo.HeaderSetCookie,
		},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.OPTIONS,
		},
	})
)

func init() {
	utility.LoadEnv()
	fmt.Printf("Initial service\n")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	var err error
	gormInstance, err = database.InitDB(DBHost, DBUser, DBPass, DBPort, DBName)
	if err != nil {
		panic(err)
	}

	m := migrations.InitMigrations(gormInstance)
	if err := m.Migrate(); err != nil {
		panic(err.Error())
	}
}

func main() {
	e := echo.New()

	e.Use(logger)
	e.Use(cors)

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "OK")
	})

	router.V1(e, gormInstance)

	e.Logger.Fatal(e.Start(":3011"))
}
