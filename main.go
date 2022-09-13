package main

import (
	"ecommerce-project/config"
	"ecommerce-project/factory"
	"ecommerce-project/migration"
	"ecommerce-project/utils/database"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.GetDataBase()
	db := database.DB(config)

	migration.InitialMigration(db)
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	factory.InitFact(e, db)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.DB_SERVERPORT)))

}
