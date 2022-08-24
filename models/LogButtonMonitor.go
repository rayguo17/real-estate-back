package models

import (
	"github.com/goBack/pkg/util"
	"github.com/jinzhu/gorm"
)

type LogButMon struct {
	gorm.Model
	Ip string
}

func (l *LogButMon) Create() (uint, error) {
	if err := DB.Create(l).Error; err != nil {
		util.Log().Warning("can't insert btn metric, %s", err)
		return 0, err
	}
	return l.ID, nil
}
