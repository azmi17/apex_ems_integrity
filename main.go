package main

import (
	"integrasi_apex_ems/app"
	"integrasi_apex_ems/controller"
	"integrasi_apex_ems/repository"
	"integrasi_apex_ems/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Init Db:
	db1, db2 := app.NewDB()

	// Init Repository:
	nasabahRepo := repository.NewNasabahRepository()
	tabungRepo := repository.NewTabungRepository()
	sysDaftarUserRepo := repository.NewSysDaftarUserRepository()

	// Init Service:
	apexService := service.NewApexService(nasabahRepo, tabungRepo, sysDaftarUserRepo, db1, db2)

	// Init Controller:
	apexController := controller.NewApexController(apexService)

	router := gin.Default()
	api := router.Group("api/v1")

	// routing path API
	api.GET("/home", apexController.Home)
	api.POST("/institution", apexController.CreateApex)

	router.Run(":4417")

}
