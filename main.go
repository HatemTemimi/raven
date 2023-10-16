package main

import (
	"flag"
	"log"
	raven "github.com/HatemTemimi/Raven/raven/cli"
)

func main() {

	cli := raven.Cli{}
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
			} else if *target == ""{
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
				if *output == ""{
				//without output, forwards to sdtdout
					cli.FetchValidToStdOut("www.google.com")
				} else if *output != ""{
				//with output
					cli.FetchValidToFile("www.google.com", *output)
				}
			} else if *target != "" {
				//with target
				if *output == ""{
					//without output
					cli.FetchValidToStdOut(*target)
				} else if *output != "" {
					//with output
					cli.FetchValidToFile(*target, *output)
				}
			}
		}

	} else if *input != ""{
		if *fetch == "all" {
		//fetching all without validation
			if *target != "" {
				//with target!!err 
				log.Println("the -t (target flag) is used with fetch set to valid --> -fetch valid")
			} else if *target == ""{
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
				if *output == ""{
					//without output, forwards to sdtdout
					//neeeds FETCHVALIDfromfileTOSTDOUT
					cli.FetchValidFromFileToStdOut("www.google.com", *input)
				} else if *output != ""{
					cli.FetchValidFromFileToFile("www.google.com", *input, *output)
				}
			} else if *target != "" {
				//with target
				if *output == ""{
					//without output
					cli.FetchValidFromFileToStdOut(*target, *input)
				} else if *output != "" {
					//with output
					cli.FetchValidFromFileToFile(*target,*input, *output)
				}
			}
		}

	} else if *fetch == "" || *help != "" {
		//asking for help
		flag.PrintDefaults()
	}
	

}
