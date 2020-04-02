package controllers

import "github.com/gin-gonic/gin"

//IStatusCheckController ...
type IStatusCheckController interface {
	Ping(c *gin.Context)
}
