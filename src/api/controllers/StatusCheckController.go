package controllers

import (
	"GoApiServicePattern/src/api/services"
	"GoApiServicePattern/src/api/viewmodels"

	"github.com/gin-gonic/gin"
)

//StatusCheckController ...
type StatusCheckController struct {
	services.IStatusCheckService
}

//Ping ...
func (controller *StatusCheckController) Ping(c *gin.Context) {
	result, err := controller.GetStatusHealth()

	if err != nil {
		c.JSON(200, viewmodels.Message{Message: err.Error()})
	}

	vm := viewmodels.Message{Message: result}

	c.JSON(200, vm)
}
