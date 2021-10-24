package model

type RailwayStation struct {
	ID       uint64
	Name     string
	Location string
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type RailwayStationEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *RailwayStation
}
