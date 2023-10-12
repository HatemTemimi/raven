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
	r.Scanner.client = &client
	r.Checker.client = &client
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

func (r *Raven) FetchAllToStdOut() error {
	proxies, err := r.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	r.Writer.WriteToStdout(proxies)
	return nil
}

func (r *Raven) FetchValidToStdOut(target string) error {
	proxies, err := r.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	r.Writer.WriteToStdout(workingAgainst)
	return nil
}

func (r *Raven) FetchAllToJsonFile(filePath string) error {
	proxies, err := r.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	err = r.Writer.WriteToJsonFile(proxies, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (r *Raven) FetchValidToJsonFile(target string, filePath string) error {
	proxies, err := r.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	err = r.Writer.WriteToJsonFile(workingAgainst, filePath)
	if err != nil {
		return err
	}
	return nil
}


func (r *Raven) FetchAllToTxtFile(filePath string) error {
	proxies, err := r.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	err = r.Writer.WriteToTxtFile(proxies, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (r *Raven) FetchValidToTxtFile(target string, filePath string) error {
	proxies, err := r.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	err = r.Writer.WriteToTxtFile(workingAgainst, filePath)
	if err != nil {
		return err
	}
	return nil
}


func (r *Raven) FetchFromTxtFile(target string, filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadTxtfile(filePath)
	if err != nil {
		return nil,err
	}
	return proxies, nil
}



func (r *Raven) FetchFromJsonFile(target string, filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadJsonFile(filePath)
	if err != nil {
		return nil,err
	}
	return proxies, nil
}


func (r *Raven) CheckFromJsonFile(target string, filePath string) ([]string, error) {
	proxies, err := r.Reader.ReadJsonFile(filePath)
	if err != nil {
		return nil,err
	}
	workingAgainst := r.Checker.Check(proxies, []string{target})
	return workingAgainst, nil
}


func (r *Raven) CheckFromTxtToTxt(url string, source string, target string) error {
	proxies, err := r.Reader.ReadTxtfile(source)
	if err != nil {
		return err
	}
	workingAgainst := r.Checker.Check(proxies, []string{url})
	err = r.Writer.WriteToTxtFile(workingAgainst, target)
	if err != nil {
		return err
	}
	return nil
}


func (r *Raven) CheckFromJsonToJson(url string, source string, target string) error {
	proxies, err := r.Reader.ReadJsonFile(source)
	if err != nil {
		return err
	}
	workingAgainst := r.Checker.Check(proxies, []string{url})
	err = r.Writer.WriteToJsonFile(workingAgainst, target)
	if err != nil {
		return err
	}
	return nil
}


