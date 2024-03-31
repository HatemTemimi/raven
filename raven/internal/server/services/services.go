package services

import "github.com/HatemTemimi/Raven/raven/pkg/lib"

type RavenService struct {
	raven lib.Raven
}

func (service *RavenService) GetAll() ([]string, error) {
	proxies, err := service.raven.FetchAll()
	if err != nil {
		return nil, err
	}
	return proxies, nil
}
