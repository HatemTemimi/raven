package raven

import (
	"errors"
	"net/http"
	"path/filepath"

	raven "github.com/HatemTemimi/Raven/raven/lib"
)

type Cli struct {
	raven raven.Raven
}

func (c *Cli) Init() {
	client := http.Client{}
	c.raven.Scanner.Client = &client
	c.raven.Checker.Client = &client
}

func (c *Cli) FetchAllToStdOut() error {
	proxies, err := c.raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	c.raven.Writer.WriteToStdout(proxies)
	return nil
}

func (c *Cli) FetchValidToStdOut(target string) error {
	proxies, err := c.raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := c.raven.Checker.Check(proxies, []string{target})
	c.raven.Writer.WriteToStdout(workingAgainst)
	return nil
}

func (c *Cli) FetchAllToFile(path string) error {
	var extension =  filepath.Ext(path)
	if extension == ".txt" {
		err := c.FetchAllToTxtFile(path)
		if err != nil {
			return err
		}
	} else if extension == ".json" {
		err := c.FetchAllToJsonFile(path)
		if err != nil {
			return err
		}
	
	}
	return nil
}

func (c *Cli) FetchAllToJsonFile(filePath string) error {
	proxies, err := c.raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	err = c.raven.Writer.WriteToJsonFile(proxies, filePath)
	if err != nil {
		return err
	}
	return nil
}


func (c *Cli) FetchAllToTxtFile(filePath string) error {
	proxies, err := c.raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	err = c.raven.Writer.WriteToTxtFile(proxies, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cli) FetchValidToJsonFile(target string, filePath string) error {
	proxies, err := c.raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := c.raven.Checker.Check(proxies, []string{target})
	err = c.raven.Writer.WriteToJsonFile(workingAgainst, filePath)
	if err != nil {
		return err
	}
	return nil
}


func (c *Cli) FetchValidToTxtFile(target string, filePath string) error {
	proxies, err := c.raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := c.raven.Checker.Check(proxies, []string{target})
	err = c.raven.Writer.WriteToTxtFile(workingAgainst, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cli) FetchValidFromFileToFile(url string, source string, target string) error {
	var sourceExt =  filepath.Ext(source)
	var targetExt =  filepath.Ext(target)
	if sourceExt == targetExt {
		if sourceExt == ".txt" {
			err := c.FetchValidFromTxtToTxt(url,source, target)
			if err != nil {
				return err
			}
		} else if sourceExt == ".json" {
			err := c.FetchValidFromJsonToJson(url,source,target)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("source file and target file type must be the same.")
	}
	
	return nil
}


func (c *Cli) FetchValidFromFileToStdOut(url string, source string) error {
	var sourceExt =  filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := c.raven.FetchValidFromTxtFile(url,source)
		if err != nil {
			return err
		}
		c.raven.Writer.WriteToStdout(proxies)
		
	} else if sourceExt == ".json" {
		proxies,err := c.raven.FetchValidFromJsonFile(url,source)
		if err != nil {
			return err
		}
		c.raven.Writer.WriteToStdout(proxies)
	}
	return nil
}


func (c *Cli) FetchAllFromFileToStdOut(source string) error {
	var sourceExt =  filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := c.raven.FetchAllFromTxtFile(source)
		if err != nil {
			return err
		}
		c.raven.Writer.WriteToStdout(proxies)
		
	} else if sourceExt == ".json" {
		proxies,err := c.raven.FetchAllFromJsonFile(source)
		if err != nil {
			return err
		}
		c.raven.Writer.WriteToStdout(proxies)
	}
	return nil
}


func (c *Cli) FetchValidFromTxtToTxt(url string, source string, target string) error {
	proxies, err := c.raven.Reader.ReadTxtfile(source)
	if err != nil {
		return err
	}
	workingAgainst := c.raven.Checker.Check(proxies, []string{url})
	err = c.raven.Writer.WriteToTxtFile(workingAgainst, target)
	if err != nil {
		return err
	}
	return nil
}


func (c *Cli) FetchValidFromJsonToJson(url string, source string, target string) error {
	proxies, err := c.raven.Reader.ReadJsonFile(source)
	if err != nil {
		return err
	}
	workingAgainst := c.raven.Checker.Check(proxies, []string{url})
	err = c.raven.Writer.WriteToJsonFile(workingAgainst, target)
	if err != nil {
		return err
	}
	return nil
}

func (r *Cli) FetchValidToFile(target string, path string) error {
	var extension =  filepath.Ext(path)
	if extension == ".txt" {
		err := r.FetchValidToTxtFile(target, path)
		if err != nil {
			return err
		}
	} else if extension == ".json" {
		err := r.FetchValidToJsonFile(target,path)
		if err != nil {
			return err
		}
	}
	return nil
}
