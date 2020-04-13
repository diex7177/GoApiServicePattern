package router

import (
	"GoApiServicePattern/src/api/controllers"
	"GoApiServicePattern/src/api/services"

	"github.com/gin-gonic/gin"
)

//InitRouter initialization router
func InitRouter() *gin.Engine {
	r := gin.New()
	service := services.StatusCheckService{}
	controller := controllers.StatusCheckController{service}
	r.GET("/ping", controller.Ping)

	return r
}
