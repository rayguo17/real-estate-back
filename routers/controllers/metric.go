package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goBack/service/metric"
)

func LogBtnHandler(c *gin.Context) {
	var service metric.LogBtnService
	res := service.Create(c)
	c.JSON(200, res)
}
