package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner1 struct {
	BaseURL string
}

type Partner1ReservationRequest struct {
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticketKind"`
	Email      string   `json:"email"`
}

type Partner1ReservationResponse struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticketKind"`
	Status     string `json:"status"`
	EventId    string `json:"eventId"`
}

func (p *Partner1) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	partnerReq := Partner1ReservationRequest{
		Spots:      req.Spots,
		TicketKind: req.TicketKind,
		Email:      req.Email,
	}

	body, err := json.Marshal(partnerReq)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/events/%s/reserve", p.BaseURL, req.EventId)

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	httpRes, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", httpRes.StatusCode)
	}

	var partnerResponses []Partner1ReservationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&partnerResponses); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResponses))

	for i, res := range partnerResponses {
		responses[i] = ReservationResponse{
			Id:     res.Id,
			Spot:   res.Spot,
			Status: res.Status,
		}
	}

	return responses, nil
}
