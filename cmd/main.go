package main

import (
	"log"
	"raven/raven"
)

func main() {
	raven := raven.Raven{}
	raven.Init()
	proxies, _ := raven.FetchAll()
	log.Println(proxies)

}
