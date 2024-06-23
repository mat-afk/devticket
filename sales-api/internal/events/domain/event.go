package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired  = errors.New("event name is required")
	ErrEventDateInThePast = errors.New("event date must be in the future")
	ErrEventCapacityZero  = errors.New("event capacity must be greater than zero")
	ErrEventPriceZero     = errors.New("event price must be greater than zero")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	RatingL16   Rating = "L16"
	RatingL18   Rating = "L18"
)

type Event struct {
	Id           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerId    int
	Spots        []Spot
	Tickets      []Ticket
}

func (e Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrEventDateInThePast
	}

	if e.Capacity <= 0 {
		return ErrEventCapacityZero
	}

	if e.Price <= 0 {
		return ErrEventPriceZero
	}

	return nil
}
