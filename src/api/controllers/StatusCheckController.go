package controllers

import (
	"GoApiServicePattern/src/api/services"

	"github.com/gin-gonic/gin"
)

//StatusCheckController ...
type StatusCheckController struct {
	services.StatusCheckService
}

//Ping ...
func (controller *StatusCheckController) Ping(c *gin.Context) {
	result, err := controller.GetStatusHealth()
	if err != nil {
		//Handle error
	}

	c.JSON(200, gin.H{
		"message": result,
	})

}
