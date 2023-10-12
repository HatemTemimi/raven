package raven

import (
	"net/http"
)

type Raven struct {
	Scanner Scanner
	Checker Checker
	Writer  Writer
	Reader Reader
}

func (r *Raven) Init() {
	client := http.Client{}
	r.Scanner.Client = &client
	r.Checker.Client = &client
}

func (r *Raven) FetchAll() ([]string, error) {
	proxies, err := r.Scanner.ScanDefaultSources()
	if err != nil {
		return nil, err
	}
	return proxies, nil
}

func (r *Raven) FetchValid(target string) ([]string, error) {
	proxies, err := r.Scanner.ScanDefaultSources()
	if err != nil {
		return nil, err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	return workingAgainst, nil
}

func (r *Raven) FetchAllFromTxtFile(filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadTxtfile(filePath)
	if err != nil {
		return nil,err
	}
	return proxies, nil
}

func (r *Raven) FetchAllFromJsonFile(filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadJsonFile(filePath)
	if err != nil {
		return nil,err
	}
	return proxies, nil
}


func (r *Raven) FetchValidFromJsonFile(target string, filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadJsonFile(filePath)
	if err != nil {
		return nil,err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	return workingAgainst, nil
}


func (r *Raven) FetchValidFromTxtFile(target string, filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadTxtfile(filePath)
	if err != nil {
		return nil,err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	return workingAgainst, nil
}




