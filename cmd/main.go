package main

import (
	"github.com/HatemTemimi/Raven/raven"
)

func main() {
	raven := raven.Raven{}
	raven.Init()
  raven.FetchValidToStdOut("www.spankbang.com")
}
