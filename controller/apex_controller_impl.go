package controller

import (
	"integrasi_apex_ems/helper"
	"integrasi_apex_ems/model/web"
	"integrasi_apex_ems/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApexControllerImpl struct {
	ApexService service.ApexService
}

// Constructor returning Apex Service
func NewApexController(apexService service.ApexService) ApexController {
	return &ApexControllerImpl{
		ApexService: apexService,
	}
}

func (controller *ApexControllerImpl) Home(c *gin.Context) {

	appInfo := map[string]interface{}{

		// "App Name":        os.Getenv("application.name"),
		// "App Description": os.Getenv("application.desc"),
		// "App Version":     os.Getenv("application.ver"),
		// "Port Listener":   os.Getenv("application.port"),

		"App Name":        "Apex App",
		"App Description": "New Apex API Endpoint",
		"App Version":     "1.0.0",
		"Author":          "Azmi Farhan",
		"Port Listener":   "4417",
	}

	response := helper.ApiResponse("Desicription App", "success", appInfo)
	c.JSON(http.StatusOK, response)
}

func (controller *ApexControllerImpl) CreateApex(c *gin.Context) {

	var request web.ApexRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {

		// Errors validations
		errors := helper.FormatValidationError(err)
		errorMesage := gin.H{"errors": errors}

		response := helper.ApiResponse("Create institution failed", "failed", errorMesage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	apexResponse := controller.ApexService.CreateApex(request)

	response := helper.ApiResponse("Institution has been created", "success", apexResponse)
	c.JSON(http.StatusOK, response)
}
