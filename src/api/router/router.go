package router

import (
	"GoApiServicePattern/src/api/controllers"

	"github.com/gin-gonic/gin"
)

//InitRouter initialization router
func InitRouter() *gin.Engine {
	r := gin.New()
	controller := controllers.StatusCheckController{}
	r.GET("/ping", controller.Ping)

	return r
}
