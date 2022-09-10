package main

import (
	"ecommerce-project/config"
	"ecommerce-project/migration"
	"ecommerce-project/utils/database"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetDataBase()
	db := database.DB(config)

	migration.InitialMigration(db)
	e := echo.New()

	// factory.InitFact(e, db)

	log.Fatal(e.Start(fmt.Sprintf(":%s", config.DB_SERVERPORT)))

}
