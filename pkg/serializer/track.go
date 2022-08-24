package serializer

import (
	"github.com/goBack/models"
	"time"
)

type TrackItem struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func BuildTrack(track models.Track) TrackItem {
	return TrackItem{
		Id:        track.ID,
		Email:     track.Email,
		CreatedAt: track.CreatedAt,
	}
}

func BuildListTracks(track []models.Track, total int) []TrackItem {
	trackItems := make([]TrackItem, 0, total)
	for i := 0; i < total; i++ {
		trackItems = append(trackItems, TrackItem{
			Id:        track[i].ID,
			Email:     track[i].Email,
			CreatedAt: track[i].CreatedAt,
		})
	}
	return trackItems
}
