package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/HatemTemimi/Raven/raven"
)

func main() {
	raven := raven.Raven{}
	raven.Init()

	target := flag.String("t", "", "Provide a Target URL to test proxies against, defaults to google")
	fetch := flag.String("f", "all", "Fetch types:\n 1/all: all proxies \n 2/valid: only valid proxies[you must provide a target with -t]")
	output := flag.String("o", "", "Path to the output file, defaults to proxies.json")
	input := flag.String("i", "", "Path to the input file")
	help := flag.String("h", "", "Raven help")

	flag.Parse()
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
		}

	} else if *input != ""{
		var extension =  filepath.Ext(*input)
		if *fetch == "all" {
		//fetching all without validation
			if *target != "" {
			//with target!!err 
				log.Println("the -t (target flag) is used with fetch set to valid --> -fetch valid")
			} else if *target == ""{
				//without target
				if *output == "" {
					//without output
					if  extension == ".txt" {
						raven.FetchFromTxtFile(*input)
					} else if  extension == ".json" {
						raven.FetchFromJsonFile(*input)
					} else {
						log.Println("please provide txt or json files as input")
					}
				} else if *output != "" {
					//with output
					if  extension == ".txt" {
						log.Println("you are fetching all from file, the exported file will be same as the input, try -fetch valid -output proxies.txt to filter out the list")
					} else if  extension == ".json" {
						log.Println("you are fetching all from file, the exported file will be same as the input, try -fetch valid -output proxies.json to filter out the list")
					} 
					//with output
				}
			}

		} else if *fetch == "valid" {
			if *target == "" {
			//fetching valid without target, forwards to google.com
				if *output == ""{
				//without output, forwards to sdtdout
					if  extension == ".txt" {
						//raven.CheckFromjj(*input)

						//needs func fetchvalidfromtxt to stdout
					} else if  extension == ".json" {
						//needs func fetchvalidfromjson to stdout
					} else {
						log.Println("please provide txt or json files as input")
					}
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
		}

	} else if *fetch == "" || *help != "" {
		//asking for help
		flag.PrintDefaults()
	}
	

}
