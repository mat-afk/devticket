package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner2 struct {
	BaseURL string
}

type Partner2ReservationRequest struct {
	Lugares      []string `json:"lugares"`
	TipoIngresso string   `json:"tipoIngresso"`
	Email        string   `json:"email"`
}

type Partner2ReservationResponse struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Lugar        string `json:"lugar"`
	TipoIngresso string `json:"tipoIngresso"`
	Estado       string `json:"estado"`
	EventoId     string `json:"eventoId"`
}

func (p *Partner2) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	partnerReq := Partner2ReservationRequest{
		Lugares:      req.Spots,
		TipoIngresso: ConvertTicketKind(req.TicketKind),
		Email:        req.Email,
	}

	body, err := json.Marshal(partnerReq)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/eventos/%s/reservar", p.BaseURL, req.EventId)

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

	var partnerResponses []Partner2ReservationResponse
	if err := json.NewDecoder(httpRes.Body).Decode(&partnerResponses); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResponses))

	for i, res := range partnerResponses {
		responses[i] = ReservationResponse{
			Id:     res.Id,
			Spot:   res.Lugar,
			Status: res.Estado,
		}
	}

	return responses, nil
}

func ConvertTicketKind(ticketKind string) string {
	if ticketKind == "full" {
		return "inteira"
	}

	return "meia"
}
