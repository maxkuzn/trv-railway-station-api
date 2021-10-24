package repo

import (
	"github.com/maxkuzn/trv-railway-station-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.RailwayStationEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.RailwayStationEvent) error
	Remove(eventIDs []uint64) error
}
