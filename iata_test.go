package iatranslator

import (
	"log"
	"testing"
)

func TestICAO(t *testing.T) {
	icao := "SAS4039"

	iata, num, err := GetFlightIATA(icao)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%s -> %s:%d\n", icao, iata, num)
}
