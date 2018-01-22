package main

import (
	"flag"
	"log"

	"github.com/jimmy-go/nobreak"
	_ "github.com/mattn/go-sqlite3"
)

var (
	configFile = flag.String("config", "", "Config file.")
)

func main() {
	flag.Parse()
	log.Printf("STARTING NO-BREAK")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	nb, err := nobreak.New(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	if err := nb.Run(); err != nil {
		log.Fatal(err)
	}
}
