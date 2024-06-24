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

func (r *EventRepositoryImpl) CreateEvent(event *domain.Event) error {

	query := `
		INSERT INTO events (id, name, location, organization, rating, date, image_url, capacity, price, partner_id)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(
		query,
		event.Id,
		event.Name,
		event.Location,
		event.Organization,
		event.Rating,
		event.Date.Format("2000-01-01 00:00:00"),
		event.ImageURL,
		event.Capacity,
		event.Price,
		event.PartnerId,
	)

	return err
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

func (r *EventRepositoryImpl) CreateSpot(spot *domain.Spot) error {
	return nil
}

func (r *EventRepositoryImpl) CreateTicket(ticket *domain.Ticket) error {
	return nil
}

func (r *EventRepositoryImpl) ReserveSpot(spotId, ticketId string) error {
	return nil
}
