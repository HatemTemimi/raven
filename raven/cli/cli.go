package cli

import (
	"errors"
	"flag"
	"log"
	"path/filepath"

	raven "github.com/HatemTemimi/Raven/raven/lib"
)

type Cli struct {
	Raven raven.Raven
}

func (cli *Cli) Init() {

	cli.Raven.Init()

	target := flag.String("t", "", "Provide a Target URL to test proxies against, defaults to google")
	fetch := flag.String("f", "all", "Fetch types:\n 1/all: all proxies \n 2/valid: only valid proxies[you must provide a target with -t]")
	output := flag.String("o", "", "Path to the output file, defaults to proxies.json")
	input := flag.String("i", "", "Path to the input file, json or txt")
	help := flag.String("h", "", "Raven help")

	flag.Parse()

	//without input, fetching from sources
	if *input == "" {
		if *fetch == "all" {
			//fetching all without validation
			if *target != "" {
				//with target!!err
				log.Println("the -t (target flag) is used with fetch set to valid --> -fetch valid")
			} else if *target == "" {
				//without target
				if *output == "" {
					//without output
					cli.FetchAllToStdOut()
				} else if *output != "" {
					//with output
					cli.FetchAllToFile(*output)
				}
			}

		} else if *fetch == "valid" {
			if *target == "" {
				log.Println("no custom target provided, testing against default target.")
				//fetching valid without target, forwards to google.com
				if *output == "" {
					//without output, forwards to sdtdout
					cli.FetchValidToStdOut("www.google.com")
				} else if *output != "" {
					//with output
					cli.FetchValidToFile("www.google.com", *output)
				}
			} else if *target != "" {
				//with target
				if *output == "" {
					//without output
					cli.FetchValidToStdOut(*target)
				} else if *output != "" {
					//with output
					cli.FetchValidToFile(*target, *output)
				}
			}
		}

	} else if *input != "" {
		if *fetch == "all" {
			//fetching all without validation
			if *target != "" {
				//with target!!err
				log.Println("the -t (target flag) is used with fetch set to valid --> -fetch valid")
			} else if *target == "" {
				//without target
				if *output == "" {
					//without output
					//FetchAllFromFile TO STDOUT
					err := cli.FetchAllFromFileToStdOut(*input)
					if err != nil {
						log.Println(err)
					}

				} else if *output != "" {
					//with output
					log.Println("Fetching same proxies from file to file, try the -fetch valid flag with a custom target.")
				}
			}

		} else if *fetch == "valid" {
			if *target == "" {
				//fetching valid without target, forwards to google.com
				log.Println("no custom target provided, testing against default target.")
				if *output == "" {
					//without output, forwards to sdtdout
					//neeeds FETCHVALIDfromfileTOSTDOUT
					cli.FetchValidFromFileToStdOut("www.google.com", *input)
				} else if *output != "" {
					cli.FetchValidFromFileToFile("www.google.com", *input, *output)
				}
			} else if *target != "" {
				//with target
				if *output == "" {
					//without output
					cli.FetchValidFromFileToStdOut(*target, *input)
				} else if *output != "" {
					//with output
					cli.FetchValidFromFileToFile(*target, *input, *output)
				}
			}
		}

	} else if *fetch == "" || *help != "" {
		//asking for help
		flag.PrintDefaults()
	}
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
	var extension = filepath.Ext(path)
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
	var sourceExt = filepath.Ext(source)
	var targetExt = filepath.Ext(target)
	if sourceExt == targetExt {
		if sourceExt == ".txt" {
			err := c.FetchValidFromTxtToTxt(url, source, target)
			if err != nil {
				return err
			}
		} else if sourceExt == ".json" {
			err := c.FetchValidFromJsonToJson(url, source, target)
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
	var sourceExt = filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := c.Raven.FetchValidFromTxtFile(url, source)
		if err != nil {
			return err
		}
		c.Raven.Writer.WriteToStdout(proxies)

	} else if sourceExt == ".json" {
		proxies, err := c.Raven.FetchValidFromJsonFile(url, source)
		if err != nil {
			return err
		}
		c.Raven.Writer.WriteToStdout(proxies)
	}
	return nil
}

func (c *Cli) FetchAllFromFileToStdOut(source string) error {
	var sourceExt = filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := c.Raven.FetchAllFromTxtFile(source)
		if err != nil {
			return err
		}
		c.Raven.Writer.WriteToStdout(proxies)

	} else if sourceExt == ".json" {
		proxies, err := c.Raven.FetchAllFromJsonFile(source)
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
	var extension = filepath.Ext(path)
	if extension == ".txt" {
		err := r.FetchValidToTxtFile(target, path)
		if err != nil {
			return err
		}
	} else if extension == ".json" {
		err := r.FetchValidToJsonFile(target, path)
		if err != nil {
			return err
		}
	}
	return nil
}
