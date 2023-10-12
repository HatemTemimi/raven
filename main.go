package main

import (
	"flag"
	"log"
	"github.com/HatemTemimi/Raven/raven"
)

func main() {
	raven := raven.Raven{}
	raven.Init()

	target := flag.String("t", "", "Provide a Target URL to test proxies against, defaults to google")
	fetch := flag.String("f", "all", "Fetch types:\n 1/all: all proxies \n 2/valid: only valid proxies[you must provide a target with -t]")
	output := flag.String("o", "", "Path to the output file, defaults to proxies.json")
	//input := flag.String("i", "", "Path to the input file")
	help := flag.String("h", "", "Raven help")

	flag.Parse()

	if *fetch == "all" {
	//fetching all without validation
		if *target != "" {
		//with target!!err 
			log.Println("the -t (target flag) is used with fetch set to valid --> -fetch valid")
		} else if *target == ""{
			//without target
			if *output == "" {
				//without output
				raven.FetchAllToStdOut()
			} else if *output != "" {
				//with output
				raven.FetchAllToJsonFile(*output)
			}

		}

	} else if *fetch == "valid" {
		if *target == "" {
		//fetching valid without target, forwards to google.com
			if *output == ""{
			//without output, forwards to sdtdout
				log.Println("no custom target provided, testing against default target.")
				raven.FetchValidToStdOut("www.google.com")
			} else if *output != ""{
			//with output
				log.Println("no custom target provided, testing against default target.")
				raven.FetchValidToJsonFile("www.google.com", *output)
			}
		} else if *target != "" {
			//with target
			if *output == ""{
				//without output
				raven.FetchValidToStdOut(*target)
			} else if *output != "" {
				//with output
				raven.FetchValidToJsonFile(*target, *output)
			}
		}
	}else if *fetch == "" || *help != "" {
		//asking for help
		flag.PrintDefaults()
	}
	

}
