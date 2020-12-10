package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/osm/flen"
	"github.com/osm/postnord"
)

func main() {
	apiKey := flag.String("api-key", "", "PostNord API key")
	id := flag.String("id", "", "Tracking identifier")
	flen.SetEnvPrefix("PN")
	flen.Parse()

	if *apiKey == "" {
		fmt.Println("error: api key can't be empty")
		os.Exit(1)
	}

	if *id == "" {
		fmt.Println("error: id can't be empty")
		os.Exit(1)
	}

	pn := postnord.New(*apiKey, "sv")
	tir, err := pn.FindByIdentifierV5(*id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, s := range tir.TrackingInformationResponse.Shipments {
		fmt.Println("Avsändare", s.Consignor.Name)
		for _, i := range s.Items {
			for _, e := range i.Events {
				fmt.Println(e.EventTime, e.Location.DisplayName, e.EventDescription)
			}
		}

	}
}
