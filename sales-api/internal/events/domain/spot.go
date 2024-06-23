package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrSpotNameRequired        = errors.New("spot name is required")
	ErrInsuficientCharacters   = errors.New("spot name must be at least 2 characters long")
	ErrSpotNameStartWithLetter = errors.New("spot name must start with a letter")
	ErrSpotNameEndWithNumber   = errors.New("spot name must end with a number")

	ErrInvalidSpotNumber   = errors.New("invalid spot number")
	ErrSpotNotFound        = errors.New("spot not found")
	ErrSpotAlreadyReserved = errors.New("spot already reserved")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	Id       string
	EventId  string
	Name     string
	Status   SpotStatus
	TicketId string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		Id:      uuid.New().String(),
		EventId: event.Id,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil
}

func (s Spot) Validate() error {
	if s.Name == "" {
		return ErrSpotNameRequired
	}

	if len(s.Name) < 2 {
		return ErrInsuficientCharacters
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStartWithLetter
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotNameEndWithNumber
	}

	return nil
}
