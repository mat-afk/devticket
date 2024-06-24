package service

import "fmt"

type PartnerFactory interface {
	CreatePartner(partnerId int) (Partner, error)
}

type DefaultPartnerFactory struct {
	PartnersBaseURLs map[int]string
}

func NewPartnerFactory(partnersBaseURLs map[int]string) PartnerFactory {
	return &DefaultPartnerFactory{
		PartnersBaseURLs: partnersBaseURLs,
	}
}

func (f *DefaultPartnerFactory) CreatePartner(partnerId int) (Partner, error) {
	baseURL, ok := f.PartnersBaseURLs[partnerId]
	if !ok {
		return nil, fmt.Errorf("partner with id %d not found", partnerId)
	}

	switch partnerId {
	case 1:
		return &Partner1{
			BaseURL: baseURL,
		}, nil
	case 2:
		return &Partner2{
			BaseURL: baseURL,
		}, nil
	}

	return nil, fmt.Errorf("partner with id %d not found", partnerId)
}
