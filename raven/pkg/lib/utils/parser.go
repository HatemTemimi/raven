package utils

import (
	"errors"
	"fmt"
	"github.com/HatemTemimi/Raven/raven/pkg/lib/models"
	"strconv"
	"strings"
)

func ParseProxyFromAddress(url string) (*models.Proxy, error) {
	parts := strings.Split(url, ":")
	if len(parts) != 2 {
		return nil, errors.New("error parsing proxy")
	}

	port, _ := strconv.ParseInt(parts[1], 10, 64)
	proxy := models.Proxy{
		Ip:     parts[0],
		Port:   port,
		Speed:  nil,
		Status: "unknown",
	}
	return &proxy, nil
}

func ProxiesToArray(proxies []models.Proxy) []string {
	var formatted []string
	for _, proxy := range proxies {
		formatted = append(formatted, fmt.Sprintf(proxy.Ip, ":", proxy.Port))
	}
	return formatted
}
