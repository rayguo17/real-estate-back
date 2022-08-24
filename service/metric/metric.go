package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/goBack/models"
	"github.com/goBack/pkg/serializer"
	"github.com/goBack/pkg/util"
)

type LogBtnService struct {
}

func (service *LogBtnService) Create(c *gin.Context) serializer.Response {
	util.Log().Debug("client ip %s\n", c.Request.Header.Get("X-Forwarded-For"))
	util.Log().Debug("client ip %s\n", c.Request.Header.Get("X-Real-Ip"))
	btn := models.LogButMon{
		Ip: c.ClientIP(),
	}
	id, err := btn.Create()
	if err != nil {
		return serializer.DBErr("failed to insert email", err)

	}
	return serializer.Response{
		Data: id,
	}
}
