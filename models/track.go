package models

import (
	"github.com/goBack/pkg/util"
	"github.com/jinzhu/gorm"
)

type Track struct {
	gorm.Model
	Email string `gorm:"unique;not null"`
	Ip    string
}

func (track *Track) Create() (uint, error) {
	if err := DB.Create(track).Error; err != nil {
		util.Log().Warning("can't insert database,%s", err)
		return 0, err
	}
	return track.ID, nil

}

func GetTrackByID(id uint) *Track {
	var track Track
	result := DB.First(&track, id)
	if result.Error != nil {
		return nil
	}
	return &track
}

func GetTracks() ([]Track, int) {
	var (
		tracks []Track
		total  int
	)
	DB.Find(&tracks)
	total = len(tracks)
	return tracks, total
}
