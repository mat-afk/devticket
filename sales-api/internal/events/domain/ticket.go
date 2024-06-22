package domain

type TicketType string

const (
	TicketTypeFull TicketType = "full"
	TicketTypeHalf TicketType = "half"
)

type Ticket struct {
	Id         string
	EventId    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}
