package raven

import (
	"bufio"
	"net/http"
)

type Scanner struct {
	Client  *http.Client
	sources []string
}

const (
	TheSpeedX   = "https://raw.githubusercontent.com/TheSpeedX/SOCKS-List/master/http.txt"
	ProxyScrape = "https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=5000&country=all&ssl=all&anonymity=all"
	Clarketm    = "https://raw.githubusercontent.com/clarketm/proxy-list/master/proxy-list-raw.txt"
)

func (s *Scanner) ScanDefaultSources() ([]string, error) {

	var proxies []string
	s.sources = []string{TheSpeedX, ProxyScrape, Clarketm}

	for _, source := range s.sources {
		chunk, err := s.ScanSource(source)
		if err != nil {
			return nil, err
		}
		proxies = append(proxies, chunk...)
	}

	return proxies, nil
}

func (s *Scanner) ScanSource(url string) ([]string, error) {

	var proxies []string

	//get proxies from url: Proxy:PORT
	resp, err := s.Client.Get(url)
	if err != nil {
		return nil, err
	}

	//marshals the proxies
	scanner := bufio.NewScanner(resp.Body)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	//append to proxy array
	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}

	return proxies, nil
}
