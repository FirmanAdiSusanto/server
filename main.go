package main

import (
	"taskApi/app/configs"
	"taskApi/app/databases"
	"taskApi/app/migrations"
	"taskApi/app/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)

	migrations.InitDBMigration(dbMysql)

	// dbPosgres := databases.InitDBPosgres(cfg)

	// create new instance echo
	e := echo.New()

	routers.InitRouter(e, dbMysql)
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	// start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
