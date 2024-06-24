package service

type ReservationRequest struct {
	EventId    string   `json:"eventId"`
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticketKind"`
	CardHash   string   `json:"cardHash"`
	Email      string   `json:"email"`
}

type ReservationResponse struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticketKind"`
	Status     string `json:"status"`
	EventId    string `json:"eventId"`
}

type Partner interface {
	MakeReservation(req *ReservationRequest) ([]ReservationResponse, error)
}
