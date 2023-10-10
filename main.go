package main

import (
	"flag"
	"fmt"

	"github.com/HatemTemimi/Raven/raven"
)

func main() {
	raven := raven.Raven{}
	raven.Init()

	target := flag.String("t", "", "Provide a Target URL to test proxies against, defaults to google")
	fetch := flag.String("fetch", "all", "Fetch type; 1/all all proxies 2/valid only valid proxies[you must provide a target with -t]")
	output := flag.String("o", "", "Path to the file, defaults to proxies.json")
	help := flag.String("h", "", "Raven help")
	flag.Parse()

	if *fetch == "all" {
		if *target != "" {
			fmt.Println("the -t (target flag) is used with fetch set to valid --> -fetch valid")
		} else if *target == "" && *output == "" {
			raven.FetchAllToStdOut()
		} else if *target == "" && *output != "" {
			raven.FetchAllToJsonFile(*output)
		}
	} else if *fetch == "valid" {
		if *target == "" && *output == "" {
			fmt.Println("no custom target provided, testing against default target.")
			raven.FetchValidToStdOut("www.google.com")
		} else if *target != "" && *output == "" {
			raven.FetchValidToStdOut(*target)
		} else if *target == "" && *output != "" {
			fmt.Println("no custom target provided, testing against default target.")
			raven.FetchValidToJsonFile("www.google.com", *output)
		} else if *target != "" && *output != "" {
			raven.FetchValidToJsonFile(*target, *output)
		}
	} else if *fetch == "" || *help != "" {
		flag.PrintDefaults()
	}

}
