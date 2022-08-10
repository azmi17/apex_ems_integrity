package main

import (
	"fmt"
	"integrasi_apex_ems/app"
	"integrasi_apex_ems/helper"
	"integrasi_apex_ems/model/web"
	"integrasi_apex_ems/repository"
	"integrasi_apex_ems/service"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db1, db2 := app.NewDB()
	nasabahRepo := repository.NewNasabahRepository()
	tabungRepo := repository.NewTabungRepository()
	sysDaftarUserRepo := repository.NewSysDaftarUserRepository()
	apexService := service.NewApexService(nasabahRepo, tabungRepo, sysDaftarUserRepo, db1, db2)

	// Test Input Request..
	request := web.ApexRequest{}
	request.KodeLkm = "0942"
	request.Nama_Lembaga = "BPRKS AZMIRA SEJAHTERA"
	request.Alamat = "Jl. Soekarno-Hatta Puteraco Gading Regency Blok A-2/06"
	request.Telpon = "082115331773"

	apexService.CreateApex(request)
	fmt.Println("data has been created..")

	router := gin.Default()
	api := router.Group("api/v1")

	api.GET("/home", Home)
	router.Run(":4417")

}

func Home(c *gin.Context) {
	appInfo := map[string]interface{}{
		// "App Name":        os.Getenv("application.name"),
		// "App Description": os.Getenv("application.desc"),
		// "App Version":     os.Getenv("application.ver"),
		// "Port Listener":   os.Getenv("application.port"),

		"App Name":        "Apex Integrity",
		"App Description": "Apex Integrity API Endpoint",
		"App Version":     "1.0.0",
		"Port Listener":   "4417",
	}

	response := helper.ApiResponse("Desicription App", "success", appInfo)
	c.JSON(http.StatusOK, response)
}
