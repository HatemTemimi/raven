package raven

import (
	"errors"
	"path/filepath"

	raven "github.com/HatemTemimi/Raven/raven/lib"
)

type Cli struct {
	Raven raven.Raven
}

func (c *Cli) Init() {
	/*
	client := http.Client{}
	c.Raven.Scanner.Client = &client
	c.Raven.Checker.Client = &client
	*/
}

func (c *Cli) FetchAllToStdOut() error {
	proxies, err := c.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	c.Raven.Writer.WriteToStdout(proxies)
	return nil
}

func (c *Cli) FetchValidToStdOut(target string) error {
	proxies, err := c.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := c.Raven.Checker.Check(proxies, []string{target})
	c.Raven.Writer.WriteToStdout(workingAgainst)
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
	proxies, err := c.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	err = c.Raven.Writer.WriteToJsonFile(proxies, filePath)
	if err != nil {
		return err
	}
	return nil
}


func (c *Cli) FetchAllToTxtFile(filePath string) error {
	proxies, err := c.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	err = c.Raven.Writer.WriteToTxtFile(proxies, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cli) FetchValidToJsonFile(target string, filePath string) error {
	proxies, err := c.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := c.Raven.Checker.Check(proxies, []string{target})
	err = c.Raven.Writer.WriteToJsonFile(workingAgainst, filePath)
	if err != nil {
		return err
	}
	return nil
}


func (c *Cli) FetchValidToTxtFile(target string, filePath string) error {
	proxies, err := c.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := c.Raven.Checker.Check(proxies, []string{target})
	err = c.Raven.Writer.WriteToTxtFile(workingAgainst, filePath)
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
		proxies, err := c.Raven.FetchValidFromTxtFile(url,source)
		if err != nil {
			return err
		}
		c.Raven.Writer.WriteToStdout(proxies)
		
	} else if sourceExt == ".json" {
		proxies,err := c.Raven.FetchValidFromJsonFile(url,source)
		if err != nil {
			return err
		}
		c.Raven.Writer.WriteToStdout(proxies)
	}
	return nil
}


func (c *Cli) FetchAllFromFileToStdOut(source string) error {
	var sourceExt =  filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := c.Raven.FetchAllFromTxtFile(source)
		if err != nil {
			return err
		}
		c.Raven.Writer.WriteToStdout(proxies)
		
	} else if sourceExt == ".json" {
		proxies,err := c.Raven.FetchAllFromJsonFile(source)
		if err != nil {
			return err
		}
		c.Raven.Writer.WriteToStdout(proxies)
	}
	return nil
}


func (c *Cli) FetchValidFromTxtToTxt(url string, source string, target string) error {
	proxies, err := c.Raven.Reader.ReadTxtfile(source)
	if err != nil {
		return err
	}
	workingAgainst := c.Raven.Checker.Check(proxies, []string{url})
	err = c.Raven.Writer.WriteToTxtFile(workingAgainst, target)
	if err != nil {
		return err
	}
	return nil
}


func (c *Cli) FetchValidFromJsonToJson(url string, source string, target string) error {
	proxies, err := c.Raven.Reader.ReadJsonFile(source)
	if err != nil {
		return err
	}
	workingAgainst := c.Raven.Checker.Check(proxies, []string{url})
	err = c.Raven.Writer.WriteToJsonFile(workingAgainst, target)
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
