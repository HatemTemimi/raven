package cli

import (
	"errors"
	"flag"
	"log"
	"path/filepath"

	ravenApi "github.com/HatemTemimi/Raven/raven/api"
	raven "github.com/HatemTemimi/Raven/raven/lib"
)

type Cli struct {
	Raven raven.Raven
	api   ravenApi.Api
}

func (cli *Cli) HandleApiFlags() {

}

func (cli *Cli) Init() {

	cli.Raven.Init()

	target := flag.String("t", "", "Provide a Target URL to test proxies against, defaults to google")
	fetch := flag.String("f", "all", "Fetch types:\n 1/all: all proxies \n 2/valid: only valid proxies[you must provide a target with -t]")
	output := flag.String("o", "", "Path to the output file, defaults to proxies.json")
	input := flag.String("i", "", "Path to the input file, json or txt")
	help := flag.String("h", "", "Raven help")

	api := flag.String("api", "", "Starts the server")

	flag.Parse()

	if *api == "start" {
		cli.api.Init()
	}

	//without input, fetching from sources
	if *api == "" {
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
						err := cli.FetchAllToStdOut()
						if err != nil {
							return
						}
					} else if *output != "" {
						//with output
						err := cli.FetchAllToFile(*output)
						if err != nil {
							return
						}
					}
				}

			} else if *fetch == "valid" {
				if *target == "" {
					log.Println("no custom target provided, testing against default target.")
					//fetching valid without target, forwards to google.com
					if *output == "" {
						//without output, forwards to sdtdout
						err := cli.FetchValidToStdOut("www.google.com")
						if err != nil {
							return
						}
					} else if *output != "" {
						//with output
						err := cli.FetchValidToFile("www.google.com", *output)
						if err != nil {
							return
						}
					}
				} else if *target != "" {
					//with target
					if *output == "" {
						//without output
						err := cli.FetchValidToStdOut(*target)
						if err != nil {
							return
						}
					} else if *output != "" {
						//with output
						err := cli.FetchValidToFile(*target, *output)
						if err != nil {
							return
						}
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
						err := cli.FetchValidFromFileToStdOut("www.google.com", *input)
						if err != nil {
							return
						}
					} else if *output != "" {
						err := cli.FetchValidFromFileToFile("www.google.com", *input, *output)
						if err != nil {
							return
						}
					}
				} else if *target != "" {
					//with target
					if *output == "" {
						//without output
						err := cli.FetchValidFromFileToStdOut(*target, *input)
						if err != nil {
							return
						}
					} else if *output != "" {
						//with output
						err := cli.FetchValidFromFileToFile(*target, *input, *output)
						if err != nil {
							return
						}
					}
				}
			}

		} else if *fetch == "" || *help != "" {
			//asking for help
			flag.PrintDefaults()
		}
	}
}

func (cli *Cli) FetchAllToStdOut() error {
	proxies, err := cli.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	cli.Raven.Writer.WriteToStdout(proxies)
	return nil
}

func (cli *Cli) FetchValidToStdOut(target string) error {
	proxies, err := cli.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := cli.Raven.Checker.Check(proxies, []string{target})
	cli.Raven.Writer.WriteToStdout(workingAgainst)
	return nil
}

func (cli *Cli) FetchAllToFile(path string) error {
	var extension = filepath.Ext(path)
	if extension == ".txt" {
		err := cli.FetchAllToTxtFile(path)
		if err != nil {
			return err
		}
	} else if extension == ".json" {
		err := cli.FetchAllToJsonFile(path)
		if err != nil {
			return err
		}

	}
	return nil
}

func (cli *Cli) FetchAllToJsonFile(filePath string) error {
	proxies, err := cli.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	err = cli.Raven.Writer.WriteToJsonFile(proxies, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Cli) FetchAllToTxtFile(filePath string) error {
	proxies, err := cli.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	err = cli.Raven.Writer.WriteToTxtFile(proxies, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Cli) FetchValidToJsonFile(target string, filePath string) error {
	proxies, err := cli.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := cli.Raven.Checker.Check(proxies, []string{target})
	err = cli.Raven.Writer.WriteToJsonFile(workingAgainst, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Cli) FetchValidToTxtFile(target string, filePath string) error {
	proxies, err := cli.Raven.Scanner.ScanDefaultSources()
	if err != nil {
		return err
	}
	workingAgainst := cli.Raven.Checker.Check(proxies, []string{target})
	err = cli.Raven.Writer.WriteToTxtFile(workingAgainst, filePath)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Cli) FetchValidFromFileToFile(url string, source string, target string) error {
	var sourceExt = filepath.Ext(source)
	var targetExt = filepath.Ext(target)
	if sourceExt == targetExt {
		if sourceExt == ".txt" {
			err := cli.FetchValidFromTxtToTxt(url, source, target)
			if err != nil {
				return err
			}
		} else if sourceExt == ".json" {
			err := cli.FetchValidFromJsonToJson(url, source, target)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("source file and target file type must be the same.")
	}

	return nil
}

func (cli *Cli) FetchValidFromFileToStdOut(url string, source string) error {
	var sourceExt = filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := cli.Raven.FetchValidFromTxtFile(url, source)
		if err != nil {
			return err
		}
		cli.Raven.Writer.WriteToStdout(proxies)

	} else if sourceExt == ".json" {
		proxies, err := cli.Raven.FetchValidFromJsonFile(url, source)
		if err != nil {
			return err
		}
		cli.Raven.Writer.WriteToStdout(proxies)
	}
	return nil
}

func (cli *Cli) FetchAllFromFileToStdOut(source string) error {
	var sourceExt = filepath.Ext(source)
	if sourceExt == ".txt" {
		proxies, err := cli.Raven.FetchAllFromTxtFile(source)
		if err != nil {
			return err
		}
		cli.Raven.Writer.WriteToStdout(proxies)

	} else if sourceExt == ".json" {
		proxies, err := cli.Raven.FetchAllFromJsonFile(source)
		if err != nil {
			return err
		}
		cli.Raven.Writer.WriteToStdout(proxies)
	}
	return nil
}

func (cli *Cli) FetchValidFromTxtToTxt(url string, source string, target string) error {
	proxies, err := cli.Raven.Reader.ReadTxtfile(source)
	if err != nil {
		return err
	}
	workingAgainst := cli.Raven.Checker.Check(proxies, []string{url})
	err = cli.Raven.Writer.WriteToTxtFile(workingAgainst, target)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Cli) FetchValidFromJsonToJson(url string, source string, target string) error {
	proxies, err := cli.Raven.Reader.ReadJsonFile(source)
	if err != nil {
		return err
	}
	workingAgainst := cli.Raven.Checker.Check(proxies, []string{url})
	err = cli.Raven.Writer.WriteToJsonFile(workingAgainst, target)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Cli) FetchValidToFile(target string, path string) error {
	var extension = filepath.Ext(path)
	if extension == ".txt" {
		err := cli.FetchValidToTxtFile(target, path)
		if err != nil {
			return err
		}
	} else if extension == ".json" {
		err := cli.FetchValidToJsonFile(target, path)
		if err != nil {
			return err
		}
	}
	return nil
}
