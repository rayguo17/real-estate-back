package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/goBack/service/track"
)

func CreateUser(c *gin.Context) {
	var service track.TrackService

	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.CreateTrack(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

func GetUsers(c *gin.Context) {
	var service track.TrackService
	res := service.GetTracks(c)
	c.JSON(200, res)
}

func GetUser(c *gin.Context) {
	var service track.TrackService
	res := service.GetTrack(c)
	c.JSON(200, res)
}

func updateUsers(c *gin.Context) {

}
