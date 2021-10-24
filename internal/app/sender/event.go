package sender

import (
	"github.com/maxkuzn/trv-railway-station-api/internal/model"
)

type EventSender interface {
	Send(event *model.RailwayStationEvent) error
}
