package repository

import (
	"database/sql"

	"github.com/mat-afk/devticket/sales-api/internal/events/domain"
)

type EventRepositoryImpl struct {
	db *sql.DB
}

func NewEventRepositoryImpl(db *sql.DB) (domain.EventRepository, error) {
	return &EventRepositoryImpl{db: db}, nil
}

func (r *EventRepositoryImpl) ListEvents() ([]domain.Event, error) {
	return nil, nil
}

func (r *EventRepositoryImpl) FindEventById(eventId string) (*domain.Event, error) {
	return nil, nil
}

func (r *EventRepositoryImpl) FindSpotsByEventId(eventId string) ([]domain.Spot, error) {
	return nil, nil
}

func (r *EventRepositoryImpl) FindSpotByName(eventId, spotName string) (*domain.Spot, error) {
	return nil, nil
}

func (r *EventRepositoryImpl) CreateEvent(event *domain.Event) error {
	return nil
}

func (r *EventRepositoryImpl) CreateSpot(spot *domain.Spot) error {
	return nil
}

func (r *EventRepositoryImpl) CreateTicket(ticket *domain.Ticket) error {
	return nil
}

func (r *EventRepositoryImpl) ReserveSpot(spotId, ticketId string) error {
	return nil
}
