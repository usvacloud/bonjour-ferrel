package main

import (
	"log"
	"os"

	"github.com/oleksandr/bonjour"
)

func main() {

	resolver, err := bonjour.NewResolver(nil)
	if err != nil {
		log.Println("Failed to initialize resolver:", err.Error())
		os.Exit(1)
	}

	results := make(chan *bonjour.ServiceEntry)

	go func(results chan *bonjour.ServiceEntry) {
		for e := range results {
			log.Printf(e.Instance, e.Service, e.AddrIPv4, e.Port, e.ServiceRecord, e.Text)
		}
	}(results)

	err = resolver.Browse("_foobar._tcp", "local.", results)
	if err != nil {
		log.Println("Failed to browse:", err.Error())
	}

	select {}
}
