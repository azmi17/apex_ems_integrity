package controller

import "github.com/gin-gonic/gin"

type ApexController interface {
	Home(c *gin.Context)
	CreateApex(c *gin.Context)
}
