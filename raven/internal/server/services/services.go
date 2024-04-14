package services

import (
	"github.com/HatemTemimi/Raven/raven/pkg/lib"
	"github.com/HatemTemimi/Raven/raven/pkg/lib/models"
)

type RavenService struct {
	raven lib.Raven
}

func (service *RavenService) GetAll() ([]models.Proxy, error) {
	proxies, err := service.raven.FetchAll()
	if err != nil {
		return nil, err
	}
	return proxies, nil
}
