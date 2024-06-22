package domain

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
