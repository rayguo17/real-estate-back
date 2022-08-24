package track

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goBack/models"
	"github.com/goBack/pkg/serializer"
	"strconv"
)

type TrackService struct {
	Email string `json:"email" binding:"required"`
	Ip    string `json:"ip" binding:"required"`
}

func (service *TrackService) CreateTrack(c *gin.Context) serializer.Response {
	email := c.PostForm("email")
	fmt.Println("get email:", email)
	fmt.Println("service:", service.Email)
	track := models.Track{Email: service.Email, Ip: service.Ip}
	id, err := track.Create()
	if err != nil {
		return serializer.DBErr("failed to insert email", err)
	}
	return serializer.Response{
		Data: id,
	}

}

func (service *TrackService) GetTrack(c *gin.Context) serializer.Response {
	id := c.Param("id")
	fmt.Println("service:", id)
	idInt, _ := strconv.Atoi(id)
	track := models.GetTrackByID(uint(idInt))
	if track == nil {
		return serializer.Err(serializer.CodeTrackNotFound, "failed to get track", nil)
	}
	fmt.Println("here is track", track)
	return serializer.Response{
		Code: 0,
		Data: serializer.BuildTrack(*track),
	}
}

func (service *TrackService) GetTracks(c *gin.Context) serializer.Response {

	tracks, total := models.GetTracks()
	
	return serializer.Response{
		Code: 0,
		Data: serializer.BuildListTracks(tracks, total),
	}

}
