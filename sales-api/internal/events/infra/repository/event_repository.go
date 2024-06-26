package repository

import (
	"database/sql"
	"errors"

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

	query := `
		SELECT id, name, location, organization, rating, date, image_url, capacity, price, partner_id
		FROM events
		WHERE id = ?
	`
	rows, err := r.db.Query(query, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event domain.Event

	err = rows.Scan(
		&event.Id,
		&event.Name,
		&event.Location,
		&event.Organization,
		&event.Rating,
		&event.Date,
		&event.ImageURL,
		&event.Capacity,
		&event.Price,
		&event.PartnerId,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *EventRepositoryImpl) FindSpotsByEventId(eventId string) ([]*domain.Spot, error) {

	query := `
		SELECT id, event_id, name, status, ticket_id
		FROM spots
		WHERE event_id = ?
	`
	rows, err := r.db.Query(query, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []*domain.Spot

	for rows.Next() {
		var spot domain.Spot

		err := rows.Scan(
			&spot.Id,
			&spot.EventId,
			&spot.Name,
			&spot.Status,
			&spot.TicketId,
		)
		if err != nil {
			return nil, err
		}

		spots = append(spots, &spot)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return spots, nil
}

func (r *EventRepositoryImpl) FindSpotByName(eventId, spotName string) (*domain.Spot, error) {

	query := `
		SELECT 
			s.id, s.event_id, s.name, s.status, s.ticket_id, t.id
		FROM spots s
		LEFT JOIN tickets t ON s.id = t.spot_id
		WHERE s.event_id = ? AND s.name = ?
	`
	row := r.db.QueryRow(query, eventId, spotName)

	var spot domain.Spot
	var ticketId sql.NullString

	err := row.Scan(
		&spot.Id,
		&spot.EventId,
		&spot.Name,
		&spot.Status,
		&spot.TicketId,
		&ticketId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrSpotNotFound
		}

		return nil, err
	}

	if ticketId.Valid {
		spot.TicketId = ticketId.String
	}

	return &spot, nil
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

func (r *EventRepositoryImpl) CreateSpot(spot *domain.Spot) error {
	query := `
		INSERT INTO spots (id, event_id, name, status, ticket_id)
		VALUES(?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, spot.Id, spot.EventId, spot.Name, spot.Status, spot.TicketId)

	return err
}

func (r *EventRepositoryImpl) CreateTicket(ticket *domain.Ticket) error {
	query := `
		INSERT INTO tickets (id, event_id, spot_id, ticket_kind, price)
		VALUES(?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, ticket.Id, ticket.EventId, ticket.Spot.Id, ticket.TicketKind, ticket.Price)

	return err
}

func (r *EventRepositoryImpl) ReserveSpot(spotId, ticketId string) error {
	query := `
		UPDATE spots 
		SET status = ?, ticket_id = ?
		WHERE id = ?
	`
	_, err := r.db.Exec(query, domain.SpotStatusSold, ticketId, spotId)

	return err
}
