package domain

import "errors"

var (
	ErrTicketPriceZero   = errors.New("ticket price must be greater than zero")
	ErrTicketTypeInvalid = errors.New("ticket type must be 'full' or 'half'")
)

type TicketKind string

const (
	TicketKindFull TicketKind = "full"
	TicketKindHalf TicketKind = "half"
)

type Ticket struct {
	Id         string
	EventId    string
	Spot       *Spot
	TicketKind TicketKind
	Price      float64
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}

	return nil
}

func IsTicketKindValid(TicketKind TicketKind) bool {
	return TicketKind == TicketKindFull || TicketKind == TicketKindHalf
}

func (t *Ticket) CalculatePrice() {
	if t.TicketKind == TicketKindHalf {
		t.Price /= 2
	}
}
