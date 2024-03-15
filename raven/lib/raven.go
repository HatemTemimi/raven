package lib

import (
	"net/http"
	"path/filepath"

	Checker "github.com/HatemTemimi/Raven/raven/lib/checker"
	Reader "github.com/HatemTemimi/Raven/raven/lib/reader"
	Scanner "github.com/HatemTemimi/Raven/raven/lib/scanner"
	Writer "github.com/HatemTemimi/Raven/raven/lib/writer"
)

type Raven struct {
	Scanner Scanner.Scanner
	Checker Checker.Checker
	Writer  Writer.Writer
	Reader  Reader.Reader
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
		return nil, err
	}
	return proxies, nil
}

func (r *Raven) FetchAllFromJsonFile(filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadJsonFile(filePath)
	if err != nil {
		return nil, err
	}
	return proxies, nil
}

func (r *Raven) FetchAllFromFile(source string) ([]string, error) {
	var sourceExt = filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := r.FetchAllFromTxtFile(source)
		if err != nil {
			return nil, err
		}
		return proxies, nil

	} else if sourceExt == ".json" {
		proxies, err := r.FetchAllFromJsonFile(source)
		if err != nil {
			return nil, err
		}
		return proxies, nil
	}
	return nil, nil
}

func (r *Raven) FetchValidFromJsonFile(target string, filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadJsonFile(filePath)
	if err != nil {
		return nil, err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	return workingAgainst, nil
}

func (r *Raven) FetchValidFromTxtFile(target string, filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadTxtfile(filePath)
	if err != nil {
		return nil, err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	return workingAgainst, nil
}

func (r *Raven) FetchValidFromFileToStdOut(url string, source string) error {
	var sourceExt = filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := r.FetchValidFromTxtFile(url, source)
		if err != nil {
			return err
		}
		r.Writer.WriteToStdout(proxies)

	} else if sourceExt == ".json" {
		proxies, err := r.FetchValidFromJsonFile(url, source)
		if err != nil {
			return err
		}

		r.Writer.WriteToStdout(proxies)
	}
	return nil
}
